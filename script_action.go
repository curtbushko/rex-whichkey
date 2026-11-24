package whichkey

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const menuScriptRun = "script.run"

type scriptAction struct {
	path string
	args []string
}
type preparedScript struct {
	path, cwd string
	args      []string
}

// The acknowledgement has a small diagnostic budget. Keep path-free repair
// guidance separate from the complete, inspectable error returned to callers.
type scriptStartupProblem struct {
	problem         error
	cause, recovery string
}

func (e *scriptStartupProblem) Error() string { return e.problem.Error() }
func (e *scriptStartupProblem) Unwrap() error { return e.problem }

func scriptProblem(problem error, cause, recovery string) error {
	return &scriptStartupProblem{problem: problem, cause: cause, recovery: recovery}
}

func scriptOSCause(problem error) string {
	var pathError *os.PathError
	if errors.As(problem, &pathError) {
		return pathError.Err.Error()
	}
	return problem.Error()
}

func decodeScriptAction(data json.RawMessage) (scriptAction, error) {
	if err := checkMenuJSON(data); err != nil {
		return scriptAction{}, err
	}
	fields, err := menuObject(data, "path", "args")
	if err != nil {
		return scriptAction{}, err
	}
	path, err := completionString(fields["path"])
	if err != nil || path == "" || strings.ContainsRune(path, 0) {
		return scriptAction{}, errors.New("script.run requires a nonempty literal path without NUL")
	}
	action := scriptAction{path: path}
	if raw, present := fields["args"]; present {
		if len(raw) == 0 || raw[0] != '[' {
			return scriptAction{}, errors.New("script.run args must be an array of strings")
		}
		if err := json.Unmarshal(raw, &action.args); err != nil {
			return scriptAction{}, errors.New("script.run args must be an array of strings")
		}
		for _, arg := range action.args {
			if strings.ContainsRune(arg, 0) {
				return scriptAction{}, errors.New("script.run arguments must not contain NUL")
			}
		}
	}
	return action, nil
}

// Only call after fresh native host/account proof. Paths are literal and never
// searched on PATH; symlinks may resolve to ordinary executable files. The OS
// still owns the final chdir/exec check: this is not a same-account sandbox.
func prepareScript(action scriptAction, cwd string) (preparedScript, error) {
	if !filepath.IsAbs(cwd) || strings.ContainsRune(cwd, 0) {
		return preparedScript{}, scriptProblem(errors.New("script not started: original foreground cwd is unavailable or not absolute"), "Original cwd unavailable or not absolute", "Check the origin pane's foreground cwd.")
	}
	directory, err := os.Stat(cwd)
	if err != nil {
		return preparedScript{}, scriptProblem(fmt.Errorf("script not started: original cwd %q: %w", cwd, err), "Original cwd: "+scriptOSCause(err), "Check the origin pane's foreground cwd.")
	}
	if !directory.IsDir() {
		return preparedScript{}, scriptProblem(fmt.Errorf("script not started: original cwd %q is not a directory", cwd), "Original cwd is not a directory", "Check the origin pane's foreground cwd.")
	}
	path := action.path
	if !filepath.IsAbs(path) {
		// Join/Clean would change traversal through symlinks and could erase
		// missing intermediate directories. Leave resolution to the filesystem.
		if !strings.HasSuffix(cwd, string(filepath.Separator)) {
			path = string(filepath.Separator) + path
		}
		path = cwd + path
	}
	info, err := os.Stat(path)
	if err != nil {
		return preparedScript{}, scriptProblem(fmt.Errorf("script not started: target %q: %w", path, err), "Target: "+scriptOSCause(err), "Check the literal path in the origin cwd.")
	}
	if !info.Mode().IsRegular() {
		return preparedScript{}, scriptProblem(fmt.Errorf("script not started: target %q must be a regular executable script", path), "Target must be a regular executable script", "Choose a regular file, not a directory or device.")
	}
	if info.Mode().Perm()&0111 == 0 {
		return preparedScript{}, scriptProblem(fmt.Errorf("script not started: target %q is not executable; set executable permission and a valid shebang", path), "Target is not executable", "Set executable permission and a valid shebang.")
	}
	return preparedScript{path: path, cwd: cwd, args: append([]string(nil), action.args...)}, nil
}
