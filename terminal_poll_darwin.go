//go:build darwin

package whichkey

import (
	"time"

	"golang.org/x/sys/unix"
)

// Darwin's /dev/tty alias is usable for I/O but poll reports it as invalid.
// select supports the controlling terminal. EOF and hangup still surface through
// the subsequent read/write; neither readiness nor a timeout is disconnection.
func pollTerminal(fds []unix.PollFd, timeout int) (int, error) {
	var readable, writable unix.FdSet
	highest := -1
	for index := range fds {
		descriptor := &fds[index]
		descriptor.Revents = 0
		fd := int(descriptor.Fd)
		// FdSet.Set indexes a fixed array. Fail safely rather than panic or corrupt
		// memory if this process has exhausted select's descriptor range.
		if fd < 0 || fd >= unix.FD_SETSIZE {
			return 0, unix.EBADF
		}
		if descriptor.Events & ^int16(unix.POLLIN|unix.POLLOUT) != 0 {
			return 0, unix.EINVAL
		}
		if descriptor.Events&unix.POLLIN != 0 {
			readable.Set(fd)
		}
		if descriptor.Events&unix.POLLOUT != 0 {
			writable.Set(fd)
		}
		highest = max(highest, fd)
	}
	delay := unix.NsecToTimeval(int64(timeout) * int64(time.Millisecond))
	_, err := unix.Select(highest+1, &readable, &writable, nil, &delay)
	if err != nil {
		return 0, err
	}
	count := 0
	for index := range fds {
		descriptor := &fds[index]
		fd := int(descriptor.Fd)
		if readable.IsSet(fd) {
			descriptor.Revents |= unix.POLLIN
		}
		if writable.IsSet(fd) {
			descriptor.Revents |= unix.POLLOUT
		}
		if descriptor.Revents != 0 {
			count++
		}
	}
	return count, nil
}
