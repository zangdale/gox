// Package xos provides ...
package xos

import (
	"os"
	"runtime"
)

// IsRoot 是否是 管理员
func IsRoot() bool {
	switch runtime.GOOS {
	case "linux":
		return os.Getuid() == 0
	case "windows":
		_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
		return err == nil
	case "darwin":
		return os.Getuid() == 0
	default:
		return os.Getuid() == 0
	}
}
