package xerrors

func HaveErrorToPrint_(printFunc func(args ...interface{}), err error, msgs ...interface{}) {
	HaveErrorPrint(printFunc, ToError(err), msgs)
}

func HaveErrorPrint(printFunc func(args ...interface{}), err *Error, msgs ...interface{}) {
	if err.IsNull() {
		return
	}
	printFunc(err, msgs)
}
