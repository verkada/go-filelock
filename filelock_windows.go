//go:build windows

package filelock

import (
	"golang.org/x/sys/windows"
)

func lockFile(fd int) error {
	var overlapped windows.Overlapped
	return windows.LockFileEx(windows.Handle(fd), windows.LOCKFILE_EXCLUSIVE_LOCK, 0, 1, 0, &overlapped)
}

func unlockFile(fd int) error {
	var overlapped windows.Overlapped
	return windows.UnlockFileEx(windows.Handle(fd), 0, 1, 0, &overlapped)
}
