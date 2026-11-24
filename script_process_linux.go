package whichkey

import "golang.org/x/sys/unix"

func scriptSocketPeer(fd int) (int, int, error) {
	credentials, err := unix.GetsockoptUcred(fd, unix.SOL_SOCKET, unix.SO_PEERCRED)
	if err != nil {
		return 0, 0, err
	}
	return int(credentials.Pid), int(credentials.Uid), nil
}
