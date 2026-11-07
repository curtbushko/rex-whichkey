//go:build darwin

package whichkey

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

const (
	fixtureDismissed          = "dismissed"
	fixtureBlocks             = "blocks"
	sandboxEndpoint           = "unix:///tmp/rex-whichkey-isolated-test.sock"
	fixtureAppSelected        = "app selected"
	fixtureWorkersDrainedCase = "workers drained"
	fixtureEffect             = "selected effect"
	fixtureFocus              = "restore focus"
	fixtureDetach             = "detach"
	fixtureKind               = "kind"
	fixtureBlockID            = "block_id"
	fixtureWorkerFailure      = "worker failure"
)

type directFixtureState struct {
	Detached bool
	Focused  bool
}

type directFixtureCall struct {
	Operation      string
	Argv           []string
	Restored       bool
	WorkersDrained bool
}

type directLifecycleCase struct {
	name    string
	problem string
	effects int
	focus   int
}

func TestDirectLifecycle(t *testing.T) {
	if os.Getenv(sandboxEnv) != "" {
		runDirectPTYChild(t)
		return
	}
	for _, test := range []directLifecycleCase{
		{"selected", "", 1, 1},
		{fixtureAppSelected, "", 1, 1},
		{fixtureWorkersDrainedCase, "", 1, 1},
		{"effect uncertain", "native resource close failed or is uncertain", 1, 1},
		{"detach uncertain", "popup detach failed or is uncertain", 0, 0},
		{"visible after detach", "detached popup layer is still visible", 0, 0},
		{"still placed", "own terminal placement or flavor could not be verified", 0, 0},
		{"focus uncertain", "original focus restoration failed or is uncertain", 0, 1},
		{"focus unverified", "exact original focus could not be verified", 0, 1},
		{"moved origin", "origin placement or explicit app connection changed", 0, 0},
		{"moved after focus", "original placement changed during focus restoration", 0, 1},
		{"generation changed", "explicit app connection changed", 0, 0},
		{"generation after focus", "origin or app connection changed", 0, 1},
		{fixtureDismissed, "", 0, 1},
		{"canceled", "context canceled", 0, 1},
	} {
		t.Run(test.name, func(t *testing.T) {
			sandbox := startPTYSandbox(t, "TestDirectLifecycle", test.name)
			sandbox.expect(t, 0, "Lifecycle action")
			switch test.name {
			case fixtureDismissed:
				sandbox.send(t, "q")
			case "canceled":
				require.NoError(t, syscall.Kill(sandbox.ready.PID, syscall.SIGTERM))
			case fixtureWorkersDrainedCase:
				sandbox.send(t, "h")
				sandboxAwait(t, func() bool { return sandboxRead(t, filepath.Join(sandbox.dir, "worker")) != "" }, sandbox.output)
				sandbox.send(t, "\x7fxy")
			default:
				sandbox.send(t, "xy")
			}
			result := sandbox.finish(t)
			assertDirectLifecycle(t, sandbox, test, result)
		})
	}
}

func runDirectPTYChild(t *testing.T) {
	t.Helper()
	dir, fd, before := preparePTYChild(t)
	config := filepath.Join(dir, "config.json")
	action := menuPaneClose
	confirm := true
	if sandboxCase() == fixtureAppSelected {
		action, confirm = menuSidebarToggle, false
	}
	spec := []Mapping{{Keys: "x", Action: action, Desc: "Lifecycle action", Confirm: confirm}}
	if sandboxCase() == fixtureWorkersDrainedCase {
		hosts := "Hosts"
		spec = append(spec, Mapping{Keys: "h", Provider: menuProviderHosts, Group: &hosts})
		require.NoError(t, unix.Mkfifo(filepath.Join(dir, "worker-fifo"), 0600))
	}
	sandboxJSON(t, config, map[string]any{"version": 1, "title": "Lifecycle action", "spec": spec})
	endpoint, client := sandboxEndpoint, sandboxClient
	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
	defer cancel()
	err := RunDirectMenu(ctx, DirectOptions{OriginSession: sandboxSession, OriginBlock: sandboxOrigin,
		OriginWindow: sandboxWindow, OriginLayer: sandboxLayer, Server: &endpoint, Client: &client,
		Config: &config, NoColor: true, ASCII: true})
	assertPTYRestored(t, fd, before)
	recordPTYResult(t, dir, MenuEvent{}, err)
}

