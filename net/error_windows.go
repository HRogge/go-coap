//go:build windows

package net

import (
	"errors"
	"syscall"
)

func IsConnError(err error) bool {
	return errors.Is(err, syscall.WSAECONNRESET) ||
		errors.Is(err, syscall.WSAECONNABORTED)
}
