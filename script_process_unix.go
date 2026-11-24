//go:build darwin || linux

package whichkey

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"syscall"

	"temp/rex"
)

// Verify OS credentials independently of REX_SERVER and the CLI's discovery
// claims. In particular a network endpoint, missing peer PID, other account or
// local socket forwarding process is not the native server described by whoami.
func verifyScriptHost(ctx context.Context, host rex.ScriptHost) (outcome error) {
	address, err := url.Parse(host.Endpoint)
	if err != nil || address.Scheme != "unix" || address.Host != "" || host.PID <= 0 || host.Principal.UID == nil {
		return errors.New("script not started: native Unix host proof unavailable")
	}
	var dialer net.Dialer
	connection, err := dialer.DialContext(ctx, "unix", address.Path)
	if err != nil {
		return errors.New("script not started: native local server socket unavailable")
	}
	defer func() { outcome = errors.Join(outcome, connection.Close()) }()
	local, ok := connection.(*net.UnixConn)
	if !ok {
		return errors.New("script not started: native local socket unavailable")
	}
	raw, err := local.SyscallConn()
	if err != nil {
		return fmt.Errorf("script not started: native socket credentials: %w", err)
	}
	var pid, uid int
	var peerErr error
	if err := raw.Control(func(fd uintptr) { pid, uid, peerErr = scriptSocketPeer(int(fd)) }); err != nil {
		return fmt.Errorf("script not started: native socket control: %w", err)
	}
	if peerErr != nil || pid != host.PID || uid != os.Geteuid() || uid != *host.Principal.UID {
		return errors.New("script not started: native server PID/account does not match local OS peer; no fallback")
	}
	return nil
}

// StartProcess uses Go's close-on-exec handshake: chdir/exec/interpreter errors
// return without a child having executed. There is no shell, PATH search,
// ENOEXEC fallback, context-owned kill or completion wait. Real /dev/null files
// prevent os/exec copy goroutines from tying startup to the menu's terminal.
func startScript(ctx context.Context, script preparedScript) (started bool, outcome error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	null, err := os.OpenFile(os.DevNull, os.O_RDWR, 0)
	if err != nil {
		return false, fmt.Errorf("script not started: open disconnected stdio: %w", err)
	}
	defer func() { outcome = errors.Join(outcome, null.Close()) }()
	argv := append([]string{script.path}, script.args...)
	if contextErr := ctx.Err(); contextErr != nil {
		return false, contextErr
	}
	process, err := os.StartProcess(script.path, argv, &os.ProcAttr{Dir: script.cwd, Env: os.Environ(), Files: []*os.File{null, null, null}, Sys: &syscall.SysProcAttr{Setsid: true}})
	if err != nil {
		return false, scriptProblem(fmt.Errorf("script not started: direct exec %q in %q failed (check cwd, executable permission and shebang interpreter): %w", script.path, script.cwd, err), "direct exec: "+scriptOSCause(err), "Check cwd, executable permission and shebang interpreter.")
	}
	// Reap while this owner remains alive, without waiting on shutdown. Unix
	// reparenting owns reaping if the short-lived menu exits before completion.
	go func() { _, _ = process.Wait() }()
	return true, nil
}
