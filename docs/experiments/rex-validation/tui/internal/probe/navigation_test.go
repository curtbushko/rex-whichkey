package probe_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNavigation(t *testing.T) {
	for _, name := range []string{"direct", "tea"} {
		t.Run(name, func(t *testing.T) {
			tr := startTerminal(t, name)
			tr.frame(t)
			for i, step := range []struct{ key, page string }{{"\t", "1"}, {"\t", "2"}, {"\x7f", "1"}, {"\x7f", "0"}, {"\x7f", "0"}} {
				tr.output.Reset()
				tr.send(t, step.key)
				// Backspace at zero can legitimately produce no new frame in Tea.
				if i == 4 {
					tr.quiet(t, 10*time.Millisecond)
					continue
				}
				marker := "PAGE=" + step.page
				if name == "tea" {
					// The pinned renderer diffs column 21. After the first update,
					// its cursor is already two columns past that digit.
					marker = "\x1b[1;21H" + step.page
					if i > 0 {
						marker = "\b\b" + step.page + " "
					}
				}
				tr.until(t, marker)
				require.NotContains(t, tr.output.String(), "RESEARCH RESULT")
			}
			tr.send(t, "y")
			tr.until(t, "RESEARCH RESULT choose")
			tr.finish(t, 0)
		})
	}
}
