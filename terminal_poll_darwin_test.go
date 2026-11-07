//go:build darwin

package whichkey

import (
	"context"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

func TestDarwinTerminalReadinessRejectsInvalidDescriptors(t *testing.T) {
	for _, fd := range []int32{-1, unix.FD_SETSIZE} {
		fds := []unix.PollFd{{Fd: fd, Events: unix.POLLIN}}
		_, err := pollTerminal(fds, 0)
		require.ErrorIs(t, err, unix.EBADF)
	}
}

// script supplies a real controlling PTY. A pipe or a /dev/null stand-in would
// not reproduce Darwin's /dev/tty readiness behavior.
func TestDarwinControllingTerminalOutput(t *testing.T) {
	if os.Getenv("REX_WHICHKEY_TTY_TEST") == "1" {
		fd, err := unix.Open("/dev/tty", unix.O_RDWR|unix.O_NOCTTY|unix.O_NONBLOCK|unix.O_CLOEXEC, 0)
		require.NoError(t, err)
		defer func() { require.NoError(t, unix.Close(fd)) }()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		require.NoError(t, writeTerminal(ctx, fd, "terminal readiness regression\r\n"))
		config := DefaultConfig()
		config.Spec = []Mapping{{Keys: "x", Action: "test_action", Desc: "Test action"}}
		idle, stop := context.WithTimeout(ctx, 100*time.Millisecond)
		defer stop()
		// With no input, the real menu must stay alive until cancellation,
		// not fail its first read-readiness check or dismiss itself.
		before, err := unix.IoctlGetTermios(fd, unix.TIOCGETA)
		require.NoError(t, err)
		_, err = RunTerminal(idle, config, true, true)
		require.ErrorIs(t, err, context.DeadlineExceeded)
		after, err := unix.IoctlGetTermios(fd, unix.TIOCGETA)
		require.NoError(t, err)
		// Darwin sets kernel-transient PENDIN on restoration (observed here).
		// Compare every configured flag, control byte and speed, excluding only
		// that transient bit. Production restores the captured state unchanged.
		before.Lflag &^= unix.PENDIN
		after.Lflag &^= unix.PENDIN
		require.Equal(t, before, after, "configured termios must be restored after cancellation")
		return
	}
	executable, err := os.Executable()
	require.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	command := exec.CommandContext(ctx, "/usr/bin/script", "-q", "/dev/null", executable,
		"-test.run", "^TestDarwinControllingTerminalOutput$")
	command.Env = append(os.Environ(), "REX_WHICHKEY_TTY_TEST=1")
	// Keep script's stdin open until the child exits, so its EOF does not
	// inject an EOT into the PTY and impersonate terminal input.
	input, err := command.StdinPipe()
	require.NoError(t, err)
	defer func() {
		if closeErr := input.Close(); closeErr != nil {
			require.ErrorIs(t, closeErr, os.ErrClosed) // Wait normally closes the pipe.
		}
	}()
	output, err := command.CombinedOutput()
	require.NoError(t, err, "%s", output)
	require.NotContains(t, string(output), "FAIL", "%s", output)
	require.Contains(t, string(output), "PASS", "%s", output)
}
