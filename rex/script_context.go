package rex

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// ScriptHost is native discovery evidence, not yet execution authority. The
// caller must independently verify its local Unix socket peer PID and UID before
// using host-local files. A remote selector or renderer location proves nothing.
type ScriptHost struct {
	Endpoint  string `json:"endpoint"`
	HostID    string `json:"host_id"`
	PID       int    `json:"pid"`
	Principal struct {
		Kind      string `json:"kind"`
		Transport string `json:"transport"`
		UID       *int   `json:"uid"`
	} `json:"principal"`
}

// NativeScriptHost deliberately bypasses all inherited target selectors. The
// frozen CLI still owns platform-default socket discovery; unknown/non-Unix
// results fail closed rather than falling back to the addressed remote server.
func (c *Client) NativeScriptHost(ctx context.Context) (ScriptHost, error) {
	environment := make([]string, 0, len(os.Environ()))
	for _, entry := range os.Environ() {
		key, _, _ := strings.Cut(entry, "=")
		switch key {
		case "REX_SERVER", "REX_CLIENT", "REX_SESSION", "REX_WINDOW", "REX_BLOCK":
			continue
		}
		environment = append(environment, entry)
	}
	ctx, cancel := context.WithTimeout(ctx, endpointTimeout)
	defer cancel()
	output, err := runSelected(ctx, c.Selection(), []string{flagNoAutostart, flagTimeout, endpointTimeout.String(), "whoami", flagJSON}, environment, 64<<10)
	if err != nil {
		return ScriptHost{}, errors.New("script not started: native local host/account discovery failed; no fallback")
	}
	return decodeNativeScriptHost(output)
}

func decodeNativeScriptHost(output []byte) (ScriptHost, error) {
	var host ScriptHost
	if json.Unmarshal(output, &host) != nil || host.HostID == "" || host.PID <= 0 || host.Principal.Kind != "peer" || host.Principal.Transport != scriptUnixTransport || host.Principal.UID == nil || *host.Principal.UID != os.Geteuid() {
		return ScriptHost{}, errors.New("script not started: native local host/account identity unavailable")
	}
	endpoint, err := identityEndpoint(output)
	if err != nil || !strings.HasPrefix(endpoint, "unix://") {
		return ScriptHost{}, errors.New("script not started: native host requires a local Unix socket; network execution is not inferred")
	}
	host.Endpoint = endpoint
	return host, nil
}

// ScriptCWD reads only the two retained terminal identities, then the original
// foreground process. The caller has already independently proved native host
// authority. Missing cwd never substitutes OSC pwd, child cwd or popup cwd.
func (c *Client) ScriptCWD(ctx context.Context, host ScriptHost, origin, own Block) (string, error) {
	for _, block := range []Block{origin, own} {
		if err := c.scriptTerminal(ctx, host, block); err != nil {
			return "", err
		}
	}
	return c.scriptForegroundCWD(ctx, origin)
}

const scriptUnixTransport = "unix"

func (c *Client) scriptTerminal(ctx context.Context, host ScriptHost, block Block) error {
	if block.client == nil || block.Server() != c.server || block.client.session != c.session {
		return errors.New("script not started: captured origin/host scope changed")
	}
	current, err := block.Inspect(ctx)
	if err != nil || current.WindowID != block.WindowID || current.LayerID != block.LayerID {
		return errors.New("script not started: captured terminal placement changed")
	}
	var result struct {
		Block struct {
			ID        string `json:"block_id"`
			HostID    string `json:"host_id"`
			WindowID  string `json:"window_id"`
			Placement string `json:"placement"`
			Creator   string `json:"creator_name"`
			Flavor    string `json:"flavor"`
		} `json:"block"`
	}
	if err := c.Call(ctx, "session.describe_block", map[string]string{fieldSession: c.session, fieldBlock: block.blockID}, &result); err != nil {
		return err
	}
	value := result.Block
	if host.HostID == "" || value.ID != block.blockID || value.HostID != host.HostID || value.WindowID != block.WindowID || value.Placement != "placed" || value.Creator != "com.superlogical.terminal" || value.Flavor != "com.superlogical.terminal.shell" {
		return errors.New("script not started: origin and menu must share the proven native machine/account; no local fallback")
	}
	return nil
}

func (c *Client) scriptForegroundCWD(ctx context.Context, origin Block) (string, error) {
	var result struct {
		Foreground *struct {
			CWD    string `json:"cwd"`
			UserID *int   `json:"user_id"`
		} `json:"foreground"`
	}
	if err := c.Call(ctx, "com.superlogical.terminal.process", map[string]any{fieldSession: c.session, fieldBlock: origin.blockID, "args": map[string]any{}}, &result); err != nil {
		return "", err
	}
	if result.Foreground == nil || !filepath.IsAbs(result.Foreground.CWD) || strings.ContainsRune(result.Foreground.CWD, 0) || result.Foreground.UserID == nil || *result.Foreground.UserID != os.Geteuid() {
		return "", errors.New("script not started: original foreground cwd/account unavailable; no cwd fallback")
	}
	return result.Foreground.CWD, nil
}
