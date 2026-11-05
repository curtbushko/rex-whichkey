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
	"unicode"
	"unicode/utf8"
)

const discoveryTimeout = 3 * time.Second

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

// run's optional frozen input is transitional for the existing concrete resource
// wrapper. An unbound invocation is rejected, never silently reselected. All
// callers must pass the selection captured by their Client.
func run(ctx context.Context, executable string, argv []string, frozen ...*CLISelection) ([]byte, error) {
	if len(frozen) != 1 || frozen[0] == nil || frozen[0].Executable.Path != executable {
		return nil, errors.New("rex: resource command requires its client's frozen CLI selection; update the integration; not retried")
	}
	return runSelected(ctx, frozen[0], argv, nil, 8<<20)
}

func runSelected(ctx context.Context, selection *CLISelection, argv, environment []string, limit int) ([]byte, error) {
	if ctx == nil {
		return nil, errors.New("rex: CLI invocation requires a context")
	}
	ctx, cancel := context.WithTimeout(ctx, discoveryTimeout)
	defer cancel()
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
		return CLIVersion{}, errCLICompatibility
	}
	var version CLIVersion
	if err := json.Unmarshal(data, &version); err != nil || version.Version == "" || len(version.Version) > 256 ||
		!utf8.ValidString(version.Version) || version.APIMinRevision > version.APIRevision {
		return CLIVersion{}, errCLICompatibility
	}
	for _, r := range version.Version {
		if unicode.IsControl(r) {
			return CLIVersion{}, errCLICompatibility
		}
	}
	for _, field := range []string{"version", "api_revision", "api_min_revision"} {
		if value, present := object[field]; present && bytes.Equal(bytes.TrimSpace(value), []byte("null")) {
			return CLIVersion{}, errCLICompatibility
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
