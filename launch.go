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

// LaunchOptions is public data/endpoint/app intent. Nil selectors are absent;
// explicitly empty selectors fail rather than acquiring ambient authority.
type LaunchOptions struct {
	Server  *string
	Client  *string
	Config  *string
	ASCII   bool
	NoColor bool
}
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

// LaunchPopup validates data before any process or resource startup, captures
// the public image/CLI/endpoint once, and enters the real Go popup directly.
// Legacy registered actions cannot select, gate or redirect this path.
func LaunchPopup(parent context.Context, options LaunchOptions) (outcome error) {
	if parent == nil {
		return errors.New("whichkey: launch requires a context")
	}
	config, err := LoadMenuConfig(options.Config)
	if err != nil {
		return err
	}
	config, err = canonicalMenuSnapshot(config)
	if err != nil {
		return err
	}
	if options.Server != nil {
		if _, endpointErr := rex.CanonicalEndpoint(*options.Server); endpointErr != nil {
			return endpointErr
		}
	}
	origin := launchOrigin{SessionID: os.Getenv("REX_SESSION"), BlockID: os.Getenv("REX_BLOCK")}
	if !popupResourceID(origin.SessionID, "session") || !popupResourceID(origin.BlockID, "block") {
		return errors.New("whichkey: run inside the originating Rex terminal with full REX_SESSION and REX_BLOCK IDs; no focused-block fallback")
	}
	selector := options.Client
	if selector == nil {
		if value, present := os.LookupEnv("REX_CLIENT"); present {
			selector = &value
		}
	}
	requested := ""
	if selector != nil {
		if !popupResourceID(*selector, "client") {
			return errors.New("whichkey: --client or REX_CLIENT must be an exact full client:UUID")
		}
		requested = *selector
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
	origin.Server, err = rex.ResolveEndpoint(ctx, options.Server, selection)
	if err != nil {
		return err
	}
	client, err := rex.NewContext(ctx, rex.Options{Selection: selection, Server: origin.Server, Session: origin.SessionID, Timeout: 3 * time.Second})
	if err != nil {
		return errors.Join(errRuntimeCompatibility, err)
	}
	origin.selection = client.Selection()
	return launchNativePopup(parent, origin, origin.selection, requested, config, options.NoColor, options.ASCII)
}

// launchOutput is isolated historical diagnostic compatibility for unchanged
// tests. No supported public launch or Rex execution path constructs it.
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
