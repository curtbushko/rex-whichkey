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
		!popupResourceID(request.ClientID, "client") || !popupConnectionGeneration(request.ConnectedAt) ||
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

// popupOriginConnection uses Rex's native info.kind classification on the
// captured session's attachments: only app is eligible, never CLI/automation or
// an unknown kind. This keeps classification and membership in one response.
// An absent selector requires exactly one live app; an explicit selector cannot
// bypass that eligibility. Missing/unknown state is connected per Rex's contract;
// only lost is disconnected. See docs/launch-contract.md for the schema source.
func popupOriginConnection(ctx context.Context, client *rex.Client, sessionID, requested string) (string, string, error) {
	connections, err := popupAppConnections(ctx, client, sessionID)
	if err != nil {
		return "", "", err
	}
	matches, id, generation := 0, "", ""
	for _, connection := range connections {
		if requested == "" || requested == connection.ID {
			matches++
			id, generation = connection.ID, connection.ConnectedAt
		}
	}
	if requested != "" && matches == 0 {
		return "", "", errors.New("whichkey: supplied client selector is not a live graphical app attached to the origin session; use --client with an exact full attached app ID or correct REX_CLIENT")
	}
	if matches == 0 {
		return "", "", errors.New("whichkey: origin session has no attached live graphical Rex app; connect the app to this session and relaunch")
	}
	if matches > 1 {
		return "", "", errors.New("whichkey: origin session has several attached live graphical Rex apps; supply --client with the exact full app client ID")
	}
	return id, generation, nil
}

// originBlock validates the connection generation independently of the Lua
// registry. Configured mode also uses it for focus restoration after dismissal;
// native mode uses it only once an executable selection has bound an app.
func (r popupRequest) originBlock(ctx context.Context, client *rex.Client) (rex.Block, error) {
	_, generation, err := popupOriginConnection(ctx, client, r.SessionID, r.ClientID)
	if err != nil {
		return rex.Block{}, err
	}
	if generation != r.ConnectedAt {
		return rex.Block{}, errors.New("whichkey: origin client connection changed; relaunch from the original terminal")
	}
	return r.resourceOriginBlock(ctx, client)
}

