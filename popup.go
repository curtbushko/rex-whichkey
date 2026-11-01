package whichkey

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"

	"golang.org/x/sys/unix"

	"temp/rex"
)

const popupMessageLimit = 64 << 10

// PreparePopup generates secret transport data for capture by the issuing Lua
// action, never for display. Resolving this executable here keeps both later
// processes on the server host, independent of a graphical client's paths.
func PreparePopup(input io.Reader, output io.Writer) error {
	return PreparePopupWithServer(context.Background(), input, output, nil)
}

// PreparePopupWithServer adds an optional immutable endpoint override to the
// captured preparation. A nil override uses the execution-host endpoint policy;
// a present empty override is invalid. Secret output must remain on a pipe.
func PreparePopupWithServer(ctx context.Context, input io.Reader, output io.Writer, server *string) error {
	if file, ok := output.(*os.File); ok {
		info, err := file.Stat()
		if err != nil || info.Mode()&os.ModeCharDevice != 0 {
			return errors.New("whichkey: preparation requires captured output, not a terminal")
		}
	}
	if _, err := DecodeConfig(input); err != nil {
		return errors.New("whichkey: invalid prepared popup configuration")
	}
	if _, err := popupRexExecutable(); err != nil {
		return err
	}
	endpoint, err := rex.ResolveEndpoint(ctx, server)
	if err != nil {
		return err
	}
	capability, err := NewInvocationCapability()
	if err != nil {
		return err
	}
	executable, err := os.Executable()
	if err != nil {
		return errors.New("whichkey: cannot resolve server-side executable")
	}
	return json.NewEncoder(output).Encode(struct {
		Capability string `json:"capability"`
		Executable string `json:"executable"`
		Server     string `json:"server"`
	}{capability, executable, endpoint})
}

// popupRequest is private launcher plumbing, not a second mapping schema. The
// configuration string is decoded exclusively through the Go-owned Config
// boundary. No callback/ticket or caller-supplied action travels in this request.
type popupRequest struct {
	Configuration string `json:"configuration"`
	Capability    string `json:"capability"`
	Server        string `json:"server"`
	SessionID     string `json:"session_id"`
	BlockID       string `json:"block_id"`
	ClientID      string `json:"client_id"`
	ConnectedAt   string `json:"connected_at"`
	Completion    string `json:"completion_action"`
	Expires       int64  `json:"expires"`
}

func decodePopupRequest(input io.Reader) (popupRequest, Config, error) {
	data, err := io.ReadAll(io.LimitReader(input, (8<<20)+1))
	if err != nil || len(data) > 8<<20 || !utf8.Valid(data) || !json.Valid(data) {
		return popupRequest{}, Config{}, errors.New("whichkey: invalid popup handoff")
	}
	if err := checkConfigJSON(json.NewDecoder(bytes.NewReader(data)), false); err != nil {
		return popupRequest{}, Config{}, errors.New("whichkey: invalid popup handoff")
	}
	var request popupRequest
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		return popupRequest{}, Config{}, errors.New("whichkey: invalid popup handoff")
	}
	// Handoffs validate directly, never ResolveEndpoint: missing transport data
	// must not acquire an ambient endpoint. Freeze the canonical spelling here.
	server, serverErr := rex.CanonicalEndpoint(request.Server)
	if ValidateInvocationCapability(request.Capability) != nil || serverErr != nil ||
		!popupResourceID(request.SessionID, "session") || !popupResourceID(request.BlockID, "block") ||
		!popupResourceID(request.ClientID, "client") || strings.TrimSpace(request.ConnectedAt) == "" ||
		request.Completion == "" || strings.Trim(request.Completion, "abcdefghijklmnopqrstuvwxyz0123456789_") != "" {
		return popupRequest{}, Config{}, errors.New("whichkey: popup requires exact origin and completion identities")
	}
	request.Server = server
	now := time.Now()
	deadline := time.Unix(request.Expires, 0)
	if !deadline.After(now) || deadline.After(now.Add(InvocationLifetimeSeconds*time.Second)) {
		return popupRequest{}, Config{}, errors.New("whichkey: invalid or expired popup lifetime")
	}
	config, err := DecodeConfig(strings.NewReader(request.Configuration))
	if err != nil {
		return popupRequest{}, Config{}, errors.New("whichkey: invalid popup configuration")
	}
	return request, config, nil
}

