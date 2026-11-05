//go:build darwin || linux

package whichkey

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

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
