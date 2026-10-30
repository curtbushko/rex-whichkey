package probe_test

import (
	"context"
	"fmt"
	"io"
	"testing"
	"time"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/stretchr/testify/require"
)

func TestReaderOSC(t *testing.T) {
	for _, terminator := range []string{"\a", "\x1b\\"} {
		reply := "\x1b]11;rgb:eeee/cccc/aaaa" + terminator
		for split := 1; split < len(reply); split++ {
			t.Run(fmt.Sprintf("%x/%d", terminator, split), func(t *testing.T) {
				testReaderReply(t, reply, split, 10*time.Millisecond)
			})
		}
	}
}

func TestReaderDelayedST(t *testing.T) {
	reply := "\x1b]11;rgb:eeee/cccc/aaaa\x1b\\"
	testReaderReply(t, reply, len(reply)-1, 2*uv.DefaultEscTimeout)
}

func TestReaderEstablishedOSC(t *testing.T) {
	for _, terminator := range []string{"\a", "\x1b\\", "\x9c"} {
		reply := "\x1b]11;rgb:eeee/cccc/aaaa" + terminator
		// Once OSC is established, Escape's ambiguity timeout no longer applies.
		for split := 3; split < len(reply); split++ {
			t.Run(fmt.Sprintf("%x/%d", terminator, split), func(t *testing.T) {
				testReaderReply(t, reply, split, 2*uv.DefaultEscTimeout)
			})
		}
	}
}

func TestReaderStandaloneEscape(t *testing.T) {
	r, w := io.Pipe()
	ctx, cancel := context.WithCancel(t.Context())
	reader := uv.NewTerminalReader(r, "xterm-256color")
	events := make(chan uv.Event, 16)
	done := make(chan error, 1)
	go func() { done <- reader.StreamEvents(ctx, events) }()
	t.Cleanup(func() {
		cancel()
		require.NoError(t, w.Close())
		require.NoError(t, r.Close())
		select {
		case err := <-done:
			require.NoError(t, err)
		case <-time.After(time.Second):
			t.Error("terminal reader failed to stop")
		}
	})
	start := time.Now()
	_, err := io.WriteString(w, "\x1b")
	require.NoError(t, err)
	select {
	case event := <-events:
		key, ok := event.(uv.KeyPressEvent)
		require.True(t, ok, "Escape emitted %T", event)
		require.Equal(t, uv.KeyEscape, key.Code)
		require.Less(t, time.Since(start), time.Second)
	case <-time.After(time.Second):
		t.Fatal("standalone Escape did not resolve")
	}
}

func testReaderReply(t *testing.T, reply string, split int, gap time.Duration) {
	t.Helper()
	r, w := io.Pipe()
	ctx, cancel := context.WithCancel(t.Context())
	reader := uv.NewTerminalReader(r, "xterm-256color")
	events := make(chan uv.Event, 64)
	done := make(chan error, 1)
	go func() { done <- reader.StreamEvents(ctx, events) }()
	t.Cleanup(func() {
		cancel()
		require.NoError(t, w.Close())
		require.NoError(t, r.Close())
		select {
		case err := <-done:
			require.NoError(t, err)
		case <-time.After(time.Second):
			t.Error("terminal reader failed to stop")
		}
	})
	_, err := io.WriteString(w, reply[:split])
	require.NoError(t, err)
	select {
	case event := <-events:
		t.Fatalf("incomplete OSC emitted %T: %v", event, event)
	case <-time.After(gap):
	}
	_, err = io.WriteString(w, reply[split:])
	require.NoError(t, err)
	select {
	case event := <-events:
		require.IsType(t, uv.BackgroundColorEvent{}, event, "complete response emitted %T: %v", event, event)
	case <-time.After(time.Second):
		t.Fatal("no complete background response")
	}
	_, err = io.WriteString(w, "y")
	require.NoError(t, err)
	select {
	case event := <-events:
		key, ok := event.(uv.KeyPressEvent)
		require.True(t, ok, "following y must be the first key, got %T", event)
		require.Equal(t, "y", key.String())
	case <-time.After(time.Second):
		t.Fatal("following y was lost")
	}
	select {
	case event := <-events:
		t.Fatalf("unexpected extra event %T: %v", event, event)
	default:
	}
}
