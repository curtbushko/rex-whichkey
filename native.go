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
	// Keep signal ownership across both UI stages, their cleanup, and the gaps
	// spent looking up targets/metadata. Per-popup contexts alone leave gaps.
	ctx, stop := signal.NotifyContext(parent, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()
	config, err := nativeDefaults()
	if err != nil {
		return errors.New("whichkey: invalid built-in menu configuration")
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
		// The original UI/files/lease are gone before any app lookup. Choice
		// changes only the target; mapping and message remain the original intent.
		client, err := rex.New(rex.Options{Executable: executable, Server: origin.Server,
			Session: origin.SessionID, Timeout: 3 * time.Second})
		if err != nil {
			return errors.New("whichkey: native action targeting refused: cannot construct the captured endpoint/session client")
		}
		connection, err := nativeActionTarget(ctx, client, origin, executable, requested, mapping.Desc)
		if err != nil {
			return fmt.Errorf("whichkey: native action targeting refused after popup cleanup: %w", err)
		}
		if connection.ID == "" {
			return nil // Target chooser cancellation never dispatches the original action.
		}
		// User interaction is over and both popup lifecycles have returned
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
		if actions[mapping.Action] != "client" {
			// Only this retained Go-owned action matters; other defaults
			// need not be available on the selected app.
			return fmt.Errorf("whichkey: native action targeting refused after popup cleanup: selected app does not provide native action %s", mapping.Action)
		}
		args := mapping.Args
		if len(args) == 0 {
			args = json.RawMessage(`{}`)
		}
		command := exec.CommandContext(dispatch, executable,
			"--server", target.Server, "--client", target.ClientID, "--autostart=false", "--timeout", "3s",
			"do", mapping.Action, "--session", origin.SessionID, "--block", origin.BlockID, "--args", "-")
		command.Stdin = bytes.NewReader(args)
		output := &launchOutput{remaining: popupMessageLimit}
		command.Stdout, command.Stderr = output, output
		command.WaitDelay = 250 * time.Millisecond
		bound := popupRequest{Server: origin.Server, SessionID: origin.SessionID, BlockID: origin.BlockID,
			ClientID: target.ClientID, ConnectedAt: target.ConnectedAt}
		// Recheck the exact captured generation and origin immediately before
		// one dispatch. A reconnect never renews or silently selects another app.
		if _, err := bound.originBlock(dispatch, client); err != nil {
			return fmt.Errorf("whichkey: native action targeting refused after popup cleanup: %w", err)
		}
		if err := command.Run(); err != nil {
			return errors.New("whichkey: native action failed or is uncertain; check app Remote Control and action compatibility; not retried")
		}
		return nil
	}
	return errors.New("whichkey: native selection is missing from the retained snapshot")
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
