// Package xerrors provides ...
package xerrors

import "errors"

var _ error = (*Error)(nil)

// Error 错误消息
// 适用于 返回结果中含有错误和消息的情况
type Error struct {
	msg  string
	data interface{}
	err  error
}

func ToError(err error) *Error {
	return &Error{
		err: err,
	}
}

func New(err, msg string, data interface{}) *Error {
	return &Error{
		msg:  msg,
		data: data,
		err:  errors.New(err),
	}
}

func Msg(err, msg string) *Error {
	return &Error{
		msg: msg,
		err: errors.New(err),
	}
}

func Data(err string, data interface{}) *Error {
	return &Error{
		data: data,
		err:  errors.New(err),
	}
}

func (e *Error) IsNull() bool {
	return e.err == nil
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) Messge() string {
	return e.msg
}

func (e *Error) Data() interface{} {
	return e.data
}
