//go:build !windows

package filelock

import (
	"syscall"
)

func lockFile(fd int) error {
	return syscall.Flock(fd, syscall.LOCK_EX)
}

func unlockFile(fd int) error {
	return syscall.Flock(fd, syscall.LOCK_UN)
}
