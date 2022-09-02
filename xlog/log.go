// Package xlog provides ...
package xlog

import (
	"io"
	"log"
	"os"
)

/*
	自定义日志输出
*/

var (
	FatalLog *log.Logger = log.New(os.Stderr, "FATAL ", log.LstdFlags)
	PanicLog *log.Logger = log.New(os.Stderr, "PANIC ", log.LstdFlags)
	ErrorLog *log.Logger = log.New(os.Stderr, "ERROR ", log.LstdFlags)
	WringLog *log.Logger = log.New(os.Stdout, "WRING ", log.LstdFlags)
	InfoLog  *log.Logger = log.New(os.Stdout, "INFO ", log.LstdFlags)
	TraceLog *log.Logger = log.New(os.Stdout, "TRACE ", log.LstdFlags)
)

var (
	Fatal   = FatalLog.Fatal
	Fatalf  = FatalLog.Fatalf
	Fatalln = FatalLog.Fatalln

	Panic   = PanicLog.Panic
	Panicf  = PanicLog.Panicf
	Panicln = PanicLog.Panicln

	Error   = ErrorLog.Print
	Errorf  = ErrorLog.Printf
	Errorln = ErrorLog.Println

	Wring   = WringLog.Print
	Wringf  = WringLog.Printf
	Wringln = WringLog.Println

	Info   = InfoLog.Print
	Infof  = InfoLog.Printf
	Infoln = InfoLog.Println

	Trace   = TraceLog.Print
	Tracef  = TraceLog.Printf
	Traceln = TraceLog.Println
)

// SetOutPut 设置某个日志的输出位置
func SetOutPut(writer io.Writer, loger ...*log.Logger) {
	if len(loger) == 0 {
		return
	}
	for _, v := range loger {
		v.SetOutput(MultiWriter(writer))
	}
	return
}
