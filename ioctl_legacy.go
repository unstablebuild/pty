//go:build !windows && !go1.12
// +build !windows,!go1.12

package pty

func ioctl(fd uintptr, cmd, ptr uintptr) error {
	return ioctlInner(fd, cmd, ptr) // fall back to blocking io (old behavior)
}
