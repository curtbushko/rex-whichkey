package whichkey

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"temp/rex"
)

// LaunchOptions selects execution-host routing and an optional configured action.
// Nil selectors are absent; present empty selectors are invalid. An empty Action
// prefers registered whichkey, otherwise the native basic menu. A nonempty Action
// always requires that configured action; it never requests native fallback.
type LaunchOptions struct {
	Server *string
	Client *string
	Action string
}

// launchOrigin is Go-owned transport context, not a capability or mapping data.
// Lua compares session/block with the real action ctx and validates the GUI
// binding independently; ClientID is the target app, not the CLI caller.
type launchOrigin struct {
	Server      string `json:"server"`
	SessionID   string `json:"session_id"`
	BlockID     string `json:"block_id"`
	ClientID    string `json:"client_id"`
	ConnectedAt string `json:"connected_at"`
}

func (origin launchOrigin) marshal() ([]byte, error) {
	server, err := rex.CanonicalEndpoint(origin.Server)
	if err != nil || server != origin.Server ||
		!popupResourceID(origin.SessionID, "session") || !popupResourceID(origin.BlockID, "block") ||
		!popupResourceID(origin.ClientID, "client") || !popupConnectionGeneration(origin.ConnectedAt) {
		return nil, errors.New("whichkey: launch requires a canonical endpoint and exact live origin identities")
	}
	data, err := json.Marshal(origin)
	if err != nil || len(data) > popupMessageLimit {
		return nil, errors.New("whichkey: cannot encode launch origin")
	}
	return data, nil
}

func popupConnectionGeneration(value string) bool {
	return len(value) <= 1024 && utf8.ValidString(value) && strings.TrimSpace(value) != "" &&
		strings.IndexFunc(value, unicode.IsControl) < 0
}

