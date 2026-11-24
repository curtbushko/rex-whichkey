package whichkey

import "golang.org/x/sys/unix"

func scriptSocketPeer(fd int) (int, int, error) {
	credentials, err := unix.GetsockoptXucred(fd, unix.SOL_LOCAL, unix.LOCAL_PEERCRED)
	if err != nil {
		return 0, 0, err
	}
	pid, err := unix.GetsockoptInt(fd, unix.SOL_LOCAL, unix.LOCAL_PEERPID)
	return pid, int(credentials.Uid), err
}
