package probe_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Direct remains the unchanged raw-byte diagnostic, not a qualified candidate.
func TestDiagnosticDirectColorReply(t *testing.T) {
	tr := startTerminal(t, "direct")
	tr.frame(t)
	tr.send(t, "\x1b]11;rgb:2222/3333/4444\x07y")
	tr.until(t, "RESEARCH RESULT cancel")
	tr.finish(t, 0)
}

func (tr *terminal) quiet(t *testing.T, duration time.Duration) {
	t.Helper()
	timer := time.NewTimer(duration)
	defer timer.Stop()
	for {
		select {
		case chunk, ok := <-tr.chunks:
			require.True(t, ok, "terminal closed during incomplete reply")
			tr.output.Write(chunk)
			require.NotContains(t, tr.output.String(), "RESEARCH RESULT")
			require.NotContains(t, tr.output.String(), "TERMINAL_RESTORED")
		case <-timer.C:
			return
		}
	}
}

func TestFragmentedOSC(t *testing.T) {
	// Initial ESC shares the bounded standalone-Escape ambiguity interval.
	// This table exercises every split inside that interval. Established OSC
	// retention beyond it is covered separately below.
	for _, terminator := range []string{"\a", "\x1b\\"} {
		reply := "\x1b]11;rgb:2222/3333/4444" + terminator
		for split := 0; split <= len(reply); split++ {
			t.Run(fmt.Sprintf("%x/%d", terminator, split), func(t *testing.T) {
				tr := startTerminal(t, "tea")
				tr.frame(t)
				tr.send(t, reply[:split])
				tr.quiet(t, 10*time.Millisecond)
				tr.send(t, reply[split:])
				tr.quiet(t, 10*time.Millisecond)
				require.False(t, strings.Contains(tr.output.String(), "RESEARCH RESULT"))
				tr.send(t, "y")
				tr.until(t, "RESEARCH RESULT choose")
				tr.finish(t, 0)
			})
		}
	}
}
