//go:build darwin

package whichkey

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	terminalConfirmResize = "confirmation resize"
	terminalEditCase      = "editing"
	terminalEscapeCase    = "escape"
	terminalPaletteCase   = "palette"
	terminalPageCase      = "pagination"
	terminalUnicodeCase   = "fragmented UTF8"
)

func terminalInteractionConfig() Config {
	config := DefaultConfig()
	config.Title = "PTY interaction"
	config.Spec = []Mapping{
		{Keys: "x", Action: "test_action", Desc: "Select action"},
		{Keys: "é", Action: "unicode_action", Desc: "Unicode action"},
		{Keys: "r", Action: menuTabRename, Desc: "Inline edit"},
		{Keys: "c", Action: "confirmed_action", Desc: "Critical action", Confirm: true},
	}
	for _, key := range "abcdefghijkl" {
		config.Spec = append(config.Spec, Mapping{Keys: "g" + string(key), Action: "group_action", Desc: fmt.Sprintf("Entry %c with long label", key)})
	}
	config.textEntry = &textEntrySpec{Initial: "original"}
	return config
}

func TestTerminalInteraction(t *testing.T) {
	if os.Getenv(sandboxEnv) != "" {
		dir, fd, before := preparePTYChild(t)
		ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
		defer cancel()
		event, err := RunTerminal(ctx, terminalInteractionConfig(), true, true)
		assertPTYRestored(t, fd, before)
		recordPTYResult(t, dir, event, err)
		return
	}
	for _, scenario := range []string{"selection", terminalUnicodeCase, terminalEditCase, terminalEscapeCase, "reports", terminalPaletteCase, terminalPageCase, terminalConfirmResize, CompletionCancel} {
		t.Run(scenario, func(t *testing.T) {
			sandbox := startPTYSandbox(t, "TestTerminalInteraction", scenario)
			sandbox.expect(t, 0, "Select action")
			driveTerminalInteraction(t, sandbox, scenario)
			result := sandbox.finish(t)
			assertTerminalInteraction(t, scenario, result)
		})
	}
}

// A bounded observation interval intentionally spans more than the 50 ms
// ambiguity window. Established sequences must not turn into Escape or keys.
func (sandbox *ptySandbox) holdSequence(t *testing.T) {
	t.Helper()
	until := time.NewTimer(120 * time.Millisecond)
	defer until.Stop()
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()
	for {
		require.Empty(t, sandboxRead(t, filepath.Join(sandbox.dir, "result")), "%s", sandbox.output())
		require.NotContains(t, sandbox.output(), "FAIL")
		select {
		case <-until.C:
			return
		case <-ticker.C:
		}
	}
}

func driveTerminalInteraction(t *testing.T, sandbox *ptySandbox, scenario string) {
	t.Helper()
	switch scenario {
	case "selection":
		sandbox.send(t, "xx") // Queued duplicate input still emits one selection.
	case terminalUnicodeCase:
		sandbox.send(t, "\xc3")
		sandbox.holdSequence(t)
		sandbox.send(t, "\xa9")
	case terminalEditCase:
		driveTerminalEditing(t, sandbox)
	case terminalEscapeCase:
		start := time.Now()
		sandbox.send(t, "\x1b")
		sandboxAwait(t, func() bool { return sandboxRead(t, filepath.Join(sandbox.dir, "result")) != "" }, sandbox.output)
		require.GreaterOrEqual(t, time.Since(start), terminalEscapeDelay)
	case "reports":
		sandbox.send(t, "\x1b[")
		sandbox.holdSequence(t)
		sandbox.send(t, "1;2R\x1b]0;xy")
		sandbox.holdSequence(t)
		sandbox.send(t, "\x07\x1bOP\x1b[?1;2c\x1b[O\x1b[I")
		sandbox.expect(t, 0, "PTY interaction")
		sandbox.holdSequence(t)
		sandbox.send(t, "x")
	case terminalPaletteCase:
		// c must establish confirmation before the reply; palette bytes must
		// not cancel it. The key after the fragmented reply must still be routed.
		mark := sandbox.send(t, "c\x1b]11;rgb:0000/")
		sandbox.expect(t, mark, "Critical action")
		sandbox.holdSequence(t)
		mark = sandbox.send(t, "0000/0000\x1b\\\x1b[I")
		sandbox.expect(t, mark, "Critical action")
		sandbox.send(t, "y")
	case terminalPageCase:
		driveTerminalPagination(t, sandbox)
	case terminalConfirmResize:
		mark := sandbox.send(t, "c")
		sandbox.expect(t, mark, "Critical action")
		mark = sandbox.resize(t, 48, 7)
		sandbox.expect(t, mark, "Critical action")
		sandbox.send(t, "\x1b[Iy")
	case CompletionCancel:
		mark := sandbox.send(t, "r\x15edited")
		sandbox.expect(t, mark, "edited|")
		require.NoError(t, syscall.Kill(sandbox.ready.PID, syscall.SIGTERM))
	}
}

