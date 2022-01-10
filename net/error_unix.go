//go:build darwin || linux

package net

import (
	"errors"
	"syscall"
)

func IsConnError(err error) bool {
	return errors.Is(err, syscall.EPIPE) ||
		errors.Is(err, syscall.ECONNRESET) ||
		errors.Is(err, syscall.ECONNABORTED)
}