func assertDirectLifecycle(t *testing.T, sandbox *ptySandbox, test directLifecycleCase, result ptyResult) {
	t.Helper()
	var calls []directFixtureCall
	for line := range strings.SplitSeq(strings.TrimSpace(sandboxRead(t, filepath.Join(sandbox.dir, "calls"))), "\n") {
		var call directFixtureCall
		require.NoError(t, json.Unmarshal([]byte(line), &call))
		calls = append(calls, call)
	}
	operations := make([]string, 0, len(calls))
	for _, call := range calls {
		operations = append(operations, call.Operation)
	}
	t.Logf("ordered executable CLI calls: %v", operations)
	require.NotContains(t, operations, "unknown", "%+v", calls)
	require.Equal(t, 1, countOperation(operations, fixtureDetach))
	require.Equal(t, test.focus, countOperation(operations, fixtureFocus), "no focus retry/fallback")
	require.Equal(t, test.effects, countOperation(operations, fixtureEffect), "exactly one effect, or none on failed cleanup/dismissal")
	if test.effects == 1 {
		detach := slices.Index(operations, fixtureDetach)
		placement := slices.Index(operations, "inventory detached")
		focus := slices.Index(operations, fixtureFocus)
		verified := slices.Index(operations, "inspect focused")
		effect := slices.Index(operations, fixtureEffect)
		require.Greater(t, placement, detach)
		require.Greater(t, focus, placement)
		require.Greater(t, verified, focus)
		require.Greater(t, effect, verified)
	}
	for _, call := range calls {
		if call.Operation == fixtureDetach || call.Operation == fixtureFocus || call.Operation == fixtureEffect {
			require.True(t, call.Restored, "TTY must already be restored before %s", call.Operation)
			require.True(t, call.WorkersDrained, "metadata subprocess must be reaped before %s", call.Operation)
		}
	}
	if test.problem == "" {
		require.Empty(t, result.Problem)
	} else {
		require.Contains(t, result.Problem, test.problem)
	}
}

func countOperation(operations []string, target string) int {
	count := 0
	for _, operation := range operations {
		if operation == target {
			count++
		}
	}
	return count
}

// This process implements only the current public argv/result shapes consumed
// by the real production methods. It never invokes Rex or accepts token matches.
func directCLIFixture() int {
	dir := os.Getenv(sandboxEnv)
	data, err := os.ReadFile(filepath.Join(dir, "state"))
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return 90
	}
	var state directFixtureState
	if len(data) > 0 && json.Unmarshal(data, &state) != nil {
		return 91
	}
	operation, result, code := directFixtureResponse(os.Args[1:], &state)
	call := directFixtureCall{Operation: operation, Argv: os.Args[1:], Restored: fixtureTTYRestored(dir), WorkersDrained: fixtureWorkersDrained(dir)}
	log, err := os.OpenFile(filepath.Join(dir, "calls"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return 92
	}
	err = json.NewEncoder(log).Encode(call)
	closeErr := log.Close()
	if err != nil || closeErr != nil {
		return 93
	}
	data, err = json.Marshal(state)
	if err != nil || os.WriteFile(filepath.Join(dir, "state"), data, 0600) != nil {
		return 94
	}
	if err := json.NewEncoder(os.Stdout).Encode(result); err != nil {
		return 95
	}
	return code
}

// A blocked inventory process makes worker joining externally observable.
// RunTerminal must cancel and Wait it before the first retirement mutation.
func fixturePendingWorker() (string, any, int) {
	dir := os.Getenv(sandboxEnv)
	data, err := json.Marshal(os.Getpid())
	if err != nil || os.WriteFile(filepath.Join(dir, "worker"), data, 0600) != nil {
		return fixtureWorkerFailure, nil, 96
	}
	file, err := os.Open(filepath.Join(dir, "worker-fifo"))
	if err != nil {
		return fixtureWorkerFailure, nil, 96
	}
	if err := file.Close(); err != nil {
		return fixtureWorkerFailure, nil, 96
	}
	return "worker unexpectedly released", map[string]any{}, 96
}

func fixtureWorkersDrained(dir string) bool {
	data, err := os.ReadFile(filepath.Join(dir, "worker"))
	if errors.Is(err, os.ErrNotExist) {
		return true
	}
	var pid int
	if err != nil || json.Unmarshal(data, &pid) != nil || pid <= 0 {
		return false
	}
	return errors.Is(syscall.Kill(pid, 0), syscall.ESRCH)
}

func fixtureTTYRestored(dir string) bool {
	data, err := os.ReadFile(filepath.Join(dir, "ready"))
	var ready ptyReady
	if err != nil || json.Unmarshal(data, &ready) != nil {
		return false
	}
	fd, err := unix.Open(ready.Name, unix.O_RDWR|unix.O_NOCTTY, 0)
	if err != nil {
		return false
	}
	after, readErr := unix.IoctlGetTermios(fd, unix.TIOCGETA)
	closeErr := unix.Close(fd)
	if readErr != nil || closeErr != nil {
		return false
	}
	ready.Original.Lflag &^= unix.PENDIN
	after.Lflag &^= unix.PENDIN
	return reflect.DeepEqual(ready.Original, *after)
}

func fixtureArgs(tail ...string) []string {
	return append([]string{"--server", sandboxEndpoint, "--autostart=false", "--timeout", "3s"}, tail...)
}

func fixtureAPI(method string, params map[string]string) []string {
	data, err := json.Marshal(params)
	if err != nil {
		return nil
	}
	return fixtureArgs("api", "call", "--session", sandboxSession, "--", method, string(data))
}

