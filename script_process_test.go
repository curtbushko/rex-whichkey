//go:build darwin || linux

package whichkey

import (
	"context"
	"encoding/json"
	"net"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"testing"
	"time"

	"temp/rex"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

const (
	scriptFixtureEnv  = "REX_WHICHKEY_SCRIPT_FIXTURE"
	scriptLiteralHome = "$HOME"
	scriptSpaceArg    = "a b"
	scriptFixtureUnix = "unix"
)

type detachedObservation struct {
	PID, Session int
	CWD          string
	Args         []string
	Null         []bool
}

func TestScriptDetachedLifetime(t *testing.T) {
	dir := t.TempDir()
	script := filepath.Join(dir, "script Ω; $(touch injected)")
	// This fixture forwards literal argv to the real OS child, never through eval.
	require.NoError(t, os.WriteFile(script, []byte("#!/bin/sh\nset -eu\nexec \"$@\"\n"), 0700))
	executable, err := os.Executable()
	require.NoError(t, err)
	launcher := exec.CommandContext(context.Background(), executable, "-test.run=^TestScriptLauncherChild$")
	launcher.Env = append(os.Environ(), scriptFixtureEnv+"="+dir, "GORACE=atexit_sleep_ms=0")
	output, err := launcher.CombinedOutput()
	require.NoError(t, err, string(output))
	var observation detachedObservation
	require.Eventually(t, func() bool {
		data, readErr := os.ReadFile(filepath.Join(dir, "observed"))
		return readErr == nil && json.Unmarshal(data, &observation) == nil
	}, 5*time.Second, 5*time.Millisecond)
	t.Cleanup(func() { _ = syscall.Kill(observation.PID, syscall.SIGKILL) })
	require.Equal(t, observation.PID, observation.Session, "separate OS session, no controlling menu PTY")
	resolved, err := filepath.EvalSymlinks(dir)
	require.NoError(t, err)
	require.Equal(t, resolved, observation.CWD)
	require.Equal(t, []string{"", scriptSpaceArg, scriptLiteralHome, ";touch injected", "*", "Ω", "\"'"}, observation.Args)
	require.Equal(t, []bool{true, true, true}, observation.Null, "all streams are real /dev/null handles")
	require.NoError(t, syscall.Kill(observation.PID, 0), "child still alive after launcher exited")
	_, err = os.Stat(filepath.Join(dir, "completed"))
	require.ErrorIs(t, err, os.ErrNotExist, "launcher must not wait for completion")
	require.NoError(t, os.WriteFile(filepath.Join(dir, "release"), nil, 0600))
	require.Eventually(t, func() bool { _, statErr := os.Stat(filepath.Join(dir, "completed")); return statErr == nil }, 5*time.Second, 5*time.Millisecond)
}

func TestScriptLauncherChild(t *testing.T) {
	dir := os.Getenv(scriptFixtureEnv)
	if dir == "" {
		return
	}
	executable, err := os.Executable()
	require.NoError(t, err)
	prepared, err := prepareScript(scriptAction{path: "script Ω; $(touch injected)", args: append([]string{executable, "-test.run=^TestScriptDetachedChild$", "--"}, "", scriptSpaceArg, scriptLiteralHome, ";touch injected", "*", "Ω", "\"'")}, dir)
	require.NoError(t, err)
	started, err := startScript(context.Background(), prepared)
	require.NoError(t, err)
	require.True(t, started)
}

func TestScriptDetachedChild(t *testing.T) {
	dir := os.Getenv(scriptFixtureEnv)
	if dir == "" {
		return
	}
	cwd, err := os.Getwd()
	require.NoError(t, err)
	session, err := unix.Getsid(0)
	require.NoError(t, err)
	observation := detachedObservation{PID: os.Getpid(), Session: session, CWD: cwd}
	for index, arg := range os.Args {
		if arg == "--" {
			observation.Args = os.Args[index+1:]
			break
		}
	}
	null, err := os.Stat(os.DevNull)
	require.NoError(t, err)
	for _, stream := range []*os.File{os.Stdin, os.Stdout, os.Stderr} {
		info, statErr := stream.Stat()
		observation.Null = append(observation.Null, statErr == nil && os.SameFile(null, info))
	}
	data, err := json.Marshal(observation)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(dir, "observed"), data, 0600))
	deadline := time.Now().Add(7 * time.Second)
	for time.Now().Before(deadline) {
		if _, err := os.Stat(filepath.Join(dir, "release")); err == nil {
			require.NoError(t, os.WriteFile(filepath.Join(dir, "completed"), nil, 0600))
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("parent did not release detached child")
}

func TestScriptLiteralTraversal(t *testing.T) {
	for _, spelling := range []struct {
		name, cwdSuffix, path string
		absolute              bool
	}{
		{name: "relative", path: "link/../run.sh"},
		{name: "absolute", absolute: true},
		{name: "trailing cwd separator", cwdSuffix: "/", path: "link/../run.sh"},
		{name: "cwd traversal", cwdSuffix: "/link/..", path: "run.sh"},
	} {
		t.Run(spelling.name, func(t *testing.T) {
			dir := t.TempDir()
			cwd := filepath.Join(dir, "work")
			target := filepath.Join(dir, "target")
			require.NoError(t, os.Mkdir(cwd, 0700))
			require.NoError(t, os.MkdirAll(filepath.Join(target, "sub"), 0700))
			require.NoError(t, os.Symlink(filepath.Join(target, "sub"), filepath.Join(cwd, "link")))
			// Distinct executables make lexical cleaning observable at native exec.
			for directory, marker := range map[string]string{cwd: "lexical", target: "filesystem"} {
				content := "#!/bin/sh\nset -eu\nprintf '%s' " + marker + " >\"$1\"\n"
				require.NoError(t, os.WriteFile(filepath.Join(directory, "run.sh"), []byte(content), 0700))
			}
			// Do not use filepath.Join here: it would erase the regression input.
			path := spelling.path
			literalPath := cwd + "/link/../run.sh"
			if spelling.absolute {
				path = literalPath
			}
			observed := filepath.Join(dir, "observed")
			prepared, err := prepareScript(scriptAction{path: path, args: []string{observed}}, cwd+spelling.cwdSuffix)
			require.NoError(t, err)
			require.Equal(t, literalPath, prepared.path, "anchoring must not canonicalize symlinks or dot-dot components")
			started, err := startScript(context.Background(), prepared)
			require.NoError(t, err)
			require.True(t, started)
			var marker []byte
			require.Eventually(t, func() bool {
				marker, err = os.ReadFile(observed)
				return err == nil && len(marker) > 0
			}, 5*time.Second, 5*time.Millisecond)
			require.Equal(t, "filesystem", string(marker), "native exec must follow the configured filesystem traversal")
		})
	}
}

func TestScriptMissingTraversal(t *testing.T) {
	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, "run.sh"), []byte("#!/bin/sh\nexit 0\n"), 0700))
	for _, path := range []string{"missing/../run.sh", dir + "/missing/../run.sh"} {
		t.Run(path, func(t *testing.T) {
			prepared, err := prepareScript(scriptAction{path: path}, dir)
			require.ErrorIs(t, err, os.ErrNotExist, "an absent intermediate directory cannot be cleaned away")
			require.Empty(t, prepared, "refusal must not return a launchable target")
		})
	}
}

