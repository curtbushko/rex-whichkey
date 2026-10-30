package probe_test

import (
	"bytes"
	"crypto/sha256"
	"debug/buildinfo"
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
	file       *os.File
	cmd        *exec.Cmd
	chunks     chan []byte
	output     bytes.Buffer
	started    time.Time
	done       chan struct{}
	readerDone chan struct{}
	waited     bool
}

func startTerminal(t *testing.T, name string) *terminal {
	t.Helper()
	root, err := filepath.Abs("../..")
	require.NoError(t, err)
	binaryRoot := os.Getenv("WK_FIXTURE_BINARY_DIR")
	if binaryRoot == "" {
		binaryRoot = root
	}
	binary := filepath.Join(binaryRoot, name)
	data, err := os.ReadFile(binary)
	require.NoError(t, err)
	info, err := buildinfo.ReadFile(binary)
	require.NoError(t, err)
	t.Logf("binary=%s sha256=%x build=%s", binary, sha256.Sum256(data), info)
	cmd := exec.Command("/bin/bash", filepath.Join(root, "restore.sh"))
	for _, entry := range os.Environ() {
		// PTY tests never inherit a live Rex callback destination.
		if !strings.HasPrefix(entry, "WK_") || strings.HasPrefix(entry, "WK_ASCII=") {
			cmd.Env = append(cmd.Env, entry)
		}
	}
	cmd.Env = append(cmd.Env, "TERM=xterm-256color", "WK_BINARY="+binary)
	started := time.Now()
	f, err := pty.StartWithSize(cmd, &pty.Winsize{Rows: 24, Cols: 80})
	require.NoError(t, err)
	tr := &terminal{file: f, cmd: cmd, chunks: make(chan []byte, 128), started: started, done: make(chan struct{}), readerDone: make(chan struct{})}
	go func() {
		defer close(tr.chunks)
		defer close(tr.readerDone)
		for {
			b := make([]byte, 4096)
			n, err := f.Read(b)
			if n > 0 {
				select {
				case tr.chunks <- b[:n]:
				case <-tr.done:
					return
				}
			}
			if err != nil {
				return
			}
		}
	}()
	t.Cleanup(func() {
		close(tr.done)
		if !tr.waited {
			if err := cmd.Process.Kill(); err != nil && !errors.Is(err, os.ErrProcessDone) {
				t.Logf("cleanup kill: %v", err)
			}
			if err := cmd.Wait(); err != nil {
				t.Logf("cleanup wait after failed test: %v", err)
			}
		}
		require.NoError(t, f.Close())
		<-tr.readerDone
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
	tr.waited = true
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
				tr.frame(t)
				start := time.Now()
				tr.send(t, tc.key)
				tr.finish(t, tc.code)
				if tc.name == "escape" {
					require.Less(t, time.Since(start), time.Second, "standalone Escape must resolve within a bounded interval")
				}
				if tc.code == 0 {
					want := "cancel"
					if tc.name == "select" || tc.name == "navigate" {
						want = "choose"
					}
					require.Contains(t, tr.output.String(), "RESEARCH RESULT "+want)
				}
			})
		}
		t.Run(name+"/resize", func(t *testing.T) {
			tr := startTerminal(t, name)
			tr.frame(t)
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
				tr.frame(t)
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
