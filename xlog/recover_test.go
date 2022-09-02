package xlog

import "testing"

func TestRecoverShow(t *testing.T) {
	defer RecoverShow()
	Panic("recover")
	defer RecoverShow()
	Fatal("fatal")
}

func TestRecover(t *testing.T) {
	defer Recover()
	Panic("recover")
	defer Recover()
	// Fatal("fatal")
}
