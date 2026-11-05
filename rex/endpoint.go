package rex

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net"
	"net/url"
	"os"
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
// Rex host label. Unix socket paths and URIs normalize to escaped unix URIs;
// host:port and bracketed IPv6 forms retain their spelling. This does
// not connect, inspect files, resolve DNS or interpret opaque Lua ctx.server.
func CanonicalEndpoint(server string) (string, error) {
	invalid := errors.New("rex: endpoint must be an absolute Unix socket path, unix URI or host:port with a nonzero numeric port; labels and controls are not allowed")
	if !utf8.ValidString(server) || strings.IndexFunc(server, unicode.IsControl) >= 0 {
		return "", invalid
	}
	if filepath.IsAbs(server) {
		server = (&url.URL{Scheme: "unix", Path: server}).String()
	}
	if strings.HasPrefix(server, "unix://") {
		return canonicalUnixEndpoint(server, invalid)
	}
	host, port, err := net.SplitHostPort(server)
	if err != nil || host == "" || port == "" || strings.Trim(port, "0123456789") != "" ||
		strings.IndexFunc(host, unicode.IsSpace) >= 0 || strings.ContainsAny(host, "/\\?#@%") {
		return "", invalid
	}
	value, err := strconv.ParseUint(port, 10, 16)
	if err != nil || value == 0 {
		return "", invalid
	}
	return server, nil
}

func canonicalUnixEndpoint(server string, invalid error) (string, error) {
	address, err := url.Parse(server)
	if err != nil || address.Host != "" || !filepath.IsAbs(address.Path) || address.Path == "/" ||
		!utf8.ValidString(address.Path) || strings.IndexFunc(address.Path, unicode.IsControl) >= 0 ||
		address.RawQuery != "" || address.ForceQuery || strings.Contains(server, "#") || address.User != nil {
		return "", invalid
	}
	// whoami uses escaped URI paths, including Application%20Support on
	// macOS. Compare normalized addresses, not escaped versus decoded text.
	return (&url.URL{Scheme: "unix", Path: address.Path}).String(), nil
}

// ResolveEndpoint captures one concrete endpoint on the execution host. A nil
// override means absent, not an explicitly empty value. Only native discovery
// invokes Rex: platform runtime/home path algorithms remain Rex's responsibility.
// Callers must pass the returned address explicitly on all subsequent calls.
func ResolveEndpoint(ctx context.Context, explicit *string, frozen ...*CLISelection) (string, error) {
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
	if len(frozen) > 1 {
		return "", errCLIInput
	}
	var options CLIOptions
	if len(frozen) == 1 {
		if frozen[0] == nil {
			return "", errCLIInput
		}
		options.Selection = frozen[0]
	}
	selection, err := ResolveCLI(ctx, options)
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, endpointTimeout)
	defer cancel()
	// Do not let the originating terminal's target selectors redirect a query
	// intended for Rex's native platform default. Preserve HOME, XDG_*, Rex's
	// runtime/home settings and authentication environment without guessing paths.
	environment := make([]string, 0, len(os.Environ()))
	for _, entry := range os.Environ() {
		key, _, _ := strings.Cut(entry, "=")
		switch key {
		case "REX_SERVER", "REX_CLIENT", "REX_SESSION", "REX_WINDOW", "REX_BLOCK":
			continue
		}
		environment = append(environment, entry)
	}
	output, err := runSelected(ctx, selection, []string{
		flagNoAutostart, flagTimeout, endpointTimeout.String(), "whoami", flagJSON,
	}, environment, 64<<10)
	if err != nil {
		// Never return process diagnostics/argv or raw identity data: preparation
		// is a private pipe boundary, not a place to display external output.
		return "", errors.Join(errors.New("rex: native endpoint discovery failed or timed out; check the running execution-host server and Rex runtime/home settings, or supply an explicit endpoint; not retried"), ctx.Err())
	}
	server, err := identityEndpoint(output)
	if err != nil {
		return "", errors.New("rex: native identity JSON has no valid concrete endpoint; use a matching Rex CLI or supply an explicit endpoint")
	}
	return server, nil
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
