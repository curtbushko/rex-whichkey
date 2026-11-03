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
	"strings"
	"syscall"
	"time"
	"unicode/utf8"

	"temp/rex"
)

// actions reads metadata only. Server registration must be queried without an
// inherited client selector: `rex actions` and `rex --client ID actions` expose
// different registries. See docs/launch-contract.md for the captured contracts.
func (origin launchOrigin) actions(parent context.Context, executable string, client bool) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()
	argv := []string{"--server", origin.Server, "--autostart=false", "--timeout", "3s"}
	stage := "server configuration"
	if client {
		argv = append(argv, "--client", origin.ClientID)
		stage = "selected app native"
	}
	argv = append(argv, "actions", "--json")
	command := exec.CommandContext(ctx, executable, argv...)
	for _, entry := range os.Environ() {
		name, _, _ := strings.Cut(entry, "=")
		switch name {
		case "REX_SERVER", "REX_CLIENT", "REX_SESSION", "REX_WINDOW", "REX_BLOCK":
			continue
		}
		command.Env = append(command.Env, entry)
	}
	// These are captured identities, not ambient focus defaults. Actions metadata
	// itself has no session/block flags and performs no mutation.
	command.Env = append(command.Env, "REX_SESSION="+origin.SessionID, "REX_BLOCK="+origin.BlockID)
	output := &launchBuffer{remaining: 1 << 20}
	command.Stdout = output
	command.Stderr = &launchOutput{remaining: popupMessageLimit}
	command.WaitDelay = 250 * time.Millisecond
	if err := command.Run(); err != nil {
		return nil, fmt.Errorf("whichkey: cannot read %s action metadata; check the captured endpoint and matching Rex CLI (no activation attempted)", stage)
	}
	data := output.buffer.Bytes()
	fields, err := completionObject(data, "actions")
	if err != nil || !utf8.Valid(data) {
		return nil, fmt.Errorf("whichkey: invalid %s action metadata; a compatible Rex actions --json response is required", stage)
	}
	var entries []struct {
		Name string `json:"name"`
		Kind string `json:"kind"`
	}
	if raw, present := fields["actions"]; !present || bytes.Equal(bytes.TrimSpace(raw), []byte("null")) ||
		json.Unmarshal(raw, &entries) != nil || len(entries) > 4096 {
		return nil, fmt.Errorf("whichkey: invalid %s action list", stage)
	}
	actions := make(map[string]string, len(entries))
	for _, entry := range entries {
		if strings.TrimSpace(entry.Name) == "" || strings.TrimSpace(entry.Kind) == "" ||
			validateText(entry.Name) != nil || validateText(entry.Kind) != nil {
			return nil, fmt.Errorf("whichkey: malformed %s action identity", stage)
		}
		if _, duplicate := actions[entry.Name]; duplicate {
			return nil, fmt.Errorf("whichkey: duplicate %s action identity", stage)
		}
		actions[entry.Name] = entry.Kind
	}
	return actions, nil
}

// launchBuffer bounds retained metadata; unlike launchOutput it keeps stdout.
// It is never used for capability-bearing or arbitrary configured-action output.
type launchBuffer struct {
	buffer    bytes.Buffer
	remaining int
}

func (output *launchBuffer) Write(data []byte) (int, error) {
	if len(data) > output.remaining {
		return 0, io.ErrShortWrite
	}
	output.remaining -= len(data)
	return output.buffer.Write(data)
}

