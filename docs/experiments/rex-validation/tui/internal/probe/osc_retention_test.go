package probe_test

import (
	"context"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/stretchr/testify/require"
)

func oscReader(t *testing.T) (io.Writer, <-chan uv.Event) {
	t.Helper()
	r, w := io.Pipe()
	ctx, cancel := context.WithCancel(t.Context())
	events := make(chan uv.Event, 64)
	done := make(chan error, 1)
	go func() { done <- uv.NewTerminalReader(r, "xterm-256color").StreamEvents(ctx, events) }()
	t.Cleanup(func() {
		cancel()
		require.NoError(t, w.Close())
		require.NoError(t, r.Close())
		select {
		case err := <-done:
			require.NoError(t, err)
		case <-time.After(time.Second):
			t.Error("reader shutdown stalled")
		}
	})
	return w, events
}

func noOSCEvent(t *testing.T, events <-chan uv.Event, gap time.Duration) {
	t.Helper()
	select {
	case ev := <-events:
		t.Fatalf("pending OSC leaked %T: %v", ev, ev)
	case <-time.After(gap):
	}
}

func followingY(t *testing.T, w io.Writer, events <-chan uv.Event) {
	t.Helper()
	_, err := io.WriteString(w, "y")
	require.NoError(t, err)
	select {
	case ev := <-events:
		require.Equal(t, uv.KeyPressEvent{Code: 'y', Text: "y"}, ev)
	case <-time.After(time.Second):
		t.Fatal("lost y")
	}
	noOSCEvent(t, events, 60*time.Millisecond)
}

func TestReaderRepeatedOSCTimeouts(t *testing.T) {
	for _, cancel := range []string{"\x18", "\x1a", "\x1b\x18", "\x1b\x1a"} {
		t.Run(fmt.Sprintf("cancel/%x", cancel), func(t *testing.T) {
			w, events := oscReader(t)
			// Preserve the original upstream OSC and payload chunks, but prohibit its
			// premature UnknownEvent and abcxxxx key delivery across repeated timeouts.
			for _, chunk := range []string{"\x1b]11;rgb:1111/2222/3333", "abc", "x", "x", "x", "x"} {
				_, err := io.WriteString(w, chunk)
				require.NoError(t, err)
				noOSCEvent(t, events, 60*time.Millisecond)
			}
			_, err := io.WriteString(w, cancel)
			require.NoError(t, err)
			followingY(t, w, events)
		})
	}
	for _, term := range []string{"\a", "\x1b\\"} {
		t.Run(fmt.Sprintf("valid/%x", term), func(t *testing.T) {
			w, events := oscReader(t)
			for _, chunk := range []string{"\x1b]11;", "rgb:1111/", "2222/3333"} {
				_, err := io.WriteString(w, chunk)
				require.NoError(t, err)
				noOSCEvent(t, events, 100*time.Millisecond)
			}
			for _, b := range []byte(term) {
				_, err := w.Write([]byte{b})
				require.NoError(t, err)
				if b == 27 {
					noOSCEvent(t, events, 100*time.Millisecond)
				}
			}
			select {
			case ev := <-events:
				require.Equal(t, uv.BackgroundColorEvent{Color: ansi.XParseColor("rgb:1111/2222/3333")}, ev)
			case <-time.After(time.Second):
				t.Fatal("missing color")
			}
			followingY(t, w, events)
		})
	}
}

func TestReaderOSCOverflowResync(t *testing.T) {
	for _, term := range []string{"\a", "\x9c", "\x1b\\", "\x18", "\x1a"} {
		t.Run(fmt.Sprintf("%x", term), func(t *testing.T) {
			w, events := oscReader(t)
			_, err := io.WriteString(w, "\x1b]999;")
			require.NoError(t, err)
			chunk := strings.Repeat("a", 4096)
			for range 256 {
				_, err = io.WriteString(w, chunk)
				require.NoError(t, err)
			}
			// Even non-ST ESC and key-like payload in discard mode cannot become keys.
			for _, payload := range []string{"\x1b[A", "\x1b[0;0;121;1;0;0_", "\x1b[2Jpayload"} {
				_, err = io.WriteString(w, payload)
				require.NoError(t, err)
				noOSCEvent(t, events, 100*time.Millisecond)
			}
			for _, b := range []byte(term) {
				_, err = w.Write([]byte{b})
				require.NoError(t, err)
				if b == 27 {
					noOSCEvent(t, events, 100*time.Millisecond)
				}
			}
			noOSCEvent(t, events, 60*time.Millisecond)
			followingY(t, w, events)
			// Synchronization must also allow the next legitimate reply.
			_, err = io.WriteString(w, "\x1b]11;rgb:1111/2222/3333\a")
			require.NoError(t, err)
			select {
			case ev := <-events:
				require.Equal(t, uv.BackgroundColorEvent{Color: ansi.XParseColor("rgb:1111/2222/3333")}, ev)
			case <-time.After(time.Second):
				t.Fatal("lost reply after resync")
			}
		})
	}
}

func TestReaderOSCShutdown(t *testing.T) {
	for _, size := range []int{0, 1 << 20} {
		t.Run(fmt.Sprint(size), func(t *testing.T) {
			w, events := oscReader(t)
			_, err := io.WriteString(w, "\x1b]999;"+strings.Repeat("a", size)+"\x1b")
			require.NoError(t, err)
			noOSCEvent(t, events, 100*time.Millisecond)
			// Cleanup cancels/closes an incomplete or discarded OSC without Escape.
		})
	}
}
