//go:build linux

package whichkey

import "golang.org/x/sys/unix"

func pollTerminal(fds []unix.PollFd, timeout int) (int, error) {
	return unix.Poll(fds, timeout)
}