func TestScriptNativeHostPeer(t *testing.T) {
	path := filepath.Join(t.TempDir(), "native.sock")
	listener, err := net.ListenUnix(scriptFixtureUnix, &net.UnixAddr{Name: path, Net: scriptFixtureUnix})
	require.NoError(t, err)
	defer func() { require.NoError(t, listener.Close()) }()
	// OS socket peer identity is independent of CLI JSON and inherited selectors.
	uid := os.Geteuid()
	host := rex.ScriptHost{Endpoint: (&url.URL{Scheme: scriptFixtureUnix, Path: path}).String(), HostID: "local", PID: os.Getpid()}
	host.Principal.UID = &uid
	require.NoError(t, verifyScriptHost(context.Background(), host))
	host.PID++
	require.Error(t, verifyScriptHost(context.Background(), host))
	host.PID = os.Getpid()
	other := uid + 1
	host.Principal.UID = &other
	require.Error(t, verifyScriptHost(context.Background(), host))
	host.Endpoint = "127.0.0.1:" + strconv.Itoa(12345)
	require.Error(t, verifyScriptHost(context.Background(), host))
}

func TestScriptStartupFailures(t *testing.T) {
	dir := t.TempDir()
	for name, content := range map[string]string{"bad-interpreter": "#!/definitely/missing/interpreter\n", "no-shebang": "touch must-not-exist\n"} {
		path := filepath.Join(dir, name)
		require.NoError(t, os.WriteFile(path, []byte(content), 0700))
		prepared, err := prepareScript(scriptAction{path: path}, dir)
		require.NoError(t, err)
		started, err := startScript(context.Background(), prepared)
		require.False(t, started)
		require.Error(t, err)
	}
	path := filepath.Join(dir, "removed")
	require.NoError(t, os.WriteFile(path, []byte("#!/bin/sh\nexit 0\n"), 0700))
	prepared, err := prepareScript(scriptAction{path: path}, dir)
	require.NoError(t, err)
	require.NoError(t, os.Rename(path, path+".moved"))
	started, err := startScript(context.Background(), prepared)
	require.False(t, started)
	require.ErrorIs(t, err, os.ErrNotExist)
	prepared.path = path + ".moved"
	prepared.cwd = filepath.Join(dir, "missing-cwd")
	started, err = startScript(context.Background(), prepared)
	require.False(t, started)
	require.ErrorIs(t, err, os.ErrNotExist)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	started, err = startScript(ctx, prepared)
	require.False(t, started)
	require.ErrorIs(t, err, context.Canceled)
	_, err = os.Stat(filepath.Join(dir, "must-not-exist"))
	require.ErrorIs(t, err, os.ErrNotExist)
	fifo := filepath.Join(dir, "fifo")
	require.NoError(t, unix.Mkfifo(fifo, 0700))
	_, err = prepareScript(scriptAction{path: fifo}, dir)
	require.Error(t, err)
}
