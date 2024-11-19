// -------------------------------------------
// @file      : errors.go
// @author    : bo cai
// @contact   : caibo923@gmail.com
// @time      : 2024/11/19 上午11:01
// -------------------------------------------

package cberrors

import (
	"bytes"
	"fmt"
	"runtime"
)

// CBError an interface for error with stack trace
type CBError interface {
	error          // inherit from system error interface
	Stack() string // get stack trace message
}

// errorHost an implement of CBError
type errorHost struct {
	stack   string // stack trace message
	message string // error message
}

// Error implement internal error
func (err *errorHost) Error() string {
	if err.message == "" {
		return fmt.Sprintf("stack trace:\n%s", err.stack)
	}
	return fmt.Sprintf("%s\nstack trace:\n%s", err.message, err.stack)
}

// String implement fmt.Stringer
func (err *errorHost) String() string {
	return err.Error()
}

// Stack get stack trace of error
func (err *errorHost) Stack() string {
	return err.stack
}

// stack get current stack trace info
func stack() []byte {
	var buff bytes.Buffer
	for skip := 2; ; skip++ {
		_, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		buff.WriteString(fmt.Sprintf("\tfile = %s, line = %d\n", file, line))
	}
	return buff.Bytes()
}

// New create new CBError object
func New(template string, args ...interface{}) CBError {
	return &errorHost{
		message: fmt.Sprintf(template, args...),
		stack:   string(stack()),
	}
}

// Wrap create new CBError object with an exist error
func Wrap(err error) CBError {
	var message string
	if err != nil {
		message = err.Error()
	}
	return &errorHost{
		message: message,
		stack:   string(stack()),
	}
}

// Panic create new CBError and panic
func Panic(template string, args ...interface{}) {
	panic(New(template, args...))
}

// PanicWrap panic with an error
func PanicWrap(err error) {
	panic(Wrap(err))
}
