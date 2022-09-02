// Package xtools provides ...
package xtools

// IF 三目运算符
func IF(b bool, value1, value2 interface{}) interface{} {
	if b {
		return value1
	}
	return value2
}

// IFFunc 三目运算符 func
func IFFunc(fun func() bool, value1, value2 interface{}) interface{} {
	if fun() {
		return value1
	}
	return value2
}
