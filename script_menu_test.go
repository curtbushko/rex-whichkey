//go:build darwin

package whichkey

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"syscall"
	"testing"
	"time"

	"temp/rex"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

const (
	scriptMinimize          = "minimize"
	scriptMissing           = "missing"
	scriptErrorRestore      = "error restore"
	scriptHostID            = "fixture-native-machine-account"
	scriptHostMismatch      = "host mismatch"
	scriptNativeUnavailable = "native unavailable"
	scriptNativePIDMismatch = "native pid mismatch"
	scriptOriginMoved       = "origin moved"
	scriptWorkers           = "workers"
	scriptRemovedBeforeExec = "removed before exec"
	scriptCreatorField      = "creator_name"
	scriptFlavorField       = "flavor"
	scriptTimeoutFlag       = "--timeout"
	scriptClients           = "clients"
	scriptTerminalFlavor    = "com.superlogical.terminal.shell"
	scriptImmediateExit     = "immediate exit"
	scriptStateField        = "state"
	scriptBadInterpreter    = "bad interpreter"
	scriptGenerationChanged = "generation changed"
	scriptDetachUncertain   = "detach uncertain"
	scriptUnknown           = "unknown"
	scriptNoAutostart       = "--autostart=false"
	scriptJSONFlag          = "--json"
	scriptClientID          = "client_id"
	scriptArgsField         = "args"
	scriptLayerID           = "layer_id"
	scriptTerminalCreator   = "com.superlogical.terminal"
	scriptPlaced            = "placed"
	scriptPlacement         = "placement"
	scriptNativeIdentity    = "native identity"
	scriptMissingAccount    = "foreground missing account"
	scriptObsoleteUID       = "foreground obsolete uid"
	scriptObsoleteOverride  = "foreground obsolete override"
	scriptStartLedger       = "starts"
	scriptForegroundAccount = "foreground account"
	scriptCWDUnavailable    = "cwd unavailable"
	scriptAccountProblem    = "original foreground cwd/account unavailable"
	scriptProcessName       = "name"
	scriptMissingParent     = "missing parent"
	scriptMissingParentAbs  = "missing parent absolute"
	scriptPathMissing       = "no such file or directory"
)

type scriptPTYObservation struct {
	PID            int
	CWD            string
	Args           []string
	WorkersDrained bool
}

func TestScriptDirectLifecycle(t *testing.T) {
	if os.Getenv(sandboxEnv) != "" {
		runScriptPTYChild(t)
		return
	}
	for scenario, expectedProblem := range scriptScenarioProblems() {
		t.Run(scenario, func(t *testing.T) {
			sandbox := startPTYSandbox(t, "TestScriptDirectLifecycle", "script "+scenario)
			sandbox.expect(t, 0, "Script fixture action")
			require.NoError(t, os.WriteFile(filepath.Join(sandbox.dir, "selected"), nil, 0600))
			if scenario == scriptWorkers {
				sandbox.send(t, "h")
				sandboxAwait(t, func() bool { return sandboxRead(t, filepath.Join(sandbox.dir, "worker")) != "" }, sandbox.output)
				sandbox.send(t, "\x7fx")
			} else {
				sandbox.send(t, "x")
			}
			// A confirmed script cannot minimize, inspect a script or start on x alone.
			time.Sleep(35 * time.Millisecond)
			require.NotContains(t, sandboxRead(t, filepath.Join(sandbox.dir, "calls")), scriptNativeIdentity)
			sandbox.send(t, "yxy")
			needsAck := scriptAccountRefusal(scenario) || slices.Contains([]string{scriptMissing, scriptMissingParent, scriptMissingParentAbs, scriptPermission, scriptHostMismatch, scriptNativeUnavailable, scriptNativePIDMismatch, scriptCWDUnavailable, scriptOriginMoved, scriptGenerationChanged, scriptBadInterpreter, scriptRemovedBeforeExec}, scenario)
			if needsAck {
				sandbox.expect(t, 0, "Script not started")
				// Repeated selection/confirmation input cannot cause a retry or action.
				sandbox.send(t, "xyy")
				time.Sleep(20 * time.Millisecond)
				sandbox.send(t, "q")
			}
			result := sandbox.finish(t)
			t.Logf("result: %s", result.Problem)
			calls := scriptReadCalls(t, sandbox)
			operations := make([]string, 0, len(calls))
			for _, call := range calls {
				operations = append(operations, call.Operation)
				if call.Operation == scriptMinimize || call.Operation == fixtureFocus || call.Operation == fixtureDetach || call.Operation == scriptErrorRestore {
					require.True(t, call.Restored, call.Operation)
					require.True(t, call.WorkersDrained, call.Operation)
				}
			}
			require.NotContains(t, operations, scriptUnknown)
			started := slices.Contains([]string{"success", scriptWorkers, scriptImmediateExit, scriptDetachUncertain}, scenario)
			if started {
				assertScriptStarted(t, sandbox, scenario, operations, result)
			} else {
				assertScriptRefusal(t, scenario, expectedProblem, result.Problem, operations)
				require.Empty(t, sandboxRead(t, filepath.Join(sandbox.dir, "effect")))
				require.Empty(t, sandboxRead(t, filepath.Join(sandbox.dir, scriptStartLedger)))
				if scenario == scriptBadInterpreter || scenario == scriptRemovedBeforeExec {
					require.Equal(t, 1, countOperation(operations, scriptErrorRestore))
					require.Contains(t, sandbox.output(), "direct exec")
				} else {
					require.NotContains(t, operations, scriptErrorRestore)
				}
			}
			if needsAck {
				require.Contains(t, sandbox.output(), "Esc/q: close; no retry")
			}
		})
	}
}

