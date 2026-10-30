package probe_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func capturePath(t *testing.T, name string) string {
	t.Helper()
	return filepath.Join(t.TempDir(), "unicode-"+name+".vt")
}

func (tr *terminal) frame(t *testing.T) {
	t.Helper()
	tr.until(t, footer)
}

func TestUnicodeObservation(t *testing.T) {
	for _, name := range []string{"direct", "tea"} {
		t.Run(name, func(t *testing.T) {
			tr := startTerminal(t, name)
			tr.frame(t)
			// Keep the actual terminal bytes as research evidence. This is not a
			// screenshot or a proof of on-screen glyph widths.
			require.NoError(t, os.WriteFile(capturePath(t, name), tr.output.Bytes(), 0600))
			preserved := strings.Contains(tr.output.String(), "界 e\u0301")
			t.Logf("%s combining_mark_preserved=%t capture=%q", name, preserved, tr.output.String())
			if os.Getenv("WK_STRICT_UNICODE") == "1" {
				require.True(t, preserved, "combining-mark acceptance gate is unmet")
			}
			tr.send(t, "y")
			tr.finish(t, 0)
		})
	}
}
