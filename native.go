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
	"strings"
	"time"
	"unicode/utf8"
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
		return nil, fmt.Errorf("whichkey: cannot read %s action metadata; check the captured endpoint, app connection and matching Rex CLI (no activation attempted)", stage)
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

func launchNativePopup(parent context.Context, origin launchOrigin, executable string) error {
	config, err := nativeDefaults()
	if err != nil {
		return errors.New("whichkey: invalid built-in menu configuration")
	}
	actions, err := origin.actions(parent, executable, true)
	if err != nil {
		return err
	}
	for _, mapping := range config.Spec {
		if mapping.Action != "" && actions[mapping.Action] != "client" {
			// Only Go-owned names, not arbitrary registry/callback output.
			return fmt.Errorf("whichkey: basic menu requires native app action %s; use a compatible Rex app or a configured --action", mapping.Action)
		}
	}
	data, err := json.Marshal(config)
	if err != nil {
		return errors.New("whichkey: cannot encode built-in menu configuration")
	}
	capability, err := NewInvocationCapability()
	if err != nil {
		return err
	}
	request := popupRequest{Configuration: string(data), Capability: capability,
		Server: origin.Server, SessionID: origin.SessionID, BlockID: origin.BlockID,
		ClientID: origin.ClientID, ConnectedAt: origin.ConnectedAt}
	// The parent alone retains this immutable snapshot. The child returns only
	// canonical keys/confirmation, never a replacement action or arguments.
	return runPopupWithLifetime(parent, request, config, executable, popupNative, func(ctx context.Context, message Completion) error {
		if err := message.ValidateSelection(config.Spec); err != nil {
			return errors.New("whichkey: invalid native retained selection")
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
			args := mapping.Args
			if len(args) == 0 {
				args = json.RawMessage(`{}`)
			}
			command := exec.CommandContext(ctx, executable,
				"--server", origin.Server, "--client", origin.ClientID, "--autostart=false", "--timeout", "3s",
				"do", mapping.Action, "--session", origin.SessionID, "--block", origin.BlockID, "--args", "-")
			command.Stdin = bytes.NewReader(args)
			output := &launchOutput{remaining: popupMessageLimit}
			command.Stdout, command.Stderr = output, output
			command.WaitDelay = 250 * time.Millisecond
			if err := command.Run(); err != nil {
				return errors.New("whichkey: native action failed or is uncertain; check app Remote Control and action compatibility; not retried")
			}
			return nil
		}
		return errors.New("whichkey: native selection is missing from the retained snapshot")
	})
}
