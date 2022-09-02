package xpath

import (
	"os"
	"os/exec"
	"path/filepath"
)

// ExecPath 执行的命令,命令所在的位置
func ExecPath() (string, string, error) {
	arg := os.Args[0]

	dir, err := filepath.Abs(filepath.Dir(arg))
	if err != nil {
		return "", "", nil
	}

	file, err := exec.LookPath(arg)
	if err != nil {
		return "", "", nil
	}

	re, err := filepath.Abs(file)
	if err != nil {
		return "", "", nil
	}

	return dir, re, err
}

// GetThisDirectory 当前执行命令的位置
func GetThisDirectory() (string, error) {
	return os.Getwd()
}