func TestScriptShortErrorLifecycle(t *testing.T) {
	if os.Getenv(sandboxEnv) != "" {
		runScriptPTYChild(t)
		return
	}
	for _, height := range []uint16{5, 7} {
		for scenario, expected := range map[string]string{
			scriptMissing:           scriptPathMissing,
			scriptBadInterpreter:    scriptPathMissing,
			scriptPermission:        "Target is not executable",
			scriptNonregular:        "regular executable script",
			scriptCWDUnavailable:    scriptPathMissing,
			scriptHostMismatch:      "proven native machine/account",
			scriptForegroundAccount: "original foreground cwd/account unavailable",
		} {
			t.Run(fmt.Sprintf("%s/%d", scenario, height), func(t *testing.T) {
				sandbox := startPTYSandbox(t, "TestScriptShortErrorLifecycle", "script short "+scenario)
				sandbox.expect(t, 0, "Script fixture action")
				sandbox.resize(t, 64, height)
				sandbox.send(t, "x")
				sandbox.expect(t, 0, "y=yes / other=cancel")
				mark := sandbox.send(t, "y")
				sandbox.expect(t, mark, "Esc/q: close; no retry")
				rows := scriptDrawnRows(t, sandbox.output()[mark:], 64, int(height))
				t.Logf("drawn before acknowledgement: %q", rows)
				drawn := strings.Join(rows, " ")
				require.Contains(t, drawn, expected)
				require.Contains(t, drawn, scriptShortRecovery(scenario))
				require.Contains(t, drawn, "Esc/q: close; no retry")
				require.Empty(t, sandboxRead(t, filepath.Join(sandbox.dir, "result")), "cause must be visible while acknowledgement is still open")
				sandbox.send(t, "xyy")
				time.Sleep(20 * time.Millisecond)
				sandbox.send(t, "q")
				result := sandbox.finish(t)
				require.NotEmpty(t, result.Problem)
				if scenario != scriptHostMismatch && scenario != scriptForegroundAccount {
					require.Contains(t, result.Problem, scriptShortCWD(sandbox.dir), "full diagnostic is retained for callers")
				}
				operations := make([]string, 0)
				for _, call := range scriptReadCalls(t, sandbox) {
					operations = append(operations, call.Operation)
					if slices.Contains([]string{scriptMinimize, fixtureFocus, fixtureDetach, scriptErrorRestore}, call.Operation) {
						require.True(t, call.Restored)
						require.True(t, call.WorkersDrained)
					}
				}
				require.Equal(t, 1, countOperation(operations, fixtureDetach))
				if scenario == scriptBadInterpreter {
					require.Equal(t, 1, countOperation(operations, scriptErrorRestore))
					require.Equal(t, 2, countOperation(operations, "foreground cwd"))
				} else {
					require.NotContains(t, operations, scriptMinimize)
					reads := 1
					if scenario == scriptHostMismatch {
						reads = 0
					}
					require.Equal(t, reads, countOperation(operations, "foreground cwd"))
				}
				require.Empty(t, sandboxRead(t, filepath.Join(sandbox.dir, scriptStartLedger)))
				require.Empty(t, sandboxRead(t, filepath.Join(sandbox.dir, "effect")))
			})
		}
	}
}

