//go:build windows
// +build windows

package pty

// Winsize is a dummy struct to enable compilation on unsupported platforms.
type Winsize struct {
	Rows, Cols, X, Y uint16
}

// Setsize resizes t to s.
func Setsize(uintptr, *Winsize) error {
	return ErrUnsupported
}

// GetsizeFull returns the full terminal size description.
func GetsizeFull(uintptr) (*Winsize, error) {
	return nil, ErrUnsupported
}
