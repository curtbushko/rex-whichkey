package rex

import (
	"context"
	"errors"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

// CLIOptions reuses an invocation's execution-host PATH selection when present.
// There is no executable override or alternate discovery authority.
type CLIOptions struct {
	Selection *CLISelection
}

// CLISelection retains only the selected executable path. Executable trust is
// ordinary PATH/OS policy, not a content, installation or API receipt. The
// canonical endpoint is separate invocation context owned by each Client.
type CLISelection struct {
	Path string `json:"path"`
}

var (
	errCLIInput  = errors.New("rex: invalid retained execution-host CLI path; not retried")
	errCLIAbsent = errors.New("rex: rex executable not found on execution-host PATH; install the licensed Rex CLI and ensure the launching process inherits its PATH; not retried")
	// Also used by strict selected-action metadata decoding, not startup admission.
	errCLICompatibility = errors.New("rex: invalid CLI result metadata; operation may have succeeded; not retried")
)

// ResolveCLI uses normal executable lookup once. Consumers and the private PTY
// child reuse that path without rediscovery, hashes or metadata negotiation.
func ResolveCLI(ctx context.Context, options CLIOptions) (*CLISelection, error) {
	if ctx == nil {
		return nil, errors.New("rex: CLI discovery requires a context")
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if options.Selection != nil {
		if !selectedPath(options.Selection.Path) {
			return nil, errCLIInput
		}
		return options.Selection.clone(), nil
	}
	path, err := exec.LookPath("rex")
	if err != nil {
		return nil, errCLIAbsent
	}
	// Freeze a relative PATH-directory result against later working-dir changes.
	path, err = filepath.Abs(path)
	if err != nil || !selectedPath(path) {
		return nil, errCLIInput
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return &CLISelection{Path: path}, nil
}

func selectedPath(path string) bool {
	return filepath.IsAbs(path) && len(path) <= 64<<10 && !strings.ContainsRune(path, '\x00')
}

func (selection *CLISelection) clone() *CLISelection {
	snapshot := *selection
	return &snapshot
}

func metadataText(value string) bool {
	return value != "" && len(value) <= 256 && utf8.ValidString(value) && strings.IndexFunc(value, unicode.IsControl) < 0
}
