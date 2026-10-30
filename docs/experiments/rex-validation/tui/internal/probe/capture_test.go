package probe_test

import (
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const footer = "Tab page; Backspace back; y select; q/Esc cancel; e fail"

func TestCapturePaths(t *testing.T) {
	first, second := capturePath(t, "tea"), capturePath(t, "tea")
	require.NotEqual(t, first, second, "captures must never overwrite one another")
	require.DirExists(t, filepath.Dir(first))
}

func TestCaptureCompleteFrame(t *testing.T) {
	frame := "RESEARCH READY PAGE=0 SIZE=80x24 界 e\u0301\r\n" + footer
	for split := 1; split < len(frame); split++ {
		t.Run(strconv.Itoa(split), func(t *testing.T) {
			tr := &terminal{chunks: make(chan []byte, 2)}
			tr.chunks <- []byte(frame[:split])
			done := make(chan struct{})
			go func() {
				defer close(done)
				time.Sleep(time.Millisecond)
				tr.chunks <- []byte(frame[split:])
			}()
			tr.frame(t)
			<-done
			require.Equal(t, frame, tr.output.String(), "split at byte %d", split)
		})
	}
}
