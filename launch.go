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
// Lua compares it with the actual fresh action ctx; it must never replace ctx.
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
// Read-only metadata selects the route BEFORE any activation. An uncertain
// activation never switches routes or retries.
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
		return errors.New("whichkey: --client or REX_CLIENT must be an exact full client ID for an attached live graphical app, not a label or suffix")
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
	// Absence of REX_CLIENT is normal: the shared resolver counts only live
	// graphical apps attached to this session, excluding CLI/automation clients.
	origin.ClientID, origin.ConnectedAt, err = popupOriginConnection(ctx, client, origin.SessionID, requested)
	if err != nil {
		return err
	}
	// Reuse the controller's exact placement and generation checks. A wrong
	// endpoint, stale session/block or reconnect fails before action activation.
	request := popupRequest{Server: server, SessionID: origin.SessionID, BlockID: origin.BlockID,
		ClientID: origin.ClientID, ConnectedAt: origin.ConnectedAt}
	if _, err := request.originBlock(ctx, client); err != nil {
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
		// Use the caller context, not the short preflight deadline, for the
		// invocation's own finite UI lifetime. The captured origin is rechecked.
		return launchNativePopup(parent, origin, executable)
	} else if kind != "lua" {
		return errors.New("whichkey: launch name is not a server configuration action; choose a registered menu.register_launch action")
	}
	if _, err := request.originBlock(ctx, client); err != nil {
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
	// Count and discard all child output, including arbitrary user-code errors.
	output := &launchOutput{remaining: popupMessageLimit}
	command.Stdout, command.Stderr = output, output
	command.WaitDelay = 250 * time.Millisecond
	if err := command.Run(); err != nil {
		if activation.Err() != nil {
			return errors.New("whichkey: registered launch activation timed out or was cancelled; result uncertain, not retried and not switched to native mode")
		}
		return errors.New("whichkey: registered launch action failed or is uncertain; check its setup/completion registration, server-host helper/modules and app Remote Control; private action diagnostics suppressed, no retry or native fallback")
	}
	return nil
}

// No handoff or arbitrary process diagnostics are retained or displayed.
type launchOutput struct{ remaining int }

func (output *launchOutput) Write(data []byte) (int, error) {
	if len(data) > output.remaining {
		return 0, io.ErrShortWrite
	}
	output.remaining -= len(data)
	return len(data), nil
}
