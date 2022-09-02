// Package xlog provides ...
package xlog

import (
	"testing"
)

func TestLog(t *testing.T) {
	msg := "abcd"

	// func() {
	// 	defer reciverShow()
	// 	Fatal(msg)
	// }()

	func() {
		defer RecoverShow()
		Panic(msg)
	}()
	func() {
		defer RecoverShow()
		Error(msg)
	}()

	func() {
		defer RecoverShow()
		Wring(msg)
	}()
	func() {
		defer RecoverShow()
		Info(msg)
	}()
	func() {
		defer RecoverShow()
		Trace(msg)
	}()
}
