package rex

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"time"
	"unicode/utf8"
)

const (
	discoveryTimeout = 3 * time.Second
	flagServer       = "--server"
	flagNoAutostart  = "--autostart=false"
	flagTimeout      = "--timeout"
	flagJSON         = "--json"
	fieldFocus       = "focus"
	fieldSession     = "session_id"
	fieldBlock       = "block_id"
	fieldBounds      = "bounds"
	fieldLabel       = "label"
	fieldArgs        = "args"
)

// ExecError deliberately excludes child diagnostics and arguments from Error.
// Stderr is kept empty; public fields remain for source compatibility only.
// Unwrap preserves cancellation without revealing a child-controlled message.
type ExecError struct {
	Executable string
	Args       []string
	ExitCode   int
	Stderr     string
	Err        error
}

func (e *ExecError) Error() string {
	return fmt.Sprintf("rex: CLI invocation failed or uncertain (exit %d); check the execution-host installation and server; not retried", e.ExitCode)
}

// Unwrap returns a fixed failure or context cancellation, never child output.
func (e *ExecError) Unwrap() error { return e.Err }

// run is the shared resource boundary. Inventory has no endpoint; resource
// commands carry their Client's explicit canonical endpoint and no-autostart.
// Execute only the requested command, with no status sandwiches or retries.
func run(ctx context.Context, executable string, argv []string, selection *CLISelection) ([]byte, error) {
	if selection == nil || selection.Path != executable {
		return nil, errCLIInput
	}
	if !inventoryArgs(argv) {
		if len(argv) < 3 || argv[0] != flagServer || argv[2] != flagNoAutostart {
			return nil, errCLIInput
		}
		endpoint, err := CanonicalEndpoint(argv[1])
		if err != nil || endpoint != argv[1] {
			return nil, errCLIInput
		}
	}
	return runSelected(ctx, selection, argv, nil, 8<<20)
}

func inventoryArgs(argv []string) bool {
	return len(argv) == 5 && argv[0] == flagNoAutostart && argv[1] == flagTimeout && argv[3] == "hosts" && argv[4] == flagJSON
}

func runSelected(ctx context.Context, selection *CLISelection, argv, environment []string, limit int) ([]byte, error) {
	if ctx == nil {
		return nil, errors.New("rex: CLI invocation requires a context")
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if selection == nil || !selectedPath(selection.Path) {
		return nil, errCLIInput
	}
	// Callers own the deadline, including long terminal waits. No shell:
	// spaces, metacharacters and JSON remain single argv values.
	command := exec.CommandContext(ctx, selection.Path, argv...)
	command.Env = environment
	var output boundedOutput
	output.limit = limit
	command.Stdout, command.Stderr = &output, io.Discard
	command.WaitDelay = 250 * time.Millisecond
	if err := command.Run(); err != nil {
		code := -1
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			code = exitError.ExitCode()
		}
		return nil, &ExecError{ExitCode: code,
			Err: errors.Join(errors.New("rex: CLI operation failed or uncertain; not retried"), ctx.Err())}
	}
	return output.buffer.Bytes(), nil
}

type boundedOutput struct {
	buffer bytes.Buffer
	limit  int
}

func (output *boundedOutput) Write(data []byte) (int, error) {
	if len(data) > output.limit-output.buffer.Len() {
		return 0, errors.New("rex: CLI output exceeds limit")
	}
	return output.buffer.Write(data)
}

// Reject duplicate keys and invalid UTF-8 in actual action metadata results.
// Unknown optional fields remain opaque.
func metadataObject(data []byte) (map[string]json.RawMessage, error) {
	if !utf8.Valid(data) || !json.Valid(data) {
		return nil, errCLICompatibility
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	token, err := decoder.Token()
	if err != nil || token != json.Delim('{') {
		return nil, errCLICompatibility
	}
	result := make(map[string]json.RawMessage)
	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			return nil, errCLICompatibility
		}
		name, ok := token.(string)
		if !ok || result[name] != nil {
			return nil, errCLICompatibility
		}
		var value json.RawMessage
		if err := decoder.Decode(&value); err != nil {
			return nil, errCLICompatibility
		}
		result[name] = value
	}
	return result, nil
}
