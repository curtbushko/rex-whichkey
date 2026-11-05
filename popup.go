package whichkey

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
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
	stage, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return prepareRuntimeInput(stage, input, output, server)
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
	Receipt       string `json:"receipt"`
	ReceiptHash   string `json:"receipt_hash"`
	Workspace     string `json:"workspace"`
	image         *runtimeImage
	retained      *runtimeReceipt
	selection     *rex.CLISelection
}

func decodePopupRequest(input io.Reader) (popupRequest, Config, error) {
	data, err := io.ReadAll(io.LimitReader(input, (8<<20)+1))
	if err != nil || !validPopupHandoff(data) {
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
		// Retain only this structurally decoded request so the selected controller
		// can authenticate and remove its own expired workspace, never run it.
		return request, Config{}, errors.New("whichkey: invalid or expired popup lifetime")
	}
	config, err := DecodeConfig(strings.NewReader(request.Configuration))
	if err != nil {
		return popupRequest{}, Config{}, errors.New("whichkey: invalid popup configuration")
	}
	return request, config, nil
}

func validPopupHandoff(data []byte) bool {
	return len(data) <= runtimeLimit && utf8.Valid(data) && json.Valid(data) && runtimeUnicode(data)
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
	block, err := r.resourceOriginBlock(ctx, client)
	if err != nil {
		return rex.Block{}, err
	}
	if r.retained != nil {
		if placementErr := r.retained.checkPlacement(block); placementErr != nil {
			return rex.Block{}, placementErr
		}
	}
	return block, nil
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
func (r popupRequest) complete(ctx context.Context, selection *rex.CLISelection, message Completion) error {
	if r.retained == nil || r.ReceiptHash == "" {
		return errRuntimeCompatibility
	}
	if err := r.retained.check(ctx, true); err != nil {
		return err
	}
	if err := r.retained.checkLoadedConsumer(ctx); err != nil {
		return err
	}
	message.Version, message.ReceiptHash = CompletionReceiptVersion, r.ReceiptHash
	if message.textInput != nil {
		if message.Selection == nil || message.Selection.Text == nil || *message.Selection.Text != *message.textInput {
			return errors.New("whichkey: rename result is missing validated name data")
		}
		message.textInput = nil // Only the typed selection text crosses the Lua bridge.
	}
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
	// --client selects the GUI for any nested configuration-action queues; Rex
	// still gives this fresh CLI activation its own actual caller identity.
	origin := launchOrigin{Server: r.Server, SessionID: r.SessionID, BlockID: r.BlockID, ClientID: r.ClientID}
	if err := runtimeAction(ctx, selection, origin, r.Completion, arguments, &launchOutput{remaining: popupMessageLimit}); err != nil {
		return errors.New("whichkey: completion failed or uncertain; not retried")
	}
	return nil
}

// RunPopupController is the finite key-launch adapter. rex.exec.start can wait
// for this command's bounded exit with h:wait; the actual controller is detached
// by Go and acknowledges its own input/checks on a private inherited pipe.
func RunPopupController(parent context.Context, input io.Reader, output io.Writer) (outcome error) {
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()
	if inherited, ok := input.(*os.File); ok {
		pipe, err := runtimePollablePipe(inherited)
		if err != nil {
			return err
		}
		deadline, _ := ctx.Deadline()
		if err := pipe.SetReadDeadline(deadline); err != nil {
			return errors.Join(errRuntimeStartup, pipe.Close())
		}
		input = pipe
	}
	request, _, err := decodePopupRequest(input)
	var inputErr error
	if closer, ok := input.(io.Closer); ok {
		inputErr = closer.Close()
	}
	// Like the child, authenticate a retained expired request for local cleanup.
	// validateRuntime adopts the owned image before checking live authority;
	// decode failure still rejects startup, and unauthenticated input owns nothing.
	validationErr := errors.Join(err, inputErr, request.validateRuntime(ctx))
	transferred := false
	defer func() {
		if request.image != nil {
			outcome = errors.Join(outcome, request.image.retire())
			outcome = errors.Join(outcome, cleanupPublicRuntime(ctx, request.Workspace, transferred, &request))
		}
	}()
	if validationErr != nil {
		return errRuntimeStartup
	}
	// This finite launcher's source and stdin must be retired before the child
	// can acknowledge startup. Never leave deferred close errors after effects.
	if retireErr := request.image.retire(); retireErr != nil {
		return retireErr
	}
	transferred, err = startRuntimeRequest(ctx, request, func() error { return reportRuntimeLauncher(output) })
	return err
}

// RunPopupControllerChild runs once outside the PTY. The selected Go launcher
// supplies finite stdin, the sole fd-3 startup writer and fd-4 proceed reader,
// never Rex API handles.
func RunPopupControllerChild(parent context.Context, input, startup, proceed *os.File) (outcome error) {
	input, startup, proceed, err := popupControllerPipes(input, startup, proceed)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(parent, 35*time.Second)
	defer cancel()
	stage, stopStage := context.WithTimeout(ctx, 3*time.Second)
	defer stopStage()
	defer func() {
		outcome = errors.Join(outcome, rejectRuntimeStartup(stage, &startup), closeRuntimePipe(&proceed))
	}()
	deadline, _ := stage.Deadline()
	if deadlineErr := input.SetReadDeadline(deadline); deadlineErr != nil {
		return errors.Join(errRuntimeStartup, input.Close())
	}
	request, config, decodeErr := decodePopupRequest(input)
	decodeErr = errors.Join(decodeErr, input.Close())
	validationErr := errors.Join(decodeErr, request.validateRuntime(stage))
	enteredLifecycle := false
	released := false
	release := func() error {
		if released || request.image == nil {
			return nil
		}
		released = true // Never repeat uncertain private cleanup.
		return errors.Join(request.image.file.Close(), removeRuntimeWorkspace(request.Workspace))
	}
	defer func() {
		// Report a known pre-popup failure before potentially yielding cleanup.
		outcome = errors.Join(outcome, rejectRuntimeStartup(stage, &startup))
		cleanupErr := release()
		outcome = errors.Join(outcome, cleanupErr)
		if request.image != nil && cleanupErr == nil && validationErr == nil && !enteredLifecycle {
			outcome = errors.Join(outcome, cancelPreparedRuntime(ctx, request))
		}
	}()
	if validationErr != nil {
		return validationErr
	}
	selection := request.retained.CLI
	config, err = prepareConfiguredRename(stage, request, config, selection)
	if err != nil {
		return err
	}
	ack := startup
	startup = nil // The finite startup guard owns this endpoint, even on failure.
	if startupErr := acknowledgePopupStartup(stage, request, ack); startupErr != nil {
		return startupErr
	}
	gate := proceed
	proceed = nil // waitRuntimeProceed owns the last startup transport close.
	if proceedErr := waitRuntimeProceed(stage, gate); proceedErr != nil {
		return proceedErr
	}
	stopStage()
	enteredLifecycle = true
	return runPopupLifecycle(ctx, request, config, selection, func(ctx context.Context, message Completion) error {
		// The runtime workspace also carries private preparation data. Remove
		// it, not just the terminal transport, before any retained callback.
		if cleanupErr := release(); cleanupErr != nil {
			return cleanupErr
		}
		return request.complete(ctx, selection, message)
	})
}

// runPopupLifecycle always enforces the configured handoff's original expiry.
// Missing/zero expiry cannot select the native foreground policy.
func runPopupLifecycle(parent context.Context, request popupRequest, config Config, selection *rex.CLISelection,
	finish func(context.Context, Completion) error,
) error {
	return runPopupWithLifetime(parent, request, config, selection, popupConfigured, finish)
}

type popupLifetime uint8

const (
	popupConfigured popupLifetime = iota
	popupNative
)

// runPopupWithLifetime shares owned transport/teardown, not lifetime authority.
// Only the Go-native caller supplies popupNative; decoded handoffs cannot do so.
// finish receives only a retained selection, after cleanup, never child actions.
func runPopupWithLifetime(parent context.Context, request popupRequest, config Config, selection *rex.CLISelection,
	lifetime popupLifetime, finish func(context.Context, Completion) error,
) (outcome error) {
	ctx, stop := signal.NotifyContext(parent, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()
	request.selection = selection
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
	client, err := rex.NewContext(ctx, rex.Options{Selection: selection, Server: request.Server,
		Session: request.SessionID, Timeout: 3 * time.Second})
	if err != nil {
		return errors.New("whichkey: origin CLI/endpoint compatibility failed; check the execution-host integration and relaunch from the original terminal; not retried")
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
	if imageErr := request.preparePopupImage(ctx, directory); imageErr != nil {
		return imageErr
	}
	command := []string{filepath.Join(directory, runtimeImageName), "popup", "--directory", directory}
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
	ready, err := request.popupCreationReady(ctx, lifetime, uiDeadline)
	if err != nil || !ready {
		return err
	}
	layer, err = origin.NewLayerWithSize(ctx, initialBounds, command, initialSize)
	if err != nil {
		return errors.New("whichkey: popup creation failed or uncertain; not retried")
	}
	// Keep the initial rectangle for this invocation. Navigation and redraws
	// adapt content to the terminal cells, never change owned layer bounds.
	err = waitPopup(ctx, request, selection, layer.BlockID, originBlock)
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
	message, err := readPopupResult(directory, request.Capability, config)
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

// Geometry and rename preparation yield. Recheck the exact loaded registry at
// the last read-only boundary; sampling must not spend the teardown reserve.
func (r popupRequest) popupCreationReady(ctx context.Context, lifetime popupLifetime, deadline time.Time) (bool, error) {
	if lifetime == popupConfigured && !deadline.After(time.Now()) {
		return false, nil
	}
	if r.retained != nil {
		if err := r.retained.checkOrigin(ctx); err != nil {
			return false, err
		}
		if err := r.retained.checkLoadedConsumer(ctx); err != nil {
			return false, err
		}
	}
	return lifetime != popupConfigured || deadline.After(time.Now()), nil
}

// waitPopup deliberately does not use Layer.Wait: that wrapper caps the whole
// CLI process at its RPC timeout. `rex wait --help` documents unbounded event
// delivery, with --timeout still bounding connections/nonstreaming requests.
// Watch captured resource liveness too, plus app generation only in configured
// mode. Native display survives GUI ambiguity/disconnection. Neither policy
// checks active-tab visibility: switching tabs must not cancel or refocus it.
func waitPopup(ctx context.Context, request popupRequest, selection *rex.CLISelection, blockID string,
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
				if err == nil && request.retained != nil {
					err = request.retained.checkLoadedConsumer(check)
				}
				done()
				if err != nil {
					watcherErr = err
					cancel(err)
					return
				}
			}
		}
	}()
	arguments := []string{"--server", request.Server, runtimeNoAutostart, runtimeTimeoutFlag, runtimeStageTimeout,
		"wait", "--session", request.SessionID, "--", blockID}
	err := runtimeCommand(waitCtx, selection, launchOrigin{Server: request.Server, SessionID: request.SessionID, BlockID: request.BlockID}, arguments, nil, &launchOutput{remaining: popupMessageLimit})
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
	for _, name := range []string{"config.json", "result.json", "lifetime", nativeTextFile, runtimeImageName, "runtime.json"} {
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
	if _, statErr := runtimeSlotStat(file, "result.json", false); statErr != nil {
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
	if _, statErr := runtimeSlotStat(lease, "lifetime", true); statErr != nil {
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
	root, selection, err := openVerifiedPopupDirectory(parent, directory)
	if err != nil {
		return err
	}
	defer func() { outcome = errors.Join(outcome, root.Close()) }()
	input, err := root.OpenFile("config.json", os.O_RDONLY|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if err != nil {
		return errors.New("whichkey: cannot open popup configuration")
	}
	if _, statErr := runtimeSlotStat(input, "config.json", false); statErr != nil {
		return errors.Join(errRuntimeStartup, input.Close())
	}
	config, decodeErr := DecodeConfig(input)
	if err := errors.Join(decodeErr, input.Close()); err != nil {
		return err
	}
	config.textEntry, err = readNativeTextSpec(root)
	if err != nil {
		return err
	}
	// Presence is explicit: an absent/bad popup selection cannot become a
	// standalone menu's initial environment/PATH capture.
	ctx := context.WithValue(parent, terminalRuntimeKey{}, terminalRuntime{selection, selection.Capabilities.Endpoint})
	if !deadline.IsZero() {
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(ctx, deadline)
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
	return commitPopupResult(ctx, root, event, config)
}
