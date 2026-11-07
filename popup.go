package whichkey

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/unix"

	"temp/rex"
)

// Full keys can consume nearly the entire snapshot. Reserve bounded space for
// private rename text and result framing without truncating admitted keys.
const popupMessageLimit = menuConfigLimit + (64 << 10)

// popupRequest is private launcher plumbing, not a second mapping schema. The
// configuration string is decoded exclusively through the Go-owned Config
// boundary. No callback/ticket or caller-supplied action travels in this request.
type popupRequest struct {
	Configuration string
	Server        string
	SessionID     string
	BlockID       string
	ClientID      string
	ConnectedAt   string
	image         *runtimeImage
	placement     *rex.Block
	snapshot      *rex.SessionSnapshot
	retained      *nativeMenuState
	invocationID  string
	selection     *rex.CLISelection
	noColor       bool
	ascii         bool
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
// Callers must supply the exact pane-owner or explicit binding. This validator
// never chooses a client from list order or prompts for one. Missing/unknown
// state is connected per Rex's contract; only lost is disconnected.
func popupOriginConnection(ctx context.Context, client *rex.Client, sessionID, requested string) (string, string, error) {
	if !popupResourceID(requested, "client") {
		return "", "", errors.New("whichkey: exact graphical app binding is missing or malformed")
	}
	connections, err := popupAppConnections(ctx, client, sessionID)
	if err != nil {
		return "", "", err
	}
	for _, connection := range connections {
		if requested == connection.ID {
			return connection.ID, connection.ConnectedAt, nil
		}
	}
	return "", "", errors.New("whichkey: bound client is not a live graphical app attached to the origin session")
}

// originBlock checks the captured app generation and exact resource origin.
// Resource-only display/restoration uses resourceOriginBlock instead.
func (r popupRequest) originBlock(ctx context.Context, client *rex.Client) (rex.Block, error) {
	_, generation, err := popupOriginConnection(ctx, client, r.SessionID, r.ClientID)
	if err != nil {
		return rex.Block{}, err
	}
	if generation != r.ConnectedAt {
		return rex.Block{}, errors.New("whichkey: origin client connection changed; relaunch from the original terminal")
	}
	block, err := r.resourceOriginBlock(ctx, client)
	if err != nil {
		return rex.Block{}, err
	}
	return block, nil
}

// resourceOriginBlock checks only the captured endpoint/session/block. Native
// display, focus and liveness do not grant authority to execute app actions.
// An executable app action additionally checks originBlock; absent app fields
// never select a replacement graphical target.
func (r popupRequest) resourceOriginBlock(ctx context.Context, client *rex.Client) (rex.Block, error) {
	session, err := client.Session(ctx, r.SessionID)
	if err != nil {
		return rex.Block{}, errors.New("whichkey: origin session is unavailable; relaunch inside the original Rex terminal")
	}
	block, err := session.Block(ctx, r.BlockID)
	if err != nil {
		return rex.Block{}, errors.New("whichkey: origin block is missing, detached or ambiguously placed in REX_SESSION; relaunch inside the original Rex terminal")
	}
	if r.placement != nil && (block.WindowID != r.placement.WindowID || block.LayerID != r.placement.LayerID) {
		return rex.Block{}, errors.New("whichkey: origin placement changed; not retargeted")
	}
	return block, nil
}

// runPopupLifecycle owns transport, PTY lifetime and ordered cleanup. There is
// no interaction deadline. finish only captures a result after cleanup succeeds.
func runPopupLifecycle(parent context.Context, request popupRequest, config Config, selection *rex.CLISelection,
	finish func(context.Context, MenuEvent) error,
) (outcome error) {
	ctx, stop := signal.NotifyContext(parent, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()
	ctx, cancelLifecycle := context.WithCancelCause(ctx)
	defer cancelLifecycle(nil)
	request.selection = selection
	request.invocationID = rand.Text()

	client, err := rex.NewContext(ctx, rex.Options{Selection: selection, Server: request.Server,
		Session: request.SessionID, Timeout: 3 * time.Second})
	if err != nil {
		return errors.New("whichkey: origin CLI/endpoint compatibility failed; check the execution-host integration and relaunch from the original terminal; not retried")
	}
	// Display and restoration require resources, not GUI eligibility.
	originBlock := func(check context.Context) (rex.Block, error) {
		return request.resourceOriginBlock(check, client)
	}
	// A wrong endpoint must fail read-only before arming cancellation/focus.
	// No viewing-host or opaque Lua server value can replace this endpoint.
	if _, err := originBlock(ctx); err != nil {
		return err
	}
	var layer rex.Layer
	var directory string
	var lease *os.File
	var updates *popupRenameUpdates
	updatesStopped := false
	stopUpdates := func() error {
		if updates == nil || updatesStopped {
			return nil
		}
		updatesStopped = true
		return updates.stop()
	}
	closed, attempted, focusAttempted := false, false, false
	closePopup := func(cleanup context.Context) error {
		if closed || layer.BlockID == "" {
			return nil
		}
		closed = true // never retry an uncertain resource mutation
		// Origin/app loss never authorizes closing anything but this block.
		if _, inspectErr := layer.Block().Inspect(cleanup); inspectErr != nil {
			return errors.New("whichkey: cannot inspect the exact owned popup for cleanup")
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
		outcome = errors.Join(outcome, stopUpdates())
		outcome = errors.Join(outcome, closePopup(cleanup))
		if !attempted && !focusAttempted {
			outcome = errors.Join(outcome, focusOrigin(cleanup))
		}
		outcome = errors.Join(outcome, removeFiles())
		if !attempted {
			attempted = true
			outcome = errors.Join(outcome, finish(cleanup, MenuEvent{Dismissed: true}))
		}
	}()
	directory, err = newRuntimeDirectory("rex-whichkey-") // Go creates mode 0700, then canonicalizes its own directory.
	if err != nil {
		return errors.New("whichkey: cannot create private popup transport")
	}
	if err := writePopupFile(directory, "config.json", []byte(request.Configuration)); err != nil {
		return err
	}
	if config.textEntry != nil {
		data, err := json.Marshal(config.textEntry)
		if err != nil || len(data) > popupMessageLimit {
			return errors.New("whichkey: cannot encode native text-entry presentation")
		}
		if err := writePopupFile(directory, nativeTextFile, data); err != nil {
			return err
		}
	}
	if err := writePopupFile(directory, "result.json", nil); err != nil {
		return err
	}
	for _, name := range popupUpdateFiles() {
		if slotErr := writePopupFile(directory, name, nil); slotErr != nil {
			return slotErr
		}
	}
	imageCtx, finishImage := context.WithTimeout(ctx, 20*time.Second)
	imageErr := request.preparePopupImage(imageCtx, directory)
	finishImage()
	if imageErr != nil {
		return imageErr
	}
	command := []string{filepath.Join(directory, runtimeImageName), "popup", "--directory", directory, "--native"}
	lease, err = createPopupLease(directory)
	if err != nil {
		return err
	}
	if request.noColor {
		command = append(command, "--no-color")
	}
	if request.ascii {
		command = append(command, "--ascii")
	}

	origin, err := originBlock(ctx)
	if err != nil {
		return err
	}
	initialBounds, initialSize, err := popupInitialSize(ctx, origin, config)
	if err != nil {
		return err
	}
	updates, err = startPopupRenameUpdates(ctx, cancelLifecycle, directory, request, client, config)
	if err != nil {
		return err
	}
	layer, err = origin.NewLayerWithSize(ctx, initialBounds, command, initialSize)
	if err != nil {
		return errors.New("whichkey: popup creation failed or uncertain; not retried")
	}
	// Keep the initial rectangle for this invocation. Navigation and redraws
	// adapt content to the terminal cells, never change owned layer bounds.
	err = waitPopup(ctx, request, client, layer.BlockID, originBlock)
	if err = errors.Join(err, stopUpdates(), context.Cause(ctx)); err != nil {
		return err
	}

	config = updates.config // Controller and every worker have joined.
	if request.retained != nil {
		request.retained.config, request.retained.targets = config, updates.targets
	}
	message, err := readNativePopupResult(directory, config)
	if err != nil {
		return err
	}
	if err := updates.validateSelection(message, config); err != nil {
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
// Watch captured placement/liveness too. Display survives GUI disconnect.
// Switching active tabs alone must not cancel or refocus the menu.
func waitPopup(ctx context.Context, request popupRequest, client *rex.Client, blockID string,
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
				// Only captured resource placement and liveness govern UI lifetime.
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
	session, err := client.Session(waitCtx, request.SessionID)
	if err == nil {
		var block rex.Block
		block, err = session.Block(waitCtx, blockID)
		if err == nil {
			err = block.WaitTerminal(waitCtx)
		}
	}

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
	for _, name := range append([]string{"config.json", "result.json", "lifetime", nativeTextFile, runtimeImageName, "runtime.json"}, popupUpdateFiles()...) {
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

// readNativePopupResult decodes data only. It never constructs a callback
// capability/completion/receipt or accepts a replacement action/resource ID.
func readNativePopupResult(directory string, config Config) (event MenuEvent, outcome error) {
	root, err := openPopupDirectory(directory)
	if err != nil {
		return MenuEvent{}, err
	}
	defer func() { outcome = errors.Join(outcome, root.Close()) }()
	file, err := root.OpenFile("result.json", os.O_RDONLY|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if err != nil {
		return MenuEvent{}, errors.New("whichkey: cannot open popup result")
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	if slotErr := checkPopupSlot(file, false); slotErr != nil {
		return MenuEvent{}, slotErr
	}
	data, err := io.ReadAll(io.LimitReader(file, popupMessageLimit+1))
	if err != nil || len(data) > popupMessageLimit {
		return MenuEvent{}, errors.New("whichkey: popup result exceeds transport bound")
	}
	return decodeNativePopupResult(data, config)
}

func decodeNativePopupResult(data []byte, config Config) (event MenuEvent, err error) {
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		return MenuEvent{Dismissed: true}, nil
	}
	fields, err := completionObject(data, "keys", "confirmed", "text", "update", "provider")
	if err != nil {
		return MenuEvent{}, err
	}
	keys, err := completionString(fields["keys"])
	if err != nil {
		return MenuEvent{}, err
	}
	var confirmed bool
	if raw, present := fields["confirmed"]; !present || json.Unmarshal(raw, &confirmed) != nil {
		return MenuEvent{}, errors.New("whichkey: missing or invalid popup confirmation")
	}
	var text *string
	if raw, present := fields["text"]; present {
		value, textErr := completionString(raw)
		if textErr != nil {
			return MenuEvent{}, textErr
		}
		text = &value
	}
	event, err = retainedNativeEvent(config, keys, confirmed, text)
	if err != nil {
		return MenuEvent{}, err
	}
	if raw, present := fields["update"]; present {
		var binding renameBinding
		if err := decodeRuntime(raw, &binding, popupMessageLimit); err != nil {
			return MenuEvent{}, err
		}
		event.update = &binding
	}
	if raw, present := fields["provider"]; present {
		var binding renameBinding
		if err := decodeRuntime(raw, &binding, popupMessageLimit); err != nil {
			return MenuEvent{}, err
		}
		event.provider = &binding
	}
	return event, nil
}

// The retained record, never the result's labels/action/args, defines execution.
func retainedNativeEvent(config Config, keys string, confirmed bool, text *string) (MenuEvent, error) {
	canonical, err := NormalizeKeys(keys)
	if err != nil || canonical != keys {
		return MenuEvent{}, errors.New("whichkey: popup result requires canonical keys")
	}
	for _, mapping := range config.Spec {
		if mapping.Keys != keys {
			continue
		}
		if !mapping.Enabled() || mapping.Action == "" || mapping.CallbackRef != "" || mapping.Confirm != confirmed {
			return MenuEvent{}, errors.New("whichkey: popup result does not match a retained Go action")
		}
		if inlineRenameMapping(mapping) != (text != nil) {
			return MenuEvent{}, errors.New("whichkey: popup text does not match retained rename intent")
		}
		if text != nil {
			if config.textEntry == nil {
				return MenuEvent{}, errors.New("whichkey: no inline rename form was retained")
			}
			if textErr := validateRenameText(*text); textErr != nil {
				return MenuEvent{}, textErr
			}
		}
		return MenuEvent{Selection: &Selection{mapping: cloneMenuMapping(mapping), confirmed: confirmed, textInput: text}}, nil
	}
	return MenuEvent{}, errors.New("whichkey: popup selection missing from retained snapshot")
}

// readPopupResult is isolated historical Completion-wire compatibility for
// unchanged tests. The public lifecycle uses readNativePopupResult instead.
func readPopupResult(directory, capability string, config Config) (message Completion, outcome error) {
	root, err := openPopupDirectory(directory)
	if err != nil {
		return Completion{}, err
	}
	defer func() { outcome = errors.Join(outcome, root.Close()) }()
	file, err := root.OpenFile("result.json", os.O_RDONLY|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if err != nil {
		return Completion{}, errors.New("whichkey: cannot open popup result")
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	if statErr := checkPopupSlot(file, false); statErr != nil {
		return Completion{}, errRuntimeStartup
	}
	data, err := io.ReadAll(io.LimitReader(file, popupMessageLimit+1))
	if err != nil || len(data) > popupMessageLimit {
		return Completion{}, errors.New("whichkey: popup result exceeds transport bound")
	}
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		return Completion{Version: CompletionVersion, Capability: capability, Kind: CompletionCancel}, nil
	}
	plain, textInput, err := nativeTextSelection(data, config.textEntry)
	if err != nil {
		return Completion{}, err
	}
	// Wrap the raw selection for the authoritative strict Completion decoder;
	// decoding CompletionSelection alone would accept missing false flags.
	wrapped, err := json.Marshal(struct {
		Version    int             `json:"version"`
		Capability string          `json:"capability"`
		Kind       string          `json:"kind"`
		Selection  json.RawMessage `json:"selection"`
	}{CompletionVersion, capability, CompletionComplete, plain})
	if err != nil || json.Unmarshal(wrapped, &message) != nil {
		return Completion{}, errors.New("whichkey: popup returned an invalid retained selection")
	}
	if message.Selection != nil {
		message.Selection.Text = textInput
	}
	if err := message.ValidateSelection(config.Spec); err != nil {
		return Completion{}, errors.New("whichkey: popup returned an invalid retained selection")
	}
	message.textInput = textInput
	return message, nil
}

// RunNativePopupTerminal has no interaction deadline and requires the foreground
// parent's private lifetime pipe plus its verified image/snapshot transport.
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
	if statErr := checkPopupSlot(lease, true); statErr != nil {
		return errors.New("whichkey: invalid native popup lifetime pipe")
	}
	ctx, cancel := context.WithCancelCause(parent)
	defer cancel(nil)
	watchCtx, cancelWatch := context.WithCancel(ctx)
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
			case <-watchCtx.Done():
				return
			case <-ticker.C:
			}
		}
	}()
	stopWatcher := func() {
		cancelWatch()
		<-stopped
	}
	defer stopWatcher()
	return runPopupTerminal(ctx, directory, noColor, ascii, stopWatcher)
}

func runPopupTerminal(parent context.Context, directory string, noColor, ascii bool, stopWatcher func()) (outcome error) {
	root, invocation, err := openVerifiedPopupDirectory(parent, directory)
	if err != nil {
		return err
	}
	defer func() { outcome = errors.Join(outcome, root.Close()) }()
	input, err := root.OpenFile("config.json", os.O_RDONLY|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if err != nil {
		return errors.New("whichkey: cannot open popup configuration")
	}
	if statErr := checkPopupSlot(input, false); statErr != nil {
		return errors.Join(errRuntimeStartup, input.Close())
	}
	config, decodeErr := decodePopupConfig(input)
	if err := errors.Join(decodeErr, input.Close()); err != nil {
		return err
	}
	config.textEntry, err = readNativeTextSpec(root)
	if err != nil {
		return err
	}
	// Presence is explicit: an absent/bad popup selection cannot become a
	// standalone menu's initial environment/PATH capture.
	ctx := context.WithValue(parent, terminalRuntimeKey{}, terminalRuntime{invocation.CLI, invocation.Endpoint})
	ctx = context.WithValue(ctx, terminalRenameKey{}, terminalRenameTransport{root, invocation.Invocation})
	// The layer already carries the requested position. Anchor popup content
	// at its top rather than centering it again inside the fixed rectangle.
	// This local copy does not change standalone RunTerminal positioning.
	config.Position = "top"
	event, err := RunTerminal(ctx, config, noColor, ascii)
	stopWatcher() // No lifetime reader survives terminal teardown/result commit.
	if err = errors.Join(err, context.Cause(parent)); err != nil {
		return err
	}
	return commitPopupResult(ctx, root, event, config)
}