func scriptShortRecovery(scenario string) string {
	switch scenario {
	case scriptBadInterpreter:
		return scriptRepairExec
	case scriptPermission:
		return "Set executable permission and a valid shebang."
	case scriptNonregular:
		return "Choose a regular file, not a directory or device."
	case scriptCWDUnavailable:
		return scriptRepairCWD
	case scriptHostMismatch, scriptForegroundAccount:
		return "Check origin host/account/cwd and trusted config."
	default:
		return "Check the literal path in the origin cwd."
	}
}

func scriptShortCWD(dir string) string {
	return filepath.Join(dir, strings.Repeat("long-origin-Ω-", 12))
}

func scriptScenarioProblems() map[string]string {
	return map[string]string{
		"success":               "",
		scriptWorkers:           "",
		scriptImmediateExit:     "",
		scriptDetachUncertain:   "script started",
		scriptMissing:           scriptPathMissing,
		scriptMissingParent:     scriptPathMissing,
		scriptMissingParentAbs:  scriptPathMissing,
		scriptPermission:        "is not executable",
		scriptHostMismatch:      "must share the proven native machine/account",
		scriptNativeUnavailable: "native local host/account discovery failed",
		scriptNativePIDMismatch: "does not match local OS peer",
		scriptCWDUnavailable:    scriptAccountProblem,
		scriptForegroundAccount: scriptAccountProblem,
		scriptMissingAccount:    scriptAccountProblem,
		scriptObsoleteUID:       scriptAccountProblem,
		scriptObsoleteOverride:  scriptAccountProblem,
		scriptOriginMoved:       "origin placement or explicit app connection changed",
		scriptGenerationChanged: "explicit app connection changed",
		"minimize uncertain":    "popup minimize failed or uncertain",
		"minimize ack only":     "not verified minimized",
		"unknown state":         "not verified minimized",
		"moved popup":           "no longer contains exactly its owned terminal",
		"focus uncertain":       "original focus restoration failed or is uncertain",
		"cancel minimized":      "popup minimize failed or uncertain",
		scriptBadInterpreter:    "direct exec",
		scriptRemovedBeforeExec: "direct exec",
	}
}

func scriptAccountRefusal(scenario string) bool {
	return slices.Contains([]string{scriptForegroundAccount, scriptMissingAccount, scriptObsoleteUID, scriptObsoleteOverride}, scenario)
}

func assertScriptRefusal(t *testing.T, scenario, expected, problem string, operations []string) {
	t.Helper()
	require.NotEmpty(t, expected, "each refusal scenario needs a specific guard assertion")
	require.Contains(t, problem, expected, "the intended guard, not an unrelated failure, must refuse startup")
	if slices.Contains([]string{scriptHostMismatch, scriptNativeUnavailable, scriptNativePIDMismatch, scriptOriginMoved, scriptGenerationChanged}, scenario) {
		require.NotContains(t, operations, "foreground cwd", "identity/placement failures must precede target filesystem access")
	}
	if scenario == scriptMissingParent || scenario == scriptMissingParentAbs {
		require.Equal(t, 1, countOperation(operations, "foreground cwd"), "missing intermediate path must refuse on first preflight")
		require.NotContains(t, operations, scriptMinimize, "preflight refusal must precede native startup")
	}
	if scriptAccountRefusal(scenario) {
		require.Equal(t, 1, countOperation(operations, "foreground cwd"), "account refusal must not retry preflight")
		require.NotContains(t, operations, scriptMinimize, "account proof must precede target preflight and startup")
	}
	if slices.Contains([]string{"minimize uncertain", "minimize ack only", "unknown state", "moved popup", "cancel minimized"}, scenario) {
		require.Equal(t, 1, countOperation(operations, "foreground cwd"), "no final script preparation or retry after failed invisibility")
	}
}

