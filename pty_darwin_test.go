//go:build darwin

package whichkey

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

const sandboxEnv = "REX_WHICHKEY_PTY_SANDBOX"

// The executable fixture is selected only by this test binary's basename and
// an isolated directory. Unknown commands fail closed; PATH contains no real rex.
func TestMain(m *testing.M) {
	if os.Getenv(sandboxEnv) != "" && filepath.Base(os.Args[0]) == "rex" {
		os.Exit(directCLIFixture())
	}
	os.Exit(m.Run())
}

type ptyReady struct {
	PID      int
	Name     string
	Original unix.Termios
}

type ptyResult struct {
	Keys      string
	Text      *string
	Confirmed bool
	Dismissed bool
	Problem   string
}

type ptySandbox struct {
	dir      string
	input    io.WriteCloser
	wait     <-chan error
	cancel   context.CancelFunc
	ready    ptyReady
	finished bool
}

func sandboxJSON(t *testing.T, path string, value any) {
	t.Helper()
	data, err := json.Marshal(value)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(path+".tmp", data, 0600))
	require.NoError(t, os.Rename(path+".tmp", path))
}

func sandboxRead(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return ""
	}
	require.NoError(t, err)
	return string(data)
}

func sandboxAwait(t *testing.T, condition func() bool, detail func() string) {
	t.Helper()
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()
	deadline := time.NewTimer(8 * time.Second)
	defer deadline.Stop()
	for !condition() {
		select {
		case <-ticker.C:
		case <-deadline.C:
			t.Fatalf("sandbox readiness timed out: %s", detail())
		}
	}
}

func startPTYSandbox(t *testing.T, test, scenario string) *ptySandbox {
	t.Helper()
	dir := t.TempDir()
	executable, err := os.Executable()
	require.NoError(t, err)
	require.NoError(t, os.Symlink(executable, filepath.Join(dir, "rex")))
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	t.Cleanup(cancel)
	command := exec.CommandContext(ctx, "/usr/bin/script", "-q", "/dev/null", executable, "-test.run", "^"+test+"$")
	command.Env = append(os.Environ(), sandboxEnv+"="+dir, "REX_WHICHKEY_PTY_CASE="+scenario,
		"PATH="+dir, "NO_COLOR=", "REX_SESSION="+sandboxSession, "REX_BLOCK="+sandboxOwn,
		// Re-executed race fixtures retain instrumentation, not the runtime's
		// default one-second exit grace per CLI call (production has none).
		"GORACE="+os.Getenv("GORACE")+" atexit_sleep_ms=0")
	output, err := os.Create(filepath.Join(dir, "output"))
	require.NoError(t, err)
	t.Cleanup(func() { assert.NoError(t, output.Close()) })
	command.Stdout, command.Stderr = output, output
	input, err := command.StdinPipe()
	require.NoError(t, err)
	sandbox := &ptySandbox{dir: dir, input: input, cancel: cancel}
	t.Cleanup(func() {
		sandbox.stop(t)
		cancel()
		closeErr := input.Close()
		assert.True(t, closeErr == nil || errors.Is(closeErr, os.ErrClosed), "%v", closeErr)
	})
	require.NoError(t, command.Start())
	wait := make(chan error, 1)
	sandbox.wait = wait
	go func() { wait <- command.Wait() }()
	sandboxAwait(t, func() bool { return sandboxRead(t, filepath.Join(dir, "ready")) != "" }, sandbox.output)
	require.NoError(t, json.Unmarshal([]byte(sandboxRead(t, filepath.Join(dir, "ready"))), &sandbox.ready))
	return sandbox
}

func (sandbox *ptySandbox) stop(t *testing.T) {
	t.Helper()
	if sandbox.finished {
		return
	}
	stopSandboxWorker(t, sandbox.dir)
	if sandbox.ready.PID != 0 {
		err := syscall.Kill(sandbox.ready.PID, syscall.SIGKILL)
		assert.True(t, err == nil || errors.Is(err, syscall.ESRCH), "%v", err)
	}
	sandbox.cancel()
	if sandbox.wait != nil {
		<-sandbox.wait
	}
}

// Failed assertions must not orphan a deliberately blocked inventory child,
// even when the menu itself has to be killed. Cleanup must continue on errors.
func stopSandboxWorker(t *testing.T, dir string) {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(dir, "worker"))
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	if !assert.NoError(t, err) {
		return
	}
	var pid int
	if !assert.NoError(t, json.Unmarshal(data, &pid)) || !assert.Positive(t, pid) {
		return
	}
	err = syscall.Kill(pid, syscall.SIGKILL)
	assert.True(t, err == nil || errors.Is(err, syscall.ESRCH), "%v", err)
}