// launchNativePopup retains requested as syntax-validated Go-owned intent, not
// a connection generation or child authority. No app is consulted until selection.
func launchNativePopup(parent context.Context, origin launchOrigin, executable, requested string) error {
	// Keep signal ownership across interaction, cleanup and metadata/dispatch
	// gaps. A context owned only by the popup would leave those gaps exposed.
	ctx, stop := signal.NotifyContext(parent, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()
	config, err := nativeDefaults()
	if err != nil {
		return errors.New("whichkey: invalid built-in menu configuration")
	}
	client, err := rex.New(rex.Options{Executable: executable, Server: origin.Server,
		Session: origin.SessionID, Timeout: 3 * time.Second})
	if err != nil {
		return errors.New("whichkey: cannot construct the captured native endpoint/session client")
	}
	// Capture resource placement independently of GUI identity. Resource closes
	// must never resolve a different tab after the user has confirmed the action.
	inspect, finishInspect := context.WithTimeout(ctx, 3*time.Second)
	request := popupRequest{Server: origin.Server, SessionID: origin.SessionID, BlockID: origin.BlockID}
	placement, err := request.resourceOriginBlock(inspect, client)
	finishInspect()
	if err != nil {
		return err
	}
	nameContext, finishName := context.WithTimeout(ctx, 3*time.Second)
	window, err := placement.Session().Window(nameContext, placement.WindowID)
	finishName()
	if err != nil {
		return errors.New("whichkey: cannot inspect the captured tab name")
	}
	config.textEntry = &textEntrySpec{Initial: window.Label}
	if len(config.textEntry.Initial) > renameTextLimit {
		config.textEntry.Initial = "" // Require a new name, never silently truncate.
	}
	message, err := nativePopupSelection(ctx, origin, executable, config)
	if err != nil {
		return err
	}
	if message.Kind == CompletionCancel {
		return nil
	}
	for _, mapping := range config.Spec {
		if mapping.Keys != message.Selection.Keys {
			continue
		}
		if mapping.CallbackRef != "" || mapping.Action == "" || !mapping.Enabled() {
			return errors.New("whichkey: native selection is not a retained native action")
		}
		if mapping.Keys == "tr" && mapping.Action == "client.tab.rename" {
			if err := message.ValidateSelection(config.Spec); err != nil || message.textInput == nil || config.textEntry == nil {
				return errors.New("whichkey: tab rename requires validated native text input")
			}
			return nativeTabRename(ctx, client, origin, placement, requested, *message.textInput)
		}
		// These operations already have an exact server resource target. They
		// need confirmation, not a GUI app or its native action registration.
		if mapping.Action == "pane.close" || mapping.Action == "client.tab.close" || mapping.Action == "session.close" {
			if err := message.ValidateSelection(config.Spec); err != nil {
				return errors.New("whichkey: invalid native retained close selection")
			}
			return nativeResourceClose(ctx, client, origin, placement, requested, mapping)
		}
		// Every app action uses the same pane-owner binding. Do not reintroduce
		// per-action exceptions that fall back to a session-wide app chooser.
		// The original UI/files/lease are already gone before this lookup.
		connection, err := nativePaneAppTarget(ctx, client, origin, placement, requested)
		if err != nil {
			return fmt.Errorf("whichkey: native action targeting refused after popup cleanup: %w", err)
		}
		// User interaction is over and the popup lifecycle has returned
		// successfully. Start a fresh short budget, still signal-cancellable.
		dispatch, done := context.WithTimeout(ctx, 3*time.Second)
		defer done()
		if err := message.ValidateSelection(config.Spec); err != nil {
			return errors.New("whichkey: invalid native retained selection after target choice")
		}
		target := origin
		target.ClientID, target.ConnectedAt = connection.ID, connection.ConnectedAt
		actions, err := target.actions(dispatch, executable, true)
		if err != nil {
			return fmt.Errorf("whichkey: native action targeting refused after popup cleanup: %w", err)
		}
		dispatchAction := mapping.Action
		if mapping.Action == "client.tab.new" {
			// Compile only retained tab intent; the child cannot supply this
			// action or its exact session/window arguments.
			dispatchAction = "session.select"
		}
		if actions[dispatchAction] != "client" {
			// Check the actual action being sent, not the translated intent's
			// name or unrelated defaults on the selected app.
			return fmt.Errorf("whichkey: native action targeting refused after popup cleanup: selected app does not provide native action %s", dispatchAction)
		}
		bound := popupRequest{Server: origin.Server, SessionID: origin.SessionID, BlockID: origin.BlockID,
			ClientID: target.ClientID, ConnectedAt: target.ConnectedAt}
		if mapping.Action == "client.tab.new" {
			return nativeNewTab(dispatch, client, bound, placement, requested, executable, mapping)
		}
		args := mapping.Args
		if len(args) == 0 {
			args = json.RawMessage(`{}`)
		}
		// Keep the transport unscoped so context is explicit in the request,
		// never injected by the adapter or taken from ambient CLI context.
		performClient, err := rex.New(rex.Options{Executable: executable, Server: target.Server,
			Timeout: 3 * time.Second})
		if err != nil {
			return errors.New("whichkey: native action targeting refused: cannot construct the captured endpoint client")
		}
		request := map[string]any{
			"client_id": target.ClientID, "action": mapping.Action, "args": args,
			"origin": "cli", "timeout_ms": 3000,
		}
		if mapping.Action != "pane.split" {
			// Preserve origin-bound context, especially for destructive actions.
			// Splits instead use the chosen app's selected tab and focused pane;
			// their origin authorizes invocation, not the destination.
			request["session_id"], request["block_id"] = origin.SessionID, origin.BlockID
		}
		// Recheck the exact captured generation and origin immediately before
		// one dispatch. A reconnect never renews or silently selects another app.
		if err := validatePaneAppOrigin(dispatch, client, bound, placement, requested); err != nil {
			return err
		}
		// Discard only a validated JSON result. Errors can follow a successful
		// mutation, so neither an uncertain reply nor a timeout permits retry.
		if err := performClient.Call(dispatch, "client.perform", request, nil); err != nil {
			return errors.New("whichkey: native action failed or is uncertain; check app Remote Control and action compatibility; not retried")
		}
		return nil
	}
	return errors.New("whichkey: native selection is missing from the retained snapshot")
}

// nativeNewTab compiles only Go-owned tn intent after popup cleanup and the
// caller's session.select registration check. Creation and app activation are
// distinct effects, never a retryable transaction.
func nativeNewTab(ctx context.Context, client *rex.Client, bound popupRequest, placement rex.Block, requested, executable string, mapping Mapping) error {
	if mapping.Action != "client.tab.new" || mapping.Keys != "tn" {
		return errors.New("whichkey: native new tab refused: invalid retained new-tab intent")
	}
	if len(mapping.Args) != 0 {
		if _, err := completionObject(mapping.Args); err != nil {
			return errors.New("whichkey: native new tab refused: retained new-tab arguments must be empty")
		}
	}
	unscoped, err := rex.New(rex.Options{Executable: executable, Server: bound.Server,
		Timeout: 3 * time.Second})
	if err != nil {
		return errors.New("whichkey: native new tab refused: cannot construct the captured endpoint client")
	}
	session, err := client.Session(ctx, bound.SessionID)
	if err != nil {
		return errors.New("whichkey: native new tab refused: cannot inspect the captured origin session")
	}
	if err := validatePaneAppOrigin(ctx, client, bound, placement, requested); err != nil {
		return err
	}
	// Exactly one default-terminal creation in the captured source session.
	// NewWindow can retain a validated ID even when later response checks fail.
	window, err := session.NewWindow(ctx, "")
	if err != nil {
		if window.ID != "" {
			return fmt.Errorf("whichkey: native new tab creation response is incomplete; inspect preserved tab %s in session %s manually; activation not attempted, no retry or rollback", window.ID, bound.SessionID)
		}
		return errors.New("whichkey: native new tab creation failed or is uncertain; inspect the origin session for a possible new tab before any manual action; no retry or rollback")
	}
	// Freeze this exact returned ID. Never create again, close it on failure,
	// resolve a position, or fall back to the original client.tab.new action.
	target := launchOrigin{Server: bound.Server, SessionID: bound.SessionID,
		BlockID: bound.BlockID, ClientID: bound.ClientID}
	actions, err := target.actions(ctx, executable, true)
	if err != nil || actions["session.select"] != "client" {
		return fmt.Errorf("whichkey: native new tab %s preserved in session %s; activation refused because session.select registration could not be verified; select that tab manually, no retry or rollback", window.ID, bound.SessionID)
	}
	if _, err := window.Inspect(ctx); err != nil {
		return fmt.Errorf("whichkey: native new tab %s created in session %s; membership could not be verified before activation; inspect that exact tab manually, no retry or rollback", window.ID, bound.SessionID)
	}
	// Creation can legitimately change GUI ownership. Keep the already-bound
	// app; revalidate its generation and origin placement, never pick a new owner.
	current, err := bound.originBlock(ctx, client)
	if err != nil || current.WindowID != placement.WindowID || current.LayerID != placement.LayerID {
		return fmt.Errorf("whichkey: native new tab %s preserved in session %s; origin or chosen app connection changed before activation; select that tab manually, no retry or rollback", window.ID, bound.SessionID)
	}
	// Only parent-derived destination args select the new tab. No outer
	// session/block override can redirect the chosen app away from that ID.
	if err := unscoped.Call(ctx, "client.perform", map[string]any{
		"client_id": bound.ClientID, "action": "session.select",
		"args":   map[string]string{"session_id": bound.SessionID, "window_id": window.ID},
		"origin": "cli", "timeout_ms": 3000,
	}, nil); err != nil {
		return fmt.Errorf("whichkey: native new tab %s preserved in session %s; activation failed or is uncertain; check app Remote Control/action compatibility and select that tab manually, no retry or rollback", window.ID, bound.SessionID)
	}
	return nil
}

// nativePopupSelection only captures finish's validated result. No app lookup,
// action, or second interactive UI can run inside finish's short operation context.
// Returning success also requires the lifecycle's deferred cleanup to succeed.
func nativePopupSelection(ctx context.Context, origin launchOrigin, executable string, config Config) (Completion, error) {
	data, err := json.Marshal(config)
	if err != nil {
		return Completion{}, errors.New("whichkey: cannot encode native menu configuration")
	}
	capability, err := NewInvocationCapability()
	if err != nil {
		return Completion{}, err
	}
	request := popupRequest{Configuration: string(data), Capability: capability,
		Server: origin.Server, SessionID: origin.SessionID, BlockID: origin.BlockID}
	var retained Completion
	err = runPopupWithLifetime(ctx, request, config, executable, popupNative, func(_ context.Context, message Completion) error {
		if err := message.ValidateSelection(config.Spec); err != nil {
			return errors.New("whichkey: invalid native retained selection")
		}
		retained = message
		return nil
	})
	if err != nil {
		return Completion{}, err
	}
	if err := retained.ValidateSelection(config.Spec); err != nil {
		return Completion{}, errors.New("whichkey: native popup returned no valid retained result")
	}
	return retained, nil
}