func assertScriptStarted(t *testing.T, sandbox *ptySandbox, scenario string, operations []string, result ptyResult) {
	t.Helper()
	if scenario != scriptDetachUncertain {
		require.Empty(t, result.Problem)
	} else {
		require.Contains(t, result.Problem, "script started")
	}
	require.Equal(t, 1, countOperation(operations, scriptMinimize))
	require.Equal(t, 1, countOperation(operations, fixtureFocus))
	require.Equal(t, 1, countOperation(operations, fixtureDetach))
	require.NotContains(t, operations, scriptErrorRestore)
	require.Less(t, slices.Index(operations, "inspect minimized"), slices.Index(operations, fixtureFocus))
	require.Greater(t, slices.Index(operations, fixtureDetach), slices.Index(operations, fixtureFocus))
	if scenario != scriptImmediateExit {
		sandboxAwait(t, func() bool { return sandboxRead(t, filepath.Join(sandbox.dir, "effect")) != "" }, sandbox.output)
		var observation scriptPTYObservation
		require.NoError(t, json.Unmarshal([]byte(sandboxRead(t, filepath.Join(sandbox.dir, "effect"))), &observation))
		t.Cleanup(func() { _ = syscall.Kill(observation.PID, syscall.SIGKILL) })
		resolved, err := filepath.EvalSymlinks(sandbox.dir)
		require.NoError(t, err)
		require.Equal(t, resolved, observation.CWD)
		require.Equal(t, []string{"", scriptSpaceArg, scriptLiteralHome, "*", "; touch injection", "Ω"}, observation.Args)
		// The controlling PTY may already be gone. Restoration is checked
		// independently at minimize, focus and detach above, before exec.
		require.True(t, observation.WorkersDrained)
		require.NoError(t, syscall.Kill(observation.PID, 0), "survived menu process exit and PTY teardown")
		require.Empty(t, sandboxRead(t, filepath.Join(sandbox.dir, "complete")))
		require.NoError(t, os.WriteFile(filepath.Join(sandbox.dir, "release"), []byte("release"), 0600))
		sandboxAwait(t, func() bool { return sandboxRead(t, filepath.Join(sandbox.dir, "complete")) == "done" }, sandbox.output)
		require.Equal(t, 1, strings.Count(sandboxRead(t, filepath.Join(sandbox.dir, scriptStartLedger)), "\n"), "exactly one native child startup")
	}
}

func scriptReadCalls(t *testing.T, sandbox *ptySandbox) []directFixtureCall {
	t.Helper()
	var calls []directFixtureCall
	for line := range strings.SplitSeq(strings.TrimSpace(sandboxRead(t, filepath.Join(sandbox.dir, "calls"))), "\n") {
		var call directFixtureCall
		require.NoError(t, json.Unmarshal([]byte(line), &call))
		calls = append(calls, call)
	}
	return calls
}