func (sandbox *ptySandbox) output() string {
	data, err := os.ReadFile(filepath.Join(sandbox.dir, "output"))
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (sandbox *ptySandbox) expect(t *testing.T, after int, text string) {
	t.Helper()
	sandboxAwait(t, func() bool {
		output := sandbox.output()
		return strings.Contains(output[min(after, len(output)):], text)
	}, sandbox.output)
}

func (sandbox *ptySandbox) send(t *testing.T, text string) int {
	t.Helper()
	mark := len(sandbox.output())
	_, err := io.WriteString(sandbox.input, text)
	require.NoError(t, err)
	return mark
}

func (sandbox *ptySandbox) resize(t *testing.T, columns, rows uint16) int {
	t.Helper()
	fd, err := unix.Open(sandbox.ready.Name, unix.O_RDWR|unix.O_NOCTTY, 0)
	require.NoError(t, err)
	defer func() { require.NoError(t, unix.Close(fd)) }()
	mark := len(sandbox.output())
	require.NoError(t, unix.IoctlSetWinsize(fd, unix.TIOCSWINSZ, &unix.Winsize{Col: columns, Row: rows}))
	require.NoError(t, syscall.Kill(sandbox.ready.PID, syscall.SIGWINCH))
	return mark
}

func (sandbox *ptySandbox) finish(t *testing.T) ptyResult {
	t.Helper()
	err := <-sandbox.wait
	sandbox.wait = nil // The waiter has been joined, including error exits.
	sandbox.finished = err == nil
	sandbox.cancel()
	require.NoError(t, err, "%s", sandbox.output())
	require.NotContains(t, sandbox.output(), "FAIL", "%s", sandbox.output())
	var result ptyResult
	require.NoError(t, json.Unmarshal([]byte(sandboxRead(t, filepath.Join(sandbox.dir, "result"))), &result), "%s", sandbox.output())
	require.Contains(t, sandbox.output(), "\x18\x1b[39;49m\x1b[2J\x1b8\x1b[?7;25;1004;2004r")
	return result
}

// script provides the controlling terminal and stays attached until the test
// exits. stdin is kept open: EOF would inject EOT and counterfeit a menu key.
func preparePTYChild(t *testing.T) (string, int, *unix.Termios) {
	t.Helper()
	dir := os.Getenv(sandboxEnv)
	fd, err := unix.Open("/dev/tty", unix.O_RDWR|unix.O_NOCTTY, 0)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, unix.Close(fd)) })
	before, err := unix.IoctlGetTermios(fd, unix.TIOCGETA)
	require.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	command := exec.CommandContext(ctx, "/usr/bin/tty")
	command.Stdin = os.Stdin
	name, err := command.Output()
	require.NoError(t, err)
	require.NoError(t, unix.IoctlSetWinsize(fd, unix.TIOCSWINSZ, &unix.Winsize{Col: 80, Row: 12}))
	sandboxJSON(t, filepath.Join(dir, "ready"), ptyReady{PID: os.Getpid(), Name: strings.TrimSpace(string(name)), Original: *before})
	return dir, fd, before
}

func assertPTYRestored(t *testing.T, fd int, before *unix.Termios) {
	t.Helper()
	after, err := unix.IoctlGetTermios(fd, unix.TIOCGETA)
	require.NoError(t, err)
	// Only Darwin's transient PENDIN is excluded; all configured state remains.
	original := *before
	original.Lflag &^= unix.PENDIN
	after.Lflag &^= unix.PENDIN
	require.Equal(t, original, *after)
}

func recordPTYResult(t *testing.T, dir string, event MenuEvent, err error) {
	t.Helper()
	result := ptyResult{Dismissed: event.Dismissed}
	if event.Selection != nil {
		result.Keys, result.Text, result.Confirmed = event.Selection.mapping.Keys, event.Selection.textInput, event.Selection.confirmed
	}
	if err != nil {
		result.Problem = err.Error()
	}
	sandboxJSON(t, filepath.Join(dir, "result"), result)
}

func sandboxCase() string { return os.Getenv("REX_WHICHKEY_PTY_CASE") }

func sandboxID(kind string, index int) string {
	return kind + ":00000000-0000-0000-0000-" + strings.Repeat("0", 12-len(strconv.Itoa(index))) + strconv.Itoa(index)
}

var (
	sandboxSession = sandboxID("session", 1)
	sandboxWindow  = sandboxID("window", 2)
	sandboxOrigin  = sandboxID("block", 3)
	sandboxLayer   = sandboxID("layer", 4)
	sandboxOwn     = sandboxID("block", 5)
	sandboxFloat   = sandboxID("layer", 6)
	sandboxClient  = sandboxID("client", 7)
)
