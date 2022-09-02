// Package xlog provides ...
package xlog

func Recover() interface{} {
	return recover()
}

// RecoverShow 输出 recover
func RecoverShow() {
	// recover() 的范围是一个{} ,在此使用 Recover() 无效
	if msg := recover(); msg != nil {
		Errorf("recover show %v", msg)
	}
	return
}