func runScriptPTYChild(t *testing.T) {
	t.Helper()
	dir, fd, before := preparePTYChild(t)
	socketDir, err := os.MkdirTemp("/tmp", "rws-")
	require.NoError(t, err)
	defer func() { require.NoError(t, os.RemoveAll(socketDir)) }()
	listener, err := net.ListenUnix(scriptFixtureUnix, &net.UnixAddr{Name: filepath.Join(socketDir, "native.sock"), Net: scriptFixtureUnix})
	require.NoError(t, err)
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			connection, acceptErr := listener.Accept()
			if acceptErr != nil {
				return
			}
			// Let the client read OS peer credentials before closing the peer.
			// Immediate close can race Darwin LOCAL_PEERPID even though the
			// listener and advertised PID belong to this same fixture process.
			_, _ = io.Copy(io.Discard, connection)
			_ = connection.Close()
		}
	}()
	defer func() { require.NoError(t, listener.Close()); <-done }()
	uid := os.Geteuid()
	host := rex.ScriptHost{Endpoint: (&url.URL{Scheme: scriptFixtureUnix, Path: listener.Addr().String()}).String(), PID: os.Getpid(), HostID: scriptHostID}
	host.Principal.Kind = "peer"
	host.Principal.Transport = scriptFixtureUnix
	host.Principal.UID = &uid
	if sandboxCase() == "script native pid mismatch" {
		host.PID++
	}
	sandboxJSON(t, filepath.Join(dir, "native"), host)
	executable, err := os.Executable()
	require.NoError(t, err)
	cwd := dir
	if strings.HasPrefix(sandboxCase(), "script short ") {
		cwd = scriptShortCWD(dir)
		require.NoError(t, os.MkdirAll(cwd, 0700))
	}
	script := filepath.Join(cwd, "run Ω; $HOME")
	content := "#!/bin/sh\nset -eu\nexec \"$@\"\n"
	if sandboxCase() == "script bad interpreter" || sandboxCase() == "script short bad interpreter" {
		content = "#!/definitely/missing/interpreter\n"
	}
	if sandboxCase() == "script immediate exit" {
		content = "#!/bin/sh\nset -eu\nexit 17\n"
	}
	require.NoError(t, os.WriteFile(script, []byte(content), 0700))
	if sandboxCase() == "script missing" || sandboxCase() == "script short missing" {
		script = filepath.Join(cwd, scriptMissing)
	}
	if sandboxCase() == "script permission" || sandboxCase() == "script short permission" {
		require.NoError(t, os.Chmod(script, 0600))
	}
	if sandboxCase() == "script short nonregular" {
		script = cwd
	}
	configuredPath := filepath.Base(script)
	if sandboxCase() == "script short nonregular" {
		configuredPath = script
	}
	if sandboxCase() == "script "+scriptMissingParent || sandboxCase() == "script "+scriptMissingParentAbs {
		// Keep the nonexistent parent and .. literal, with a real executable at
		// the lexically cleaned path that would record any unintended launch.
		configuredPath = "missing/../" + configuredPath
		if sandboxCase() == "script "+scriptMissingParentAbs {
			configuredPath = dir + "/" + configuredPath
		}
	}
	args, err := json.Marshal(map[string]any{"path": configuredPath, scriptArgsField: []string{executable, "-test.run=^TestScriptPTYEffect$", "--", "", scriptSpaceArg, scriptLiteralHome, "*", "; touch injection", "Ω"}})
	require.NoError(t, err)
	spec := []Mapping{{Keys: "x", Desc: "Script fixture action", Action: menuScriptRun, Args: args, Confirm: true}}
	if sandboxCase() == "script workers" {
		group := "Hosts"
		spec = append(spec, Mapping{Keys: "h", Provider: menuProviderHosts, Group: &group})
		require.NoError(t, unix.Mkfifo(filepath.Join(dir, "worker-fifo"), 0600))
	}
	config := filepath.Join(dir, "config.json")
	sandboxJSON(t, config, map[string]any{"version": 1, "title": "Script fixture action", "spec": spec})
	endpoint, client := sandboxEndpoint, sandboxClient
	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
	defer cancel()
	err = RunDirectMenu(ctx, DirectOptions{OriginSession: sandboxSession, OriginBlock: sandboxOrigin, OriginWindow: sandboxWindow, OriginLayer: sandboxLayer, Server: &endpoint, Client: &client, Config: &config, NoColor: true, ASCII: true})
	assertPTYRestored(t, fd, before)
	recordPTYResult(t, dir, MenuEvent{}, err)
}

