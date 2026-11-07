//go:build darwin || linux

package whichkey

import (
	"context"
	"io"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

func TestTerminalDisplayRollback(t *testing.T) {
	for _, test := range []struct {
		name    string
		display terminalDisplay
		want    string
	}{
		{"unarmed", terminalDisplay{}, "\x18"},
		{"cursor only", terminalDisplay{cursorSaved: true}, "\x18\x1b8"},
		{"full", terminalDisplay{cursorSaved: true, modesSaved: true, viewportUsed: true}, "\x18\x1b[39;49m\x1b[2J\x1b8\x1b[?7;25;1004;2004r"},
	} {
		t.Run(test.name, func(t *testing.T) {
			reader, writer, err := os.Pipe()
			require.NoError(t, err)
			defer func() { require.NoError(t, reader.Close()) }()
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			require.NoError(t, test.display.restore(ctx, int(writer.Fd())), "cleanup must outlive input cancellation")
			require.NoError(t, writer.Close())
			data, err := io.ReadAll(reader)
			require.NoError(t, err)
			require.Equal(t, test.want, string(data))
		})
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var display terminalDisplay
	require.ErrorIs(t, display.setup(ctx, -1), context.Canceled)
	require.Equal(t, terminalDisplay{}, display, "failed saves must not arm rollback")
	require.Error(t, display.restore(ctx, -1))
}

func TestTerminalDisplayCleanupIsBounded(t *testing.T) {
	var fds [2]int
	require.NoError(t, unix.Pipe(fds[:]))
	defer func() { require.NoError(t, unix.Close(fds[0])) }()
	defer func() { require.NoError(t, unix.Close(fds[1])) }()
	require.NoError(t, unix.SetNonblock(fds[1], true))
	data := make([]byte, 4096)
	for {
		_, err := unix.Write(fds[1], data)
		if err != nil {
			require.True(t, terminalRetry(err))
			break
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	started := time.Now()
	err := (terminalDisplay{cursorSaved: true, modesSaved: true}).restore(ctx, fds[1])
	require.ErrorIs(t, err, context.DeadlineExceeded)
	require.Less(t, time.Since(started), 2*time.Second)
}

func TestTerminalReadiness(t *testing.T) {
	reader, writer, err := os.Pipe()
	require.NoError(t, err)
	defer func() { require.NoError(t, reader.Close()) }()
	defer func() { require.NoError(t, writer.Close()) }()
	fds := []unix.PollFd{{Fd: int32(reader.Fd()), Events: unix.POLLIN, Revents: unix.POLLNVAL}}
	count, err := pollTerminal(fds, 0)
	require.NoError(t, err)
	require.Zero(t, count)
	require.Zero(t, fds[0].Revents, "timeout must clear stale readiness flags")
	fds[0] = unix.PollFd{Fd: int32(writer.Fd()), Events: unix.POLLOUT}
	count, err = pollTerminal(fds, terminalPollInterval)
	require.NoError(t, err)
	require.Equal(t, 1, count)
	require.NotZero(t, fds[0].Revents&unix.POLLOUT)
	_, err = writer.Write([]byte("x"))
	require.NoError(t, err)
	fds[0] = unix.PollFd{Fd: int32(reader.Fd()), Events: unix.POLLIN}
	count, err = pollTerminal(fds, terminalPollInterval)
	require.NoError(t, err)
	require.Equal(t, 1, count)
	require.NotZero(t, fds[0].Revents&unix.POLLIN)
}
