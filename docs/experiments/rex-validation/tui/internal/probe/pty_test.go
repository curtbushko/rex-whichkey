package probe_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/creack/pty"
	"github.com/stretchr/testify/require"
)

type terminal struct {
	file    *os.File
	cmd     *exec.Cmd
	chunks  chan []byte
	output  bytes.Buffer
	started time.Time
}

func startTerminal(t *testing.T, name string) *terminal {
	t.Helper()
	root, err := filepath.Abs("../..")
	require.NoError(t, err)
	cmd := exec.Command("/bin/bash", filepath.Join(root, "restore.sh"))
	cmd.Env = append(os.Environ(), "TERM=xterm-256color", "WK_BINARY="+filepath.Join(root, name))
	started := time.Now()
	f, err := pty.StartWithSize(cmd, &pty.Winsize{Rows: 24, Cols: 80})
	require.NoError(t, err)
	tr := &terminal{file: f, cmd: cmd, chunks: make(chan []byte, 128), started: started}
	go func() {
		defer close(tr.chunks)
		for {
			b := make([]byte, 4096)
			n, err := f.Read(b)
			if n > 0 {
				tr.chunks <- b[:n]
			}
			if err != nil {
				return
			}
		}
	}()
	t.Cleanup(func() {
		_ = f.Close()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
	})
	return tr
}
func (tr *terminal) until(t *testing.T, marker string) {
	t.Helper()
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	for !strings.Contains(tr.output.String(), marker) {
		select {
		case b, ok := <-tr.chunks:
			require.True(t, ok, "terminal closed before %q: %q", marker, tr.output.String())
			tr.output.Write(b)
		case <-timer.C:
			t.Fatalf("terminal timed out before %q: %q", marker, tr.output.String())
		}
	}
}
func (tr *terminal) send(t *testing.T, text string) {
	t.Helper()
	_, err := io.WriteString(tr.file, text)
	require.NoError(t, err)
}
func (tr *terminal) finish(t *testing.T, want int) {
	t.Helper()
	tr.until(t, "TERMINAL_RESTORED")
	err := tr.cmd.Wait()
	if want == 0 {
		require.NoError(t, err)
	} else {
		var exit *exec.ExitError
		require.True(t, errors.As(err, &exit))
		require.Equal(t, want, exit.ExitCode())
	}
}

func TestPTY(t *testing.T) {
	for _, name := range []string{"direct", "tea"} {
		for _, tc := range []struct {
			name, key string
			code      int
		}{{"select", "y", 0}, {"cancel", "q", 0}, {"escape", "\x1b", 0}, {"error", "e", 1}, {"navigate", "\t\x7fy", 0}} {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				tr := startTerminal(t, name)
				tr.until(t, "RESEARCH READY")
				tr.send(t, tc.key)
				tr.finish(t, tc.code)
			})
		}
		t.Run(name+"/resize", func(t *testing.T) {
			tr := startTerminal(t, name)
			tr.until(t, "RESEARCH READY")
			tr.output.Reset()
			require.NoError(t, pty.Setsize(tr.file, &pty.Winsize{Rows: 12, Cols: 40}))
			tr.until(t, "40x12")
			tr.send(t, "y")
			tr.finish(t, 0)
		})
		t.Run(name+"/ascii", func(t *testing.T) {
			t.Setenv("WK_ASCII", "1")
			t.Setenv("NO_COLOR", "1")
			tr := startTerminal(t, name)
			tr.until(t, "ASCII")
			require.NotContains(t, tr.output.String(), "界")
			tr.send(t, "y")
			tr.finish(t, 0)
		})
	}
}

func TestStartupSamples(t *testing.T) {
	for _, name := range []string{"direct", "tea"} {
		t.Run(name, func(t *testing.T) {
			first := make([]float64, 0, 20)
			exit := make([]float64, 0, 20)
			for range 20 {
				tr := startTerminal(t, name)
				tr.until(t, "RESEARCH READY")
				first = append(first, float64(time.Since(tr.started).Microseconds())/1000)
				start := time.Now()
				tr.send(t, "y")
				tr.finish(t, 0)
				exit = append(exit, float64(time.Since(start).Microseconds())/1000)
			}
			sort.Float64s(first)
			sort.Float64s(exit)
			t.Logf("%s first-content ms min=%.3f median=%.3f p95=%.3f; key-to-restored-exit ms min=%.3f median=%.3f p95=%.3f; n=20", name, first[0], first[10], first[18], exit[0], exit[10], exit[18])
			require.True(t, len(first) == 20, fmt.Sprintf("samples: %v", first))
		})
	}
}
