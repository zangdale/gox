package xlog

import (
	"io"
	"io/ioutil"
	"os"
)

var (
	Stdout  = os.Stdout
	Stderr  = os.Stderr
	Discard = ioutil.Discard
	NULL    = io.Discard
)

// MultiWriter 合并输出位置
func MultiWriter(writer ...io.Writer) io.Writer {
	if len(writer) == 0 {
		return ioutil.Discard
	}
	return io.MultiWriter(writer...)
}

// WriterFile 输出到文件
func WriterFile(name string) io.Writer {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Fatalf("writer file %s", err.Error())
	}
	return f
}