func TestScriptPTYEffect(t *testing.T) {
	dir := os.Getenv(sandboxEnv)
	if dir == "" {
		return
	}
	cwd, err := os.Getwd()
	require.NoError(t, err)
	observation := scriptPTYObservation{PID: os.Getpid(), CWD: cwd, WorkersDrained: fixtureWorkersDrained(dir)}
	for index, arg := range os.Args {
		if arg == "--" {
			observation.Args = os.Args[index+1:]
			break
		}
	}
	starts, err := os.OpenFile(filepath.Join(dir, scriptStartLedger), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	require.NoError(t, err)
	require.NoError(t, json.NewEncoder(starts).Encode(observation))
	require.NoError(t, starts.Close())
	sandboxJSON(t, filepath.Join(dir, "effect"), observation)
	deadline := time.Now().Add(10 * time.Second)
	for time.Now().Before(deadline) {
		if sandboxRead(t, filepath.Join(dir, "release")) != "" {
			require.NoError(t, os.WriteFile(filepath.Join(dir, "complete"), []byte("done"), 0600))
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("fixture not released")
}

func scriptFixtureResponse(argv []string, state *directFixtureState) (string, any, int, bool) {
	if !strings.HasPrefix(sandboxCase(), "script ") {
		return "", nil, 0, false
	}
	dir := os.Getenv(sandboxEnv)
	if slices.Equal(argv, []string{scriptNoAutostart, scriptTimeoutFlag, "2s", "whoami", scriptJSONFlag}) {
		if os.Getenv("REX_SERVER") != "" || os.Getenv("REX_SESSION") != "" || os.Getenv("REX_BLOCK") != "" || os.Getenv("REX_CLIENT") != "" {
			return "selectors leaked", nil, 98, true
		}
		if sandboxCase() == "script native unavailable" {
			return scriptNativeIdentity, nil, 1, true
		}
		var host rex.ScriptHost
		data, err := os.ReadFile(filepath.Join(dir, "native"))
		if err != nil || json.Unmarshal(data, &host) != nil {
			return scriptNativeIdentity, nil, 1, true
		}
		return scriptNativeIdentity, host, 0, true
	}
	if slices.Equal(argv, []string{scriptNoAutostart, scriptTimeoutFlag, "3s", "hosts", scriptJSONFlag}) && sandboxCase() == "script workers" {
		operation, result, code := fixturePendingWorker()
		return operation, result, code, true
	}
	if slices.Equal(argv, fixtureArgs("session", "inspect", scriptJSONFlag, "--", sandboxSession)) {
		operation := "inspect placed"
		if state.Minimized {
			operation = "inspect minimized"
		}
		if state.Detached {
			operation = "inspect retired"
		}
		return operation, scriptFixtureView(*state), 0, true
	}
	for _, id := range []string{sandboxOrigin, sandboxOwn} {
		if slices.Equal(argv, fixtureAPI("session.describe_block", map[string]string{menuSessionID: sandboxSession, fixtureBlockID: id})) {
			host := scriptHostID
			if sandboxCase() == "script host mismatch" || sandboxCase() == "script short host mismatch" {
				host = "remote"
			}
			return "describe " + id, map[string]any{"block": map[string]any{fixtureBlockID: id, "host_id": host, "window_id": sandboxWindow, scriptPlacement: scriptPlaced, scriptCreatorField: scriptTerminalCreator, scriptFlavorField: scriptTerminalFlavor}}, 0, true
		}
	}
	processArgs, _ := json.Marshal(map[string]any{menuSessionID: sandboxSession, fixtureBlockID: sandboxOrigin, scriptArgsField: map[string]any{}})
	if slices.Equal(argv, fixtureArgs("api", "call", "--session", sandboxSession, "--", "com.superlogical.terminal.process", string(processArgs))) {
		state.ScriptReads++
		return "foreground cwd", scriptFixtureProcess(dir), 0, true
	}
	for _, minimized := range []bool{true, false} {
		value := "normal"
		params := map[string]any{menuSessionID: sandboxSession, scriptLayerID: sandboxFloat, scriptStateField: value, "focus": true}
		operation := scriptErrorRestore
		if minimized {
			value = "minimized"
			params = map[string]any{menuSessionID: sandboxSession, scriptLayerID: sandboxFloat, scriptStateField: value}
			operation = scriptMinimize
		}
		encoded, _ := json.Marshal(params)
		if slices.Equal(argv, fixtureArgs("api", "call", "--session", sandboxSession, "--", "session.set_layer_state", string(encoded))) {
			code := scriptFixtureLayerState(state, minimized)
			return operation, map[string]int{"revision": 1}, code, true
		}
	}
	if slices.Equal(argv, fixtureAPI("session.focus_block", map[string]string{menuSessionID: sandboxSession, fixtureBlockID: sandboxOrigin})) {
		state.Focused = true
		if sandboxCase() == "script focus uncertain" {
			return fixtureFocus, nil, 1, true
		}
		return fixtureFocus, map[string]any{}, 0, true
	}
	if slices.Equal(argv, fixtureAPI("session.detach_block", map[string]string{menuSessionID: sandboxSession, fixtureBlockID: sandboxOwn})) {
		state.Detached = true
		if sandboxCase() == "script detach uncertain" {
			return fixtureDetach, nil, 1, true
		}
		return fixtureDetach, map[string]any{}, 0, true
	}
	if slices.Equal(argv, fixtureAPI("session.list_clients", map[string]string{menuSessionID: sandboxSession})) && sandboxCase() == "script generation changed" {
		generation := "2026-11-07T00:00:00Z"
		if _, err := os.Stat(filepath.Join(dir, "selected")); err == nil {
			generation = "2026-11-07T01:00:00Z"
		}
		return scriptClients, map[string]any{scriptClients: []any{map[string]any{scriptClientID: sandboxClient, "connected_at": generation, "connection_state": "connected", "info": map[string]string{fixtureKind: "app"}}}}, 0, true
	}
	if slices.Equal(argv, fixtureAPI("session.list_blocks", map[string]string{menuSessionID: sandboxSession})) && state.ScriptReads == 2 && sandboxCase() == "script removed before exec" {
		source := filepath.Join(dir, "run Ω; $HOME")
		if _, err := os.Stat(source); err == nil {
			_ = os.Rename(source, source+".moved")
		}
	}
	return "", nil, 0, false
}

// Match terminal.process's native shape, not whoami's principal.uid schema.
// A valid child must not become fallback account/cwd evidence for a bad foreground.
func scriptFixtureProcess(dir string) map[string]any {
	if strings.HasPrefix(sandboxCase(), "script short ") {
		dir = scriptShortCWD(dir)
	}
	foreground := map[string]any{
		"pid": os.Getpid(), "user_id": os.Geteuid(), "user": "fixture",
		scriptProcessName: "sh", "path": "/bin/sh", "invoked_path": "/bin/sh", "argv0": "sh", "cwd": dir,
	}
	result := map[string]any{"last_exit": nil, "child": maps.Clone(foreground), "foreground": foreground, "home": dir}
	scenario := strings.TrimPrefix(strings.TrimPrefix(sandboxCase(), "script "), "short ")
	if sandboxCase() == "script short cwd unavailable" {
		foreground["cwd"] = filepath.Join(dir, "missing-cwd")
		return result
	}
	if scenario == scriptCWDUnavailable {
		delete(foreground, "cwd")
	}
	if scriptAccountRefusal(scenario) {
		// If account proof is bypassed, prepareScript will fail its first Stat
		// with a different diagnostic. The asserted account error must win.
		foreground["cwd"] = filepath.Join(dir, "must-not-stat-without-account")
	}
	switch scenario {
	case scriptForegroundAccount:
		foreground["user_id"] = os.Geteuid() + 1
	case scriptMissingAccount:
		delete(foreground, "user_id")
	case scriptObsoleteUID:
		delete(foreground, "user_id")
		foreground["uid"] = os.Geteuid()
	case scriptObsoleteOverride:
		foreground["user_id"] = os.Geteuid() + 1
		foreground["uid"] = os.Geteuid()
	}
	return result
}

func scriptFixtureLayerState(state *directFixtureState, minimized bool) int {
	state.Minimized = minimized
	if !minimized {
		state.Focused = false
		return 0
	}
	switch sandboxCase() {
	case "script minimize uncertain":
		return 1
	case "script minimize ack only":
		state.Minimized = false
	case "script cancel minimized":
		var ready ptyReady
		data, _ := os.ReadFile(filepath.Join(os.Getenv(sandboxEnv), "ready"))
		_ = json.Unmarshal(data, &ready)
		_ = syscall.Kill(ready.PID, syscall.SIGTERM)
	}
	return 0
}

func scriptFixtureView(state directFixtureState) any {
	layer := sandboxLayer
	if sandboxCase() == "script origin moved" {
		if _, err := os.Stat(filepath.Join(os.Getenv(sandboxEnv), "selected")); err == nil {
			layer = sandboxID("layer", 9)
		}
	}
	layers := []any{map[string]any{scriptLayerID: layer, fixtureKind: "tiling", fixtureBlocks: []any{map[string]string{fixtureBlockID: sandboxOrigin}}}}
	if !state.Detached {
		value := ""
		if state.Minimized {
			value = "minimized"
		}
		if state.Minimized && sandboxCase() == "script unknown state" {
			value = scriptUnknown
		}
		window := map[string]any{scriptLayerID: sandboxFloat, fixtureKind: "floating", scriptStateField: value, fixtureBlocks: []any{map[string]string{fixtureBlockID: sandboxOwn}}}
		if state.Minimized && sandboxCase() == "script moved popup" {
			window[fixtureBlocks] = []any{map[string]string{fixtureBlockID: sandboxID("block", 9)}}
		}
		layers = append(layers, window)
	}
	focus := sandboxOwn
	if state.Focused {
		focus = sandboxOrigin
	}
	return map[string]any{menuSessionID: sandboxSession, "active_window_id": sandboxWindow, "windows": []any{map[string]any{menuWindowID: sandboxWindow, "focused_block_id": focus, "layers": layers}}}
}