// LaunchPopup opens the basic native popup when the implicit default action is
// unregistered, or activates a registered original-VM action once. Configured
// success acknowledges activation only; native success follows UI exit/cleanup.
// Read-only server metadata selects the route before GUI binding or activation.
// Native display needs no app; configured activation binds its app at launch.
// An uncertain activation never switches routes or retries.
func LaunchPopup(parent context.Context, options LaunchOptions) error {
	if parent == nil {
		return errors.New("whichkey: launch requires a context")
	}
	action := options.Action
	if action == "" {
		action = "whichkey"
	}
	if strings.Trim(action, "abcdefghijklmnopqrstuvwxyz0123456789_") != "" {
		return errors.New("whichkey: --action must name a registered configuration action (lowercase letters, digits and underscores only)")
	}
	origin := launchOrigin{SessionID: os.Getenv("REX_SESSION"), BlockID: os.Getenv("REX_BLOCK")}
	if !popupResourceID(origin.SessionID, "session") || !popupResourceID(origin.BlockID, "block") {
		return errors.New("whichkey: run inside the originating Rex terminal with full REX_SESSION and REX_BLOCK IDs; no focused-block fallback is available")
	}
	selector := options.Client
	if selector == nil {
		if value, present := os.LookupEnv("REX_CLIENT"); present {
			selector = &value
		}
	}
	if selector != nil && !popupResourceID(*selector, "client") {
		return errors.New("whichkey: --client or REX_CLIENT must be an exact full client:UUID, not an empty value, label or suffix")
	}
	ctx, cancel := context.WithTimeout(parent, 20*time.Second)
	defer cancel()
	server, err := rex.ResolveEndpoint(ctx, options.Server)
	if err != nil {
		return err
	}
	origin.Server = server
	executable, err := popupRexExecutable()
	if err != nil {
		return err
	}
	client, err := rex.New(rex.Options{Executable: executable, Server: server, Session: origin.SessionID, Timeout: 3 * time.Second})
	if err != nil {
		return errors.New("whichkey: cannot construct the explicitly targeted launch client")
	}
	requested := ""
	if selector != nil {
		requested = *selector
	}
	// Validate server resources before choosing a route, without inspecting GUI
	// attachments. A requested native selector is intent, not a live binding.
	request := popupRequest{Server: server, SessionID: origin.SessionID, BlockID: origin.BlockID}
	placement, err := request.resourceOriginBlock(ctx, client)
	if err != nil {
		return err
	}
	actions, err := origin.actions(ctx, executable, false)
	if err != nil {
		return err
	}
	if kind, registered := actions[action]; !registered {
		if options.Action != "" {
			return errors.New("whichkey: requested configuration launch action is not registered on the captured server; register it with menu.register_launch or omit --action for the basic menu")
		}
		// Native interaction has no idle deadline. Preserve only the optional
		// selector until an executable retained selection needs an app.
		return launchNativePopup(parent, origin, executable, requested)
	} else if kind != "lua" {
		return errors.New("whichkey: launch name is not a server configuration action; choose a registered menu.register_launch action")
	}
	// Configured launch still requires an exact live app and generation before
	// activation in the original Lua VM; no decoded handoff can defer binding.
	connection, err := nativePaneAppTarget(ctx, client, origin, placement, requested)
	if err != nil {
		return err
	}
	origin.ClientID, origin.ConnectedAt = connection.ID, connection.ConnectedAt
	request.ClientID, request.ConnectedAt = origin.ClientID, origin.ConnectedAt
	if err := validatePaneAppOrigin(ctx, client, request, placement, requested); err != nil {
		return err
	}
	data, err := origin.marshal()
	if err != nil {
		return err
	}
	arguments, err := json.Marshal(struct {
		Origin string `json:"origin"`
	}{string(data)})
	if err != nil {
		return errors.New("whichkey: cannot encode launch arguments")
	}
	activation, stop := context.WithTimeout(ctx, 10*time.Second)
	defer stop()
	command := exec.CommandContext(activation, executable,
		"--server", server, "--client", origin.ClientID, "--autostart=false", "--timeout", "10s",
		"do", action, "--session", origin.SessionID, "--block", origin.BlockID, "--args", "-")
	command.Stdin = bytes.NewReader(arguments)
	// Share one comparable writer: os/exec serializes stdout/stderr writes.
	// Discard arbitrary child output. Only recognize the fixed legacy mismatch
	// diagnostic to explain a stale Lua installation without exposing child text.
	output := &launchOutput{remaining: popupMessageLimit}
	command.Stdout, command.Stderr = output, output
	command.WaitDelay = 250 * time.Millisecond
	if err := command.Run(); err != nil {
		if activation.Err() != nil {
			return errors.New("whichkey: registered launch activation timed out or was cancelled; result uncertain, not retried and not switched to native mode")
		}
		if output.legacyClientBinding {
			return errors.New("whichkey: installed Lua module still equates the CLI caller with the GUI target; reinstall the matching Lua modules and reload Rex configuration; no retry or native fallback")
		}
		return errors.New("whichkey: registered launch action failed or is uncertain; check its setup/completion registration, server-host helper/modules and app Remote Control; private action diagnostics suppressed, no retry or native fallback")
	}
	return nil
}

// No handoff or arbitrary process diagnostics are retained or displayed.
type launchOutput struct {
	remaining           int
	legacyMatch         int
	legacyClientBinding bool
}

func (output *launchOutput) Write(data []byte) (int, error) {
	if len(data) > output.remaining {
		return 0, io.ErrShortWrite
	}
	output.remaining -= len(data)
	// Retain only a prefix length and a boolean, never process output (which
	// may contain capabilities). The literal has no overlapping proper prefix.
	const legacy = "rex_whichkey: launch origin client_id does not match the actual action context"
	for _, char := range data {
		if char != legacy[output.legacyMatch] {
			output.legacyMatch = 0
		}
		if char == legacy[output.legacyMatch] {
			output.legacyMatch++
			if output.legacyMatch == len(legacy) {
				output.legacyClientBinding = true
				output.legacyMatch = 0
			}
		}
	}
	return len(data), nil
}