func driveTerminalEditing(t *testing.T, sandbox *ptySandbox) {
	t.Helper()
	mark := sandbox.send(t, "r\x15A\xc3")
	sandbox.expect(t, mark, "A|")
	sandbox.holdSequence(t)
	mark = sandbox.send(t, "\xa9B\x1b[D\x7f")
	sandbox.expect(t, mark, "A|B")
	mark = sandbox.resize(t, 48, 7)
	sandbox.expect(t, mark, "A|B")
	sandbox.send(t, "\x1b[200~猫")
	sandbox.holdSequence(t)
	mark = sandbox.send(t, "Z\x1b[201~")
	sandbox.expect(t, mark, "A猫Z|B")
	mark = sandbox.send(t, "\x1b[3~\x1b[H\x1b[C!")
	sandbox.expect(t, mark, "A!|猫Z")
	sandbox.send(t, "\r")
}

func driveTerminalPagination(t *testing.T, sandbox *ptySandbox) {
	t.Helper()
	// The two-page root becomes a four-page group before the queued Tabs.
	// Reusing the old count for this read would incorrectly wrap back to 1/4.
	mark := sandbox.resize(t, 40, 5)
	sandbox.expect(t, mark, "Select action")
	mark = sandbox.send(t, "g\t\t")
	sandbox.expect(t, mark, "3/4")
	mark = sandbox.resize(t, 40, 6)
	sandbox.expect(t, mark, "3/3")
	mark = sandbox.send(t, "\t\t")
	sandbox.expect(t, mark, "2/3")
	sandbox.send(t, "l") // Off-page keys remain active, exact group retained.
}

func assertTerminalInteraction(t *testing.T, scenario string, result ptyResult) {
	t.Helper()
	if scenario == CompletionCancel {
		require.Empty(t, result.Keys)
		require.Nil(t, result.Text)
		require.False(t, result.Dismissed)
		require.Contains(t, result.Problem, "context canceled")
		return
	}
	require.Empty(t, result.Problem)
	require.Equal(t, scenario == terminalEscapeCase, result.Dismissed)
	switch scenario {
	case terminalEscapeCase:
		require.Empty(t, result.Keys)
	case terminalUnicodeCase:
		require.Equal(t, "é", result.Keys)
	case terminalEditCase:
		require.Equal(t, "r", result.Keys)
		require.NotNil(t, result.Text)
		require.Equal(t, "A!猫Z", *result.Text)
	case terminalPaletteCase, terminalConfirmResize:
		require.Equal(t, "c", result.Keys)
		require.True(t, result.Confirmed)
	case terminalPageCase:
		require.Equal(t, "gl", result.Keys)
	default:
		require.Equal(t, "x", result.Keys)
	}
	require.False(t, strings.Contains(result.Problem, "deadline"))
}
