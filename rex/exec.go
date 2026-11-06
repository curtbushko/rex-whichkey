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
	layoutProvenance = "layout-dependency"
	schemaNull       = "null"
	schemaObject     = "object"
	schemaArray      = "array"
	schemaString     = "string"
	schemaBoolean    = "boolean"
	schemaNumber     = "number"
	fieldFocus       = "focus"
	fieldSession     = "session_id"
	fieldBlock       = "block_id"
	fieldBounds      = "bounds"
	fieldLabel       = "label"
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

// run is the shared resource boundary. Inventory has no endpoint; all other
// commands require the client's retained endpoint/schema receipt. Identity loss
// after an effect is uncertain, never permission to repeat it.
func run(ctx context.Context, executable string, argv []string, selection *CLISelection) ([]byte, error) {
	if selection == nil || selection.Executable.Path != executable {
		return nil, errCLIChanged
	}
	inventory := inventoryArgs(argv)
	if !inventory {
		if len(argv) < 2 || argv[0] != flagServer || selection.Capabilities == nil || argv[1] != selection.Capabilities.Endpoint {
			return nil, errCLIChanged
		}
		if err := selection.checkServer(ctx); err != nil {
			return nil, err
		}
	}
	output, err := runSelected(ctx, selection, argv, nil, 8<<20)
	if err != nil {
		return nil, err
	}
	if !inventory {
		if err := selection.checkServer(ctx); err != nil {
			return nil, err
		}
	}
	return output, nil
}

func inventoryArgs(argv []string) bool {
	return len(argv) == 5 && argv[0] == flagNoAutostart && argv[1] == flagTimeout && argv[3] == "hosts" && argv[4] == flagJSON
}

func runSelected(ctx context.Context, selection *CLISelection, argv, environment []string, limit int) ([]byte, error) {
	if ctx == nil {
		return nil, errors.New("rex: CLI invocation requires a context")
	}
	// Callers own the invocation deadline (including long terminal waits).
	// Metadata callers supply their own short discovery bound.
	if err := selection.checkFiles(ctx); err != nil {
		return nil, err
	}
	// No shell: spaces, metacharacters and JSON remain single argv values.
	command := exec.CommandContext(ctx, selection.Executable.Path, argv...)
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
	// A changed dependency/link invalidates even a read-only reply. A failed
	// post-effect check is uncertain, never permission to repeat the command.
	if err := selection.checkFiles(ctx); err != nil {
		return nil, err
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

func decodeVersion(data []byte) (CLIVersion, error) {
	object, err := metadataObject(data)
	if err != nil {
		return CLIVersion{}, compatibilityFailure("server.version", "object")
	}
	var version CLIVersion
	if json.Unmarshal(data, &version) != nil {
		return CLIVersion{}, compatibilityFailure("server.version", "known field types")
	}
	if !metadataText(version.Version) {
		return CLIVersion{}, compatibilityFailure("server.version", "rex_version text")
	}
	if version.APIMinRevision > version.APIRevision {
		return CLIVersion{}, compatibilityFailure("server.version", "api_min_revision exceeds api_revision")
	}
	for _, field := range []string{"rex_version", "api_revision", "api_min_revision"} {
		if value, present := object[field]; present && bytes.Equal(bytes.TrimSpace(value), []byte(schemaNull)) {
			return CLIVersion{}, compatibilityFailure("server.version", field+" null")
		}
	}
	return version, nil
}

// Reject duplicate keys and invalid UTF-8 at metadata boundaries. Unknown
// optional fields remain opaque; no unknown field is taken as capability proof.
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
