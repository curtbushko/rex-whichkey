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
	// Selection reuses the bootstrap's already-frozen CLI without discovery.
	Selection *CLISelection
	// Server is a concrete endpoint. Empty uses REX_SERVER. A server remains
	// required here: endpoint discovery is a separate, explicit bootstrap step.
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
	selection  *CLISelection
	server     string
	session    string
	timeout    time.Duration
}

// New is the context-free convenience entry for NewContext. Construction does
// not contact, autostart or mutate a server.
func New(options Options) (*Client, error) {
	return NewContext(context.Background(), options)
}

// NewContext retains one PATH-selected CLI and a concrete canonical target.
// Use this entry when bootstrap cancellation is available.
func NewContext(ctx context.Context, options Options) (*Client, error) {
	if ctx == nil {
		return nil, errors.New("rex: client construction requires a context")
	}
	ctx, cancel := context.WithTimeout(ctx, discoveryTimeout)
	defer cancel()
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
	server, err := CanonicalEndpoint(options.Server)
	if err != nil {
		return nil, err
	}
	selection, err := ResolveCLI(ctx, CLIOptions{Selection: options.Selection})
	if err != nil {
		return nil, err
	}
	client := &Client{
		executable: selection.Path,
		selection:  selection,
		server:     server,
		session:    options.Session,
		timeout:    options.Timeout,
	}
	return client, nil
}

// Server returns this client's captured concrete endpoint.
func (c *Client) Server() string {
	if c == nil {
		return ""
	}
	return c.server
}

// Selection returns a detached path snapshot for invocation-local reuse.
// It carries no endpoint, resource, GUI or compatibility authority.
func (c *Client) Selection() *CLISelection {
	if c == nil || c.selection == nil {
		return nil
	}
	return c.selection.clone()
}

// Call invokes a documented API method exactly once. Arguments must marshal to
// one JSON object, or be nil to omit ARGS. Pass json.RawMessage for raw JSON.
// Result must be a non-nil pointer, or nil to discard the validated JSON result.
// Child output and user-defined decoding errors are never diagnostic text.
// A timeout or decode error does not mean the operation had no effect; callers
// must not blindly retry mutations.
func (c *Client) Call(ctx context.Context, method string, arguments, result any) error {
	if err := c.validateCall(ctx, method, result); err != nil {
		return err
	}

	argv := []string{
		flagServer, c.server,
		flagNoAutostart,
		flagTimeout, c.timeout.String(),
		"api", "call",
	}
	if c.session != "" {
		argv = append(argv, "--session", c.session)
	}
	// End flag parsing before positional values, including user-supplied ones.
	argv = append(argv, "--", method)
	if arguments != nil {
		encoded, err := encodeArguments(arguments)
		if err != nil {
			return err
		}
		argv = append(argv, encoded)
	}

	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	output, err := run(ctx, c.executable, argv, c.selection)
	if err != nil {
		return errors.Join(errors.New("rex: API call failed or uncertain; not retried"), err)
	}
	if result == nil {
		var discarded json.RawMessage
		result = &discarded
	}
	if err := json.Unmarshal(output, result); err != nil {
		return errors.New("rex: invalid API result (operation may have succeeded); not retried")
	}
	return nil
}

func (c *Client) validateCall(ctx context.Context, method string, result any) error {
	if c == nil || c.executable == "" || c.server == "" || c.timeout <= 0 {
		return errors.New("rex: client must be constructed with New")
	}
	if ctx == nil {
		return errors.New("rex: call requires a context")
	}
	if err := ctx.Err(); err != nil {
		return errors.Join(errors.New("rex: API call canceled; not retried"), err)
	}
	if strings.TrimSpace(method) == "" || strings.ContainsRune(method, '\x00') {
		return errors.New("rex: API method must not be empty or contain NUL")
	}
	if result != nil {
		value := reflect.ValueOf(result)
		if value.Kind() != reflect.Pointer || value.IsNil() {
			return fmt.Errorf("rex: invalid result destination: %w", &json.InvalidUnmarshalError{Type: value.Type()})
		}
	}

	return nil
}

func encodeArguments(arguments any) (string, error) {
	data, err := json.Marshal(arguments)
	if err != nil {
		return "", errors.New("rex: API arguments cannot be encoded as a JSON object")
	}
	var object map[string]json.RawMessage
	if err := json.Unmarshal(data, &object); err != nil {
		return "", errors.New("rex: API arguments must encode one JSON object")
	}
	if object == nil {
		return "", errors.New("rex: API arguments must be a JSON object, not null")
	}
	return string(data), nil
}
