package whichkey

import (
	"context"
	"errors"
	"io"
	"os"
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
	selection   *rex.CLISelection
	image       *runtimeImage
}

func popupConnectionGeneration(value string) bool {
	return len(value) <= 1024 && utf8.ValidString(value) && strings.TrimSpace(value) != "" &&
		strings.IndexFunc(value, unicode.IsControl) < 0
}

// LaunchPopup opens the basic native popup when the implicit default action is
// unregistered, or activates a registered original-VM action once. Configured
// success acknowledges activation and finite startup checks, not popup display;
// native success follows UI exit/cleanup.
// Read-only server metadata selects the route before GUI binding or activation.
// Native display needs no app; configured activation binds its app at launch.
// An uncertain activation never switches routes or retries.
func LaunchPopup(parent context.Context, options LaunchOptions) (outcome error) {
	if parent == nil {
		return errors.New("whichkey: launch requires a context")
	}
	action, err := publicLaunchAction(options)
	if err != nil {
		return err
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
	image, err := captureSelf(ctx)
	if err != nil {
		return err
	}
	defer func() { outcome = errors.Join(outcome, image.retire()) }()
	origin.image = image
	selection, err := rex.ResolveCLI(ctx, rex.CLIOptions{})
	if err != nil {
		return err
	}
	server, err := rex.ResolveEndpoint(ctx, options.Server, selection)
	if err != nil {
		return err
	}
	origin.Server = server
	client, err := rex.NewContext(ctx, rex.Options{Selection: selection, Server: server, Session: origin.SessionID, Timeout: 3 * time.Second})
	if err != nil {
		return errors.Join(errRuntimeCompatibility, err)
	}
	selection = client.Selection()
	origin.selection = selection
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
	actions, err := origin.actions(ctx, selection, false)
	if err != nil {
		return err
	}
	if kind, registered := actions[action]; !registered {
		if options.Action != "" {
			return errors.New("whichkey: requested configuration launch action is not registered on the captured server; register it with menu.register_launch or omit --action for the basic menu")
		}
		// Native interaction has no idle deadline. Preserve only the optional
		// selector until an executable retained selection needs an app.
		terminalEndpoint, endpointErr := rex.ResolveEndpoint(ctx, nil, selection)
		if endpointErr != nil || terminalEndpoint != server {
			return errors.New("whichkey: native popup requires the originating execution-host terminal endpoint; run on that host without a conflicting server override")
		}
		return launchNativePopup(parent, origin, selection, requested)
	} else if kind != "lua" {
		return errors.New("whichkey: launch name is not a server configuration action; choose a registered menu.register_launch action")
	}
	if actions["whichkey_runtime"] != "lua" {
		return errRuntimeCompatibility
	}
	consumer, err := queryRuntimeConsumer(ctx, selection, origin, action)
	if err != nil {
		return err
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
	return launchPublicRuntime(parent, origin, placement, consumer)
}

func publicLaunchAction(options LaunchOptions) (string, error) {
	action := options.Action
	if action == "" {
		action = "whichkey"
	}
	if strings.Trim(action, "abcdefghijklmnopqrstuvwxyz0123456789_") != "" || len(action) > 256 {
		return "", errors.New("whichkey: --action must name a registered configuration action (lowercase letters, digits and underscores only)")
	}
	if options.Server != nil {
		if _, err := rex.CanonicalEndpoint(*options.Server); err != nil {
			return "", err
		}
	}
	return action, nil
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