func directFixtureResponse(argv []string, state *directFixtureState) (string, any, int) {
	if slices.Equal(argv, []string{"--autostart=false", "--timeout", "3s", "hosts", "--json"}) && sandboxCase() == fixtureWorkersDrainedCase {
		return fixturePendingWorker()
	}
	if slices.Equal(argv, fixtureArgs("session", "inspect", "--json", "--", sandboxSession)) {
		operation := "inspect placed"
		if state.Detached {
			operation = "inspect retired"
		}
		if state.Focused {
			operation = "inspect focused"
		}
		return operation, fixtureSessionView(*state), 0
	}
	if slices.Equal(argv, fixtureAPI("session.list_clients", map[string]string{menuSessionID: sandboxSession})) {
		generation := "2026-11-07T00:00:00Z"
		if state.Detached && sandboxCase() == "generation changed" || state.Focused && sandboxCase() == "generation after focus" {
			generation = "2026-11-07T01:00:00Z"
		}
		return "clients", map[string]any{"clients": []any{map[string]any{"client_id": sandboxClient,
			"connected_at": generation, "connection_state": "connected", "info": map[string]string{fixtureKind: "app"}}}}, 0
	}
	if slices.Equal(argv, fixtureAPI("session.list_blocks", map[string]string{menuSessionID: sandboxSession})) {
		placement, window := "placed", sandboxWindow
		if state.Detached && sandboxCase() != "still placed" {
			placement, window = "detached", ""
		}
		return "inventory " + placement, map[string]any{fixtureBlocks: []any{
			map[string]string{fixtureBlockID: sandboxOrigin, "placement": "placed", menuWindowID: sandboxWindow,
				"creator_name": "com.superlogical.terminal", "flavor": "com.superlogical.terminal.shell"},
			map[string]string{fixtureBlockID: sandboxOwn, "placement": placement, menuWindowID: window,
				"creator_name": "com.superlogical.terminal", "flavor": "com.superlogical.terminal.shell"},
		}}, 0
	}
	if slices.Equal(argv, fixtureAPI("session.detach_block", map[string]string{menuSessionID: sandboxSession, fixtureBlockID: sandboxOwn})) {
		state.Detached = true
		return fixtureDetach, map[string]any{}, fixtureExit("detach uncertain")
	}
	if slices.Equal(argv, fixtureAPI("session.focus_block", map[string]string{menuSessionID: sandboxSession, fixtureBlockID: sandboxOrigin})) {
		state.Focused = true
		return fixtureFocus, map[string]any{}, fixtureExit("focus uncertain")
	}
	if slices.Equal(argv, fixtureArgs("block", "close", "--session", sandboxSession, "--", sandboxOrigin)) {
		return fixtureEffect, map[string]any{}, fixtureExit("effect uncertain")
	}
	if slices.Equal(argv, fixtureArgs("--client", sandboxClient, "actions", "--json")) {
		return "actions", map[string]any{"actions": []any{map[string]string{"name": menuSidebarToggle, fixtureKind: "client"}}}, 0
	}
	if slices.Equal(argv, fixturePerformArgs()) {
		return fixtureEffect, map[string]any{}, 0
	}
	return "unknown", map[string]string{"error": fmt.Sprintf("unexpected argv: %q", argv)}, 97
}

func fixturePerformArgs() []string {
	data, err := json.Marshal(map[string]any{"client_id": sandboxClient, "action": menuSidebarToggle, "args": map[string]any{},
		"origin": "cli", "timeout_ms": 3000, menuSessionID: sandboxSession, fixtureBlockID: sandboxOrigin})
	if err != nil {
		return nil
	}
	return fixtureArgs("api", "call", "--", "client.perform", string(data))
}

func fixtureExit(scenario string) int {
	if sandboxCase() == scenario {
		return 1
	}
	return 0
}

func fixtureSessionView(state directFixtureState) any {
	originLayer := sandboxLayer
	if state.Detached && sandboxCase() == "moved origin" || state.Focused && sandboxCase() == "moved after focus" {
		originLayer = sandboxID("layer", 9)
	}
	layers := []any{map[string]any{"layer_id": originLayer, fixtureKind: "tiling", fixtureBlocks: []any{map[string]string{fixtureBlockID: sandboxOrigin}}}}
	if !state.Detached || sandboxCase() == "visible after detach" {
		layers = append(layers, map[string]any{"layer_id": sandboxFloat, fixtureKind: "floating", fixtureBlocks: []any{map[string]string{fixtureBlockID: sandboxOwn}}})
	}
	focus := sandboxOwn
	if state.Detached {
		focus = "" // No focused pane is valid read metadata, never focus authority.
	}
	if state.Focused && sandboxCase() != "focus unverified" {
		focus = sandboxOrigin
	}
	return map[string]any{menuSessionID: sandboxSession, "label": "Original session", "active_window_id": sandboxWindow,
		"windows": []any{map[string]any{menuWindowID: sandboxWindow, "label": "Original tab", "focused_block_id": focus, "layers": layers}}}
}