func popupResourceID(id, kind string) bool {
	value := strings.TrimPrefix(id, kind+":")
	if value == id || len(value) != 36 {
		return false
	}
	for i, char := range value {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			if char != '-' {
				return false
			}
		} else if !strings.ContainsRune("0123456789abcdefABCDEF", char) {
			return false
		}
	}
	return true
}

func popupBounds(position string) rex.Bounds {
	switch position {
	case "top":
		return rex.Bounds{X: 0, Y: 0, W: 1, H: 0.4}
	case "bottom":
		return rex.Bounds{X: 0, Y: 0.6, W: 1, H: 0.4}
	case "left":
		return rex.Bounds{X: 0, Y: 0, W: 0.4, H: 1}
	case "right":
		return rex.Bounds{X: 0.6, Y: 0, W: 0.4, H: 1}
	default: // Config.Validate already restricts this to center.
		return rex.Bounds{X: 0.1, Y: 0.2, W: 0.8, H: 0.6}
	}
}

// originBlock validates the connection generation independently of the Lua
// registry. This also guards focus restoration after dismissal, for which the
// cancellation handler deliberately performs no focus operation.
func (r popupRequest) originBlock(ctx context.Context, client *rex.Client) (rex.Block, error) {
	var result struct {
		Clients []struct {
			ID          string `json:"client_id"`
			ConnectedAt string `json:"connected_at"`
			State       string `json:"connection_state"`
		} `json:"clients"`
	}
	if err := client.Call(ctx, "session.list_clients", map[string]string{"session_id": r.SessionID}, &result); err != nil {
		return rex.Block{}, errors.New("whichkey: origin client inspection failed")
	}
	matches := 0
	for _, connection := range result.Clients {
		if connection.ID == r.ClientID {
			if connection.State == "lost" || connection.ConnectedAt != r.ConnectedAt {
				return rex.Block{}, errors.New("whichkey: origin client connection changed")
			}
			matches++
		}
	}
	if matches != 1 {
		return rex.Block{}, errors.New("whichkey: origin client is missing or ambiguous")
	}
	session, err := client.Session(ctx, r.SessionID)
	if err != nil {
		return rex.Block{}, errors.New("whichkey: origin session is unavailable")
	}
	return session.Block(ctx, r.BlockID)
}

// complete invokes the installed CONFIG action, never a new Lua script VM. It
// sends capability-bearing data only on stdin and suppresses child diagnostics.
// Failure is uncertain: callers must mark the attempt before calling, never
// retry it or send a cancellation that could mask a possibly executed action.
func (r popupRequest) complete(ctx context.Context, executable string, message Completion) error {
	data, err := json.Marshal(message)
	if err != nil || len(data) > popupMessageLimit {
		return errors.New("whichkey: invalid popup completion")
	}
	arguments, err := json.Marshal(struct {
		Message string `json:"message"`
	}{string(data)})
	if err != nil {
		return errors.New("whichkey: cannot encode popup completion")
	}
	command := exec.CommandContext(ctx, executable,
		"--server", r.Server, "--client", r.ClientID, "--autostart=false", "--timeout", "3s",
		"do", r.Completion, "--session", r.SessionID, "--block", r.BlockID, "--args", "-")
	command.Stdin = bytes.NewReader(arguments)
	command.Stdout, command.Stderr = io.Discard, io.Discard
	command.WaitDelay = time.Second
	if err := command.Run(); err != nil {
		return errors.New("whichkey: completion failed or uncertain; not retried")
	}
	return nil
}

func popupRexExecutable() (string, error) {
	name := "rex"
	if directory := os.Getenv("REX_BIN_DIR"); directory != "" {
		if !filepath.IsAbs(directory) {
			return "", errors.New("whichkey: REX_BIN_DIR must be server-local and absolute")
		}
		name = filepath.Join(directory, name)
	}
	path, err := exec.LookPath(name)
	if err != nil {
		return "", errors.New("whichkey: cannot resolve Rex on the execution host")
	}
	return path, nil
}

