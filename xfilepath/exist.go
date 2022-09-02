// Package xfilepath provides ...
package xfilepath

import (
	"os"
)

// IfNotExists 文件路径 是否存在
func IfNotExists(filePath string, noExistdoFn, existdoFn func(string) error) error {
	_, err := os.Stat(filePath)
	if err == nil {
		if existdoFn != nil {
			return existdoFn(filePath)
		}
		return nil
	}

	// 不存在执行
	if os.IsNotExist(err) {
		if noExistdoFn != nil {
			return noExistdoFn(filePath)
		}
	}

	return err
}
