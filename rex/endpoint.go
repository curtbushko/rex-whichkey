package rex

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const endpointTimeout = 2 * time.Second

// CanonicalEndpoint validates a concrete popup transport address, never a mutable
// Rex host label. Absolute Unix socket paths become unix URIs; existing strict
// unix URI, host:port and bracketed IPv6 forms retain their spelling. This does
// not connect, inspect files, resolve DNS or interpret opaque Lua ctx.server.
func CanonicalEndpoint(server string) (string, error) {
	invalid := errors.New("rex: endpoint must be an absolute Unix socket path, unix URI or host:port with a nonzero numeric port; labels, whitespace and controls are not allowed")
	if !utf8.ValidString(server) || strings.IndexFunc(server, func(char rune) bool {
		return unicode.IsSpace(char) || unicode.IsControl(char)
	}) >= 0 {
		return "", invalid
	}
	if filepath.IsAbs(server) {
		server = "unix://" + server
	}
	if strings.HasPrefix(server, "unix://") {
		address, err := url.Parse(server)
		if err == nil && address.Host == "" && filepath.IsAbs(address.Path) &&
			address.Path != "/" && address.Path == strings.TrimPrefix(server, "unix://") &&
			address.RawQuery == "" && !address.ForceQuery && address.Fragment == "" && address.User == nil {
			return server, nil
		}
		return "", invalid
	}
	host, port, err := net.SplitHostPort(server)
	if err != nil || host == "" || port == "" || strings.Trim(port, "0123456789") != "" ||
		strings.ContainsAny(host, "/\\?#@%") {
		return "", invalid
	}
	value, err := strconv.ParseUint(port, 10, 16)
	if err != nil || value == 0 {
		return "", invalid
	}
	return server, nil
}

// ResolveEndpoint captures one concrete endpoint on the execution host. A nil
// override means absent, not an explicitly empty value. Only native discovery
// invokes Rex: platform runtime/home path algorithms remain Rex's responsibility.
// Callers must pass the returned address explicitly on all subsequent calls.
func ResolveEndpoint(ctx context.Context, explicit *string) (string, error) {
	if explicit != nil {
		return CanonicalEndpoint(*explicit)
	}
	if server := os.Getenv("REX_SERVER"); server != "" {
		return CanonicalEndpoint(server)
	}
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		return "", errors.New("rex: automatic endpoint discovery requires Linux or macOS; supply an explicit endpoint")
	}
	if ctx == nil {
		return "", errors.New("rex: endpoint discovery requires a context")
	}
	executable, err := resolveExecutable("")
	if err != nil {
		return "", errors.New("rex: endpoint discovery requires the matching Rex CLI on the execution host (REX_BIN_DIR or PATH)")
	}
	ctx, cancel := context.WithTimeout(ctx, endpointTimeout)
	defer cancel()
	command := exec.CommandContext(ctx, executable,
		"--autostart=false", "--timeout", endpointTimeout.String(), "whoami", "--json")
	// Do not let the originating terminal's target selectors redirect a query
	// intended for Rex's native platform default. Preserve HOME, XDG_*, Rex's
	// runtime/home settings and authentication environment without guessing paths.
	command.Env = make([]string, 0, len(os.Environ()))
	for _, entry := range os.Environ() {
		key, _, _ := strings.Cut(entry, "=")
		switch key {
		case "REX_SERVER", "REX_CLIENT", "REX_SESSION", "REX_WINDOW", "REX_BLOCK":
			continue
		}
		command.Env = append(command.Env, entry)
	}
	var output endpointOutput
	command.Stdout, command.Stderr, command.Stdin = &output, io.Discard, nil
	command.WaitDelay = 250 * time.Millisecond
	if err := command.Run(); err != nil {
		// Never return process diagnostics/argv or raw identity data: preparation
		// is a private pipe boundary, not a place to display external output.
		return "", errors.New("rex: native endpoint discovery failed or timed out; check the running execution-host server and Rex runtime/home settings, or supply an explicit endpoint; not retried")
	}
	server, err := identityEndpoint(output.buffer.Bytes())
	if err != nil {
		return "", errors.New("rex: native identity JSON has no valid concrete endpoint; use a matching Rex CLI or supply an explicit endpoint")
	}
	return server, nil
}

// Bound captured identity output without leaving a pipe reader or child behind.
// A writer error makes os/exec fail; no truncated response can become a target.
type endpointOutput struct{ buffer bytes.Buffer }

func (output *endpointOutput) Write(data []byte) (int, error) {
	if len(data) > (64<<10)-output.buffer.Len() {
		return 0, errors.New("rex: native identity output exceeds limit")
	}
	return output.buffer.Write(data)
}

func identityEndpoint(data []byte) (string, error) {
	invalid := errors.New("rex: invalid native identity JSON")
	if !utf8.Valid(data) || !json.Valid(data) {
		return "", invalid
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	token, err := decoder.Token()
	if err != nil || token != json.Delim('{') {
		return "", invalid
	}
	var endpoint string
	seen := make(map[string]bool)
	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			return "", invalid
		}
		name, ok := token.(string)
		if !ok || seen[name] {
			return "", invalid
		}
		seen[name] = true
		var value json.RawMessage
		if err := decoder.Decode(&value); err != nil {
			return "", invalid
		}
		if name == "endpoint" {
			if err := json.Unmarshal(value, &endpoint); err != nil {
				return "", invalid
			}
		}
	}
	// json.Valid already requires the closing object and a single JSON value.
	return CanonicalEndpoint(endpoint)
}