// resourceOriginBlock checks only the captured endpoint/session/block. Native
// display, focus and liveness do not grant authority to execute app actions.
// Configured callers must use originBlock as well, never infer policy from an
// absent client field in a handoff.
func (r popupRequest) resourceOriginBlock(ctx context.Context, client *rex.Client) (rex.Block, error) {
	session, err := client.Session(ctx, r.SessionID)
	if err != nil {
		return rex.Block{}, errors.New("whichkey: origin session is unavailable; relaunch inside the original Rex terminal")
	}
	block, err := session.Block(ctx, r.BlockID)
	if err != nil {
		return rex.Block{}, errors.New("whichkey: origin block is missing, detached or ambiguously placed in REX_SESSION; relaunch inside the original Rex terminal")
	}
	return block, nil
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
func RunPopupController(parent context.Context, input io.Reader) error {
	request, config, err := decodePopupRequest(input)
	if err != nil {
		return err
	}
	executable, err := popupRexExecutable()
	if err != nil {
		return err
	}
	return runPopupLifecycle(parent, request, config, executable, func(ctx context.Context, message Completion) error {
		return request.complete(ctx, executable, message)
	})
}

// runPopupLifecycle always enforces the configured handoff's original expiry.
// Missing/zero expiry cannot select the native foreground policy.
func runPopupLifecycle(parent context.Context, request popupRequest, config Config, executable string,
	finish func(context.Context, Completion) error,
) error {
	return runPopupWithLifetime(parent, request, config, executable, popupConfigured, finish)
}

type popupLifetime uint8

const (
	popupConfigured popupLifetime = iota
	popupNative
)

// runPopupWithLifetime shares owned transport/teardown, not lifetime authority.
// Only the Go-native caller supplies popupNative; decoded handoffs cannot do so.
// finish receives only a retained selection, after cleanup, never child actions.
func runPopupWithLifetime(parent context.Context, request popupRequest, config Config, executable string,
	lifetime popupLifetime, finish func(context.Context, Completion) error,
) (outcome error) {
	ctx, stop := signal.NotifyContext(parent, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()
	var expires, uiDeadline time.Time
	switch lifetime {
	case popupConfigured:
		expires = time.Unix(request.Expires, 0)
		now := time.Now()
		if !expires.After(now) || expires.After(now.Add(InvocationLifetimeSeconds*time.Second)) {
			return errors.New("whichkey: invalid or expired popup lifetime")
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(ctx, expires)
		defer cancel()
		// Reserve the original five seconds for restoration/cleanup/completion.
		uiDeadline = expires.Add(-5 * time.Second)
	case popupNative:
		if request.Expires != 0 || request.Completion != "" || request.ClientID != "" || request.ConnectedAt != "" {
			return errors.New("whichkey: native foreground popup cannot carry a configured lifetime, completion action or pre-bound app")
		}
	default:
		return errors.New("whichkey: invalid popup lifetime policy")
	}
	client, err := rex.New(rex.Options{Executable: executable, Server: request.Server,
		Session: request.SessionID, Timeout: 3 * time.Second})
	if err != nil {
		return errors.New("whichkey: cannot construct the origin client")
	}
	// Only the trusted Go caller selects native resource-only liveness. Empty
	// decoded client fields or expiry values cannot select this policy.
	originBlock := func(check context.Context) (rex.Block, error) {
		if lifetime == popupNative {
			return request.resourceOriginBlock(check, client)
		}
		return request.originBlock(check, client)
	}
	// A wrong endpoint must fail read-only before arming cancellation/focus.
	// Configured mode additionally requires its captured app generation here.
	// Never compare, stringify or substitute Rex's opaque Lua server identity.
	if _, err := originBlock(ctx); err != nil {
		return err
	}
	var layer rex.Layer
	var directory string
	var lease *os.File
	closed, attempted, focusAttempted := false, false, false
	closePopup := func(cleanup context.Context) error {
		if closed || layer.BlockID == "" {
			return nil
		}
		closed = true // never retry an uncertain resource mutation
		if lifetime == popupNative {
			// The exact owned block can be cleaned even after origin loss or
			// app disconnect. Inspect its placement in the captured session;
			// never search another endpoint/session or close a whole layer.
			if _, err := layer.Block().Inspect(cleanup); err != nil {
				return errors.New("whichkey: cannot inspect the exact owned popup for cleanup")
			}
		} else if _, err := originBlock(cleanup); err != nil {
			return err
		}
		if err := layer.Close(cleanup); err != nil {
			return errors.New("whichkey: owned popup close failed or could not confirm exact layer removal; no cascading cleanup or retry")
		}
		return nil
	}
	focusOrigin := func(cleanup context.Context) error {
		focusAttempted = true // never repeat a failed/uncertain focus mutation
		origin, err := originBlock(cleanup)
		if err != nil {
			return err
		}
		if err := origin.Focus(cleanup); err != nil {
			return errors.New("whichkey: cannot restore the exact origin focus; no selected action dispatched")
		}
		return nil
	}
	removeFiles := func() error {
		if lease != nil {
			writer := lease
			lease = nil // a failed close is not retried
			if err := writer.Close(); err != nil {
				return errors.New("whichkey: cannot close native popup lifetime pipe; no selected action dispatched")
			}
		}
		if directory == "" {
			return nil
		}
		if err := removePopupFiles(directory); err != nil {
			return errors.New("whichkey: cannot remove private popup transport; no selected action dispatched")
		}
		directory = ""
		return nil
	}
	defer func() {
		// Cleanup survives caller cancellation, but is itself bounded. If only
		// an incomplete creation identity was returned, no cascading layer,
		// window, session or label-based cleanup is safe or attempted.
		cleanup, done := context.WithTimeout(context.WithoutCancel(ctx), 5*time.Second)
		defer done()
		outcome = errors.Join(outcome, closePopup(cleanup))
		if !attempted && !focusAttempted {
			outcome = errors.Join(outcome, focusOrigin(cleanup))
		}
		outcome = errors.Join(outcome, removeFiles())
		if !attempted {
			attempted = true
			// Lua has already revoked an expired capability. Do not activate
			// completion after expiry, even for cancellation, or retry it.
			if expires.IsZero() || time.Now().Before(expires) {
				outcome = errors.Join(outcome, finish(cleanup, Completion{
					Version: CompletionVersion, Capability: request.Capability, Kind: CompletionCancel,
				}))
			}
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
	command := []string{self, "popup", "--directory", directory}
	if lifetime == popupNative {
		// The child watches this pipe, not a renewable timestamp or PID. EOF
		// cancels it even if the foreground parent is killed without defers.
		lease, err = createPopupLease(directory)
		if err != nil {
			return err
		}
		command = append(command, "--native")
	} else {
		if !uiDeadline.After(time.Now()) {
			return nil // Too little authenticated lifetime remains; cancel safely.
		}
		command = append(command, "--deadline", uiDeadline.UTC().Format(time.RFC3339))
	}
	origin, err := originBlock(ctx)
	if err != nil {
		return err
	}
	initialBounds, initialSize, err := popupInitialSize(ctx, origin, config)
	if err != nil {
		return err
	}
	if lifetime == popupConfigured && !uiDeadline.After(time.Now()) {
		return nil // Read-only geometry sampling must not spend the teardown reserve.
	}
	layer, err = origin.NewLayerWithSize(ctx, initialBounds, command, initialSize)
	if err != nil {
		return errors.New("whichkey: popup creation failed or uncertain; not retried")
	}
	// Keep the initial rectangle for this invocation. Navigation and redraws
	// adapt content to the terminal cells, never change owned layer bounds.
	err = waitPopup(ctx, request, executable, layer.BlockID, originBlock)
	if err != nil {
		if lifetime == popupConfigured && errors.Is(err, context.DeadlineExceeded) &&
			ctx.Err() == context.DeadlineExceeded && parent.Err() == nil && !time.Now().Before(expires) {
			return nil // Controller expiry: cleanup only, never a stale selection.
		}
		return err
	}
	if !expires.IsZero() && !time.Now().Before(expires) {
		return nil
	}
	message, err := readPopupResult(directory, request.Capability, config.Spec)
	if err != nil {
		return err
	}
	if err := closePopup(ctx); err != nil {
		return err // never dispatch after uncertain owned-resource cleanup
	}
	if err := focusOrigin(ctx); err != nil {
		return err
	}
	// An action may close the caller's pane/tab/session. All terminal/Rex/file
	// cleanup must therefore finish BEFORE invoking either dispatch route.
	if err := removeFiles(); err != nil {
		return err
	}
	if _, err := originBlock(ctx); err != nil {
		return err
	}
	attempted = true // a failed/uncertain dispatch must never send a second result
	dispatch, done := context.WithTimeout(ctx, 3*time.Second)
	defer done()
	return finish(dispatch, message)
}

// waitPopup deliberately does not use Layer.Wait: that wrapper caps the whole
// CLI process at its RPC timeout. `rex wait --help` documents unbounded event
// delivery, with --timeout still bounding connections/nonstreaming requests.
// Watch captured resource liveness too, plus app generation only in configured
// mode. Native display survives GUI ambiguity/disconnection. Neither policy
// checks active-tab visibility: switching tabs must not cancel or refocus it.
func waitPopup(ctx context.Context, request popupRequest, executable, blockID string,
	originBlock func(context.Context) (rex.Block, error),
) error {
	waitCtx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)
	stopped := make(chan struct{})
	var watcherErr error // Read only after stopped closes; no watcher survives cleanup.
	go func() {
		defer close(stopped)
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-waitCtx.Done():
				return
			case <-ticker.C:
				// Only resource liveness and configured app generation remain.
				// Child exit stops checks; teardown joins an in-flight bounded
				// check. Parent cancellation still interrupts it through ctx.
				check, done := context.WithTimeout(ctx, 3*time.Second)
				_, err := originBlock(check)
				done()
				if err != nil {
					watcherErr = err
					cancel(err)
					return
				}
			}
		}
	}()
	command := exec.CommandContext(waitCtx, executable,
		"--server", request.Server, "--autostart=false", "--timeout", "3s",
		"wait", "--session", request.SessionID, "--", blockID)
	output := &launchOutput{remaining: popupMessageLimit}
	command.Stdout, command.Stderr = output, output
	command.WaitDelay = 250 * time.Millisecond
	err := command.Run()
	cause := context.Cause(waitCtx)
	cancel(nil)
	<-stopped // No watcher/RPC goroutine outlives teardown.
	if watcherErr != nil {
		return watcherErr // Includes failures reported after the child's wait completed.
	}
	if cause != nil {
		return cause
	}
	if err != nil {
		return errors.New("whichkey: popup child did not exit successfully or its status was lost; no selected action dispatched")
	}
	return nil
}

// A read end makes opening the writer nonblocking and portable; only the writer
// is retained by the foreground parent. No bytes, credentials or renewals travel
// over the pipe. A child sees EOF as soon as the last parent writer disappears.
func createPopupLease(directory string) (*os.File, error) {
	path := filepath.Join(directory, "lifetime")
	if err := unix.Mkfifo(path, 0600); err != nil {
		return nil, errors.New("whichkey: cannot create native popup lifetime pipe")
	}
	reader, err := os.OpenFile(path, os.O_RDONLY|unix.O_NONBLOCK|unix.O_NOFOLLOW, 0)
	if err != nil {
		return nil, errors.New("whichkey: cannot open native popup lifetime pipe")
	}
	writer, err := os.OpenFile(path, os.O_WRONLY|unix.O_NONBLOCK|unix.O_NOFOLLOW, 0)
	closeErr := reader.Close()
	if err != nil || closeErr != nil {
		failure := errors.New("whichkey: cannot retain native popup lifetime pipe")
		if writer != nil {
			return nil, errors.Join(failure, writer.Close())
		}
		return nil, failure
	}
	return writer, nil
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

// Only the reserved transport names inside the private directory are unlinked, then
// the empty directory. Root-scoped removal cannot traverse a child symlink;
// unexpected contents cause directory removal to fail, never recursive deletion.
func removePopupFiles(directory string) error {
	root, err := openPopupDirectory(directory)
	if err != nil {
		return err
	}
	var outcome error
	for _, name := range []string{"config.json", "result.json", "lifetime"} {
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

// RunPopupTerminal runs the bounded configured PTY child. Like the native child,
// it sees only configuration/private results, never capabilities or Lua tickets.
// Terminal cleanup must succeed before committing a result and acknowledging it
// by process exit. The independent controller owns all Rex cleanup and dispatch.
func RunPopupTerminal(parent context.Context, directory string, deadline time.Time, noColor, ascii bool) error {
	if !deadline.After(time.Now()) || deadline.After(time.Now().Add(InvocationLifetimeSeconds*time.Second)) {
		return errors.New("whichkey: invalid popup terminal deadline")
	}
	return runPopupTerminal(parent, directory, deadline, noColor, ascii)
}

// RunNativePopupTerminal has no interaction deadline, but requires the native
// parent's private lifetime pipe. Missing configured deadlines never enter here.
func RunNativePopupTerminal(parent context.Context, directory string, noColor, ascii bool) (outcome error) {
	root, err := openPopupDirectory(directory)
	if err != nil {
		return err
	}
	defer func() { outcome = errors.Join(outcome, root.Close()) }()
	lease, err := root.OpenFile("lifetime", os.O_RDONLY|unix.O_NONBLOCK|unix.O_NOFOLLOW, 0)
	if err != nil {
		return errors.New("whichkey: native popup requires its foreground lifetime pipe")
	}
	defer func() { outcome = errors.Join(outcome, lease.Close()) }()
	info, err := lease.Stat()
	if err != nil || info.Mode()&os.ModeNamedPipe == 0 || info.Mode().Perm() != 0600 {
		return errors.New("whichkey: invalid native popup lifetime pipe")
	}
	ctx, cancel := context.WithCancelCause(parent)
	defer cancel(nil)
	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		var buffer [1]byte
		for {
			// Nonblocking read: EAGAIN means a writer is alive; EOF also
			// catches parent death before the child opened the pipe.
			n, err := unix.Read(int(lease.Fd()), buffer[:])
			if err == nil || !terminalRetry(err) {
				if n > 0 {
					cancel(errors.New("whichkey: unexpected native popup lifetime data"))
				} else {
					cancel(errors.New("whichkey: native popup foreground parent was lost"))
				}
				return
			}
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
			}
		}
	}()
	defer func() {
		cancel(nil)
		<-stopped
	}()
	return runPopupTerminal(ctx, directory, time.Time{}, noColor, ascii)
}

func runPopupTerminal(parent context.Context, directory string, deadline time.Time, noColor, ascii bool) (outcome error) {
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
	ctx := parent
	if !deadline.IsZero() {
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(parent, deadline)
		defer cancel()
	}
	// The layer already carries the requested position. Anchor popup content
	// at its top rather than centering it again inside the fixed rectangle.
	// This local copy does not change standalone RunTerminal positioning.
	config.Position = "top"
	event, err := RunTerminal(ctx, config, noColor, ascii)
	// Only the exact expected deadline error is graceful. Joined terminal
	// restoration/I/O failures and other child errors must still fail closed.
	if !deadline.IsZero() && err == context.DeadlineExceeded && ctx.Err() == context.DeadlineExceeded && parent.Err() == nil {
		event, err = MenuEvent{Dismissed: true}, nil
	}
	if err != nil {
		return err
	}
	if !deadline.IsZero() && !time.Now().Before(deadline) {
		event = MenuEvent{Dismissed: true}
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
