// Package rex calls the Rex Server API through the rex executable.
package rex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

const defaultTimeout = 10 * time.Second

// Options selects the CLI and its target in the host running this Go process.
type Options struct {
	// Executable is a program name or path on the execution host, not the
	// remote server. Empty uses REX_BIN_DIR/rex, or rex from PATH.
	Executable string
	// Server is passed unchanged to --server. Empty uses REX_SERVER. A server
	// is required: losing originating context must not select a local server.
	Server string
	// Session scopes every call made by this client. Empty leaves calls
	// unscoped; api call does not default to REX_SESSION. A session_id in the
	// arguments can also select a session, as documented by Rex.
	Session string
	// Timeout bounds the entire CLI process as well as Rex requests. Zero
	// uses ten seconds; an earlier context deadline still takes precedence.
	Timeout time.Duration
}

// Client is a fixed Rex server/session target. It does not own Rex resources.
// Use an unscoped client for methods such as session.list and session.create.
type Client struct {
	executable string
	server     string
	session    string
	timeout    time.Duration
}

// New captures the target and resolves the executable in the execution host's
// environment. It does not connect to Rex or start a server.
func New(options Options) (*Client, error) {
	if options.Server == "" {
		options.Server = os.Getenv("REX_SERVER")
	}
	if strings.TrimSpace(options.Server) == "" {
		return nil, errors.New("rex: server is required (set Options.Server or REX_SERVER)")
	}
	if strings.ContainsRune(options.Server, '\x00') || strings.ContainsRune(options.Session, '\x00') {
		return nil, errors.New("rex: server and session must not contain NUL")
	}
	if options.Timeout < 0 {
		return nil, errors.New("rex: timeout must not be negative")
	}
	if options.Timeout == 0 {
		options.Timeout = defaultTimeout
	}
	executable, err := resolveExecutable(options.Executable)
	if err != nil {
		return nil, err
	}
	return &Client{
		executable: executable,
		server:     options.Server,
		session:    options.Session,
		timeout:    options.Timeout,
	}, nil
}

// Call invokes a documented API method exactly once. Arguments must marshal to
// one JSON object, or be nil to omit ARGS. Pass json.RawMessage for raw JSON.
// Result must be a non-nil pointer, or nil to discard the validated JSON result.
// JSON decoding errors retain their encoding/json types through error wrapping.
// A timeout or decode error does not mean the operation had no effect; callers
// must not blindly retry mutations.
func (c *Client) Call(ctx context.Context, method string, arguments, result any) error {
	if c == nil || c.executable == "" || c.server == "" || c.timeout <= 0 {
		return errors.New("rex: client must be constructed with New")
	}
	if ctx == nil {
		return errors.New("rex: call requires a context")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("rex call %q: %w", method, err)
	}
	if strings.TrimSpace(method) == "" || strings.ContainsRune(method, '\x00') {
		return errors.New("rex: API method must not be empty or contain NUL")
	}
	if result != nil {
		value := reflect.ValueOf(result)
		if value.Kind() != reflect.Pointer || value.IsNil() {
			return fmt.Errorf("rex call %q result: %w", method, &json.InvalidUnmarshalError{Type: value.Type()})
		}
	}

	argv := []string{
		"--server", c.server,
		"--autostart=false",
		"--timeout", c.timeout.String(),
		"api", "call",
	}
	if c.session != "" {
		argv = append(argv, "--session", c.session)
	}
	// End flag parsing before positional values, including user-supplied ones.
	argv = append(argv, "--", method)
	if arguments != nil {
		data, err := json.Marshal(arguments)
		if err != nil {
			return fmt.Errorf("rex call %q arguments: %w", method, err)
		}
		var object map[string]json.RawMessage
		if err := json.Unmarshal(data, &object); err != nil {
			return fmt.Errorf("rex call %q arguments must be a JSON object: %w", method, err)
		}
		if object == nil {
			return fmt.Errorf("rex call %q arguments must be a JSON object, not null", method)
		}
		argv = append(argv, string(data))
	}

	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	output, err := run(ctx, c.executable, argv)
	if err != nil {
		return fmt.Errorf("rex call %q: %w", method, err)
	}
	if result == nil {
		var discarded json.RawMessage
		result = &discarded
	}
	if err := json.Unmarshal(output, result); err != nil {
		return fmt.Errorf("rex call %q decode result (operation may have succeeded): %w", method, err)
	}
	return nil
}
