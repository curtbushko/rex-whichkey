package rex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// ExecError describes a failed CLI invocation. Stderr preserves Rex's diagnostic
// message, including any server status/code. ExitCode is -1 if no exit code is
// available. Unwrap preserves process errors and context cancellation/deadlines.
type ExecError struct {
	Executable string
	Args       []string
	ExitCode   int
	Stderr     string
	Err        error
}

func (e *ExecError) Error() string {
	message := fmt.Sprintf("rex executable %q (exit %d): %v", e.Executable, e.ExitCode, e.Err)
	if diagnostic := strings.TrimSpace(e.Stderr); diagnostic != "" {
		message += ": " + diagnostic
	}
	return message
}

// Unwrap returns the underlying execution failure.
func (e *ExecError) Unwrap() error { return e.Err }

func resolveExecutable(name string) (string, error) {
	if name == "" {
		name = "rex"
		if directory := os.Getenv("REX_BIN_DIR"); directory != "" {
			// This directory belongs to the execution host. Do not replace it
			// with a path discovered on a client or silently try another rex.
			name = filepath.Join(directory, name)
			if !filepath.IsAbs(name) {
				// Keep even REX_BIN_DIR=. an explicit path, not a PATH search.
				name = "." + string(os.PathSeparator) + name
			}
		}
	}
	path, err := exec.LookPath(name)
	if err != nil {
		return "", fmt.Errorf("rex: resolve executable %q on execution host: %w", name, err)
	}
	return path, nil
}

func run(ctx context.Context, executable string, argv []string) ([]byte, error) {
	// No shell: spaces, metacharacters and JSON remain single argv values.
	command := exec.CommandContext(ctx, executable, argv...)
	var stdout, stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr
	// A child inheriting the pipes must not leave cancellation waiting forever.
	command.WaitDelay = time.Second
	if err := command.Run(); err != nil {
		if contextErr := ctx.Err(); contextErr != nil {
			err = errors.Join(err, contextErr)
		}
		code := -1
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			code = exitError.ExitCode()
		}
		return nil, &ExecError{
			Executable: executable,
			Args:       append([]string(nil), argv...),
			ExitCode:   code,
			Stderr:     stderr.String(),
			Err:        err,
		}
	}
	return stdout.Bytes(), nil
}