// RunPopupController runs once in a detached server-host process, not in the
// popup's PTY. Its lifetime, state and files belong to exactly one invocation.
// Closing the UI (or a selected action closing the origin) cannot kill this
// process through terminal teardown. The caller never logs its secret stdin.
func RunPopupController(parent context.Context, input io.Reader) (outcome error) {
	request, config, err := decodePopupRequest(input)
	if err != nil {
		return err
	}
	executable, err := popupRexExecutable()
	if err != nil {
		return err
	}
	ctx, stop := signal.NotifyContext(parent, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()
	ctx, cancel := context.WithDeadline(ctx, time.Unix(request.Expires, 0))
	defer cancel()
	client, err := rex.New(rex.Options{Executable: executable, Server: request.Server,
		Session: request.SessionID, Timeout: InvocationLifetimeSeconds * time.Second})
	if err != nil {
		return errors.New("whichkey: cannot construct the origin client")
	}
	// A wrong endpoint must fail read-only. In particular, do not arm deferred
	// cancellation/focus until the exact session, placed block and captured
	// client connection generation have all been found through that endpoint.
	// Never compare, stringify or substitute Rex's opaque Lua server identity.
	origin, err := request.originBlock(ctx, client)
	if err != nil {
		return err
	}
	var layer rex.Layer
	var directory string
	closed, attempted := false, false
	closePopup := func(cleanup context.Context) error {
		if closed || layer.BlockID == "" {
			return nil
		}
		closed = true // never retry an uncertain resource mutation
		return layer.Close(cleanup)
	}
	defer func() {
		// Cleanup survives caller cancellation, but is itself bounded. If only
		// an incomplete creation identity was returned, no cascading layer,
		// window, session or label-based cleanup is safe or attempted.
		cleanup, done := context.WithTimeout(context.WithoutCancel(ctx), 5*time.Second)
		defer done()
		outcome = errors.Join(outcome, closePopup(cleanup))
		if !attempted {
			if origin, err := request.originBlock(cleanup, client); err == nil {
				outcome = errors.Join(outcome, origin.Focus(cleanup))
			} else {
				outcome = errors.Join(outcome, err)
			}
			outcome = errors.Join(outcome, request.complete(cleanup, executable, Completion{
				Version: CompletionVersion, Capability: request.Capability, Kind: CompletionCancel,
			}))
		}
		if directory != "" {
			outcome = errors.Join(outcome, removePopupFiles(directory))
		}
	}()
	directory, err = os.MkdirTemp("", "rex-whichkey-") // Go creates this with mode 0700.
	if err != nil {
		return errors.New("whichkey: cannot create private popup transport")
	}
	if err := writePopupFile(directory, "config.json", []byte(request.Configuration)); err != nil {
		return err
	}
	if err := writePopupFile(directory, "result.json", nil); err != nil {
		return err
	}
	self, err := os.Executable()
	if err != nil {
		return errors.New("whichkey: cannot resolve popup child executable")
	}
	// Leave five seconds of the original (never renewed) lifetime for terminal
	// restoration, exact-resource cleanup and the fresh completion activation.
	uiDeadline := time.Unix(request.Expires, 0).Add(-5 * time.Second)
	if !uiDeadline.After(time.Now()) {
		return errors.New("whichkey: insufficient invocation lifetime to open popup")
	}
	layer, err = origin.NewLayer(ctx, popupBounds(config.Position), []string{
		self, "popup", "--directory", directory, "--deadline", uiDeadline.UTC().Format(time.RFC3339),
	})
	if err != nil {
		return errors.New("whichkey: popup creation failed or uncertain; not retried")
	}
	waitCtx, done := context.WithDeadline(ctx, uiDeadline)
	err = layer.Wait(waitCtx)
	done()
	if err != nil {
		return errors.New("whichkey: popup child did not exit successfully")
	}
	message, err := readPopupResult(directory, request.Capability, config.Spec)
	if err != nil {
		return err
	}
	if err := closePopup(ctx); err != nil {
		return err // never dispatch after uncertain owned-resource cleanup
	}
	origin, err = request.originBlock(ctx, client)
	if err != nil {
		return err
	}
	if err := origin.Focus(ctx); err != nil {
		return err
	}
	attempted = true
	return request.complete(ctx, executable, message)
}

func writePopupFile(directory, name string, data []byte) (outcome error) {
	file, err := os.OpenFile(filepath.Join(directory, name), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return errors.New("whichkey: cannot create private popup file")
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	if _, err := file.Write(data); err != nil {
		return errors.New("whichkey: cannot write private popup file")
	}
	return nil
}

// Only the two reserved names inside the private directory are unlinked, then
// the empty directory. Root-scoped removal cannot traverse a child symlink;
// unexpected contents cause directory removal to fail, never recursive deletion.
func removePopupFiles(directory string) error {
	root, err := openPopupDirectory(directory)
	if err != nil {
		return err
	}
	var outcome error
	for _, name := range []string{"config.json", "result.json"} {
		if err := root.Remove(name); err != nil && !errors.Is(err, os.ErrNotExist) {
			outcome = errors.Join(outcome, err)
		}
	}
	outcome = errors.Join(outcome, root.Close())
	if err := os.Remove(directory); err != nil && !errors.Is(err, os.ErrNotExist) {
		outcome = errors.Join(outcome, err)
	}
	return outcome
}

func openPopupDirectory(directory string) (*os.Root, error) {
	info, err := os.Lstat(directory)
	if err != nil || !filepath.IsAbs(directory) || !info.IsDir() || info.Mode().Perm() != 0700 {
		return nil, errors.New("whichkey: popup directory must be private and absolute")
	}
	owner, ok := info.Sys().(*syscall.Stat_t)
	if !ok || owner.Uid != uint32(os.Geteuid()) {
		return nil, errors.New("whichkey: popup directory has the wrong owner")
	}
	root, err := os.OpenRoot(directory)
	if err != nil {
		return nil, errors.New("whichkey: cannot open private popup directory")
	}
	return root, nil
}

func readPopupResult(directory, capability string, mappings []Mapping) (message Completion, outcome error) {
	root, err := openPopupDirectory(directory)
	if err != nil {
		return Completion{}, err
	}
	defer func() { outcome = errors.Join(outcome, root.Close()) }()
	file, err := root.OpenFile("result.json", os.O_RDONLY|unix.O_NOFOLLOW, 0)
	if err != nil {
		return Completion{}, errors.New("whichkey: cannot open popup result")
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	data, err := io.ReadAll(io.LimitReader(file, popupMessageLimit+1))
	if err != nil || len(data) > popupMessageLimit {
		return Completion{}, errors.New("whichkey: popup result exceeds transport bound")
	}
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		return Completion{Version: CompletionVersion, Capability: capability, Kind: CompletionCancel}, nil
	}
	// Wrap the raw selection for the authoritative strict Completion decoder;
	// decoding CompletionSelection alone would accept missing false flags.
	wrapped, err := json.Marshal(struct {
		Version    int             `json:"version"`
		Capability string          `json:"capability"`
		Kind       string          `json:"kind"`
		Selection  json.RawMessage `json:"selection"`
	}{CompletionVersion, capability, CompletionComplete, data})
	if err != nil || json.Unmarshal(wrapped, &message) != nil || message.ValidateSelection(mappings) != nil {
		return Completion{}, errors.New("whichkey: popup returned an invalid retained selection")
	}
	return message, nil
}

// RunPopupTerminal is the only PTY-owned child. It sees configuration and a
// private result path, never the capability or original Lua ticket. Terminal
// cleanup must succeed before a result is written, closed and acknowledged by
// process exit. The independent controller owns all Rex cleanup and dispatch.
func RunPopupTerminal(parent context.Context, directory string, deadline time.Time, noColor, ascii bool) (outcome error) {
	if !deadline.After(time.Now()) || deadline.After(time.Now().Add(InvocationLifetimeSeconds*time.Second)) {
		return errors.New("whichkey: invalid popup terminal deadline")
	}
	root, err := openPopupDirectory(directory)
	if err != nil {
		return err
	}
	defer func() { outcome = errors.Join(outcome, root.Close()) }()
	input, err := root.OpenFile("config.json", os.O_RDONLY|unix.O_NOFOLLOW, 0)
	if err != nil {
		return errors.New("whichkey: cannot open popup configuration")
	}
	config, decodeErr := DecodeConfig(input)
	if err := errors.Join(decodeErr, input.Close()); err != nil {
		return err
	}
	ctx, cancel := context.WithDeadline(parent, deadline)
	defer cancel()
	event, err := RunTerminal(ctx, config, noColor, ascii)
	if err != nil {
		return err
	}
	output, err := root.OpenFile("result.json", os.O_WRONLY|os.O_TRUNC|unix.O_NOFOLLOW, 0)
	if err != nil {
		return errors.New("whichkey: cannot open popup result output")
	}
	if err := errors.Join(WriteMenuResult(output, event), output.Close()); err != nil {
		return fmt.Errorf("whichkey: commit popup result: %w", err)
	}
	return nil
}
