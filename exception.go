package goexception

import (
	"errors"
	"fmt"
)

type Exception struct {
	code    Code
	message string
	error   any
}

func (e *Exception) GetError() any {
	return e.error
}

func (e *Exception) GetCode() Code {
	return e.code
}

func (e *Exception) GetMessage() string {
	return e.message
}

func (e *Exception) GetDetailError() error {
	return errors.New(fmt.Sprintf("msg: %s, error: %v", e.message, e.error))
}

func _createException(code Code, message string, errs ...any) *Exception {
	var err any
	if len(errs) > 0 {
		err = errs[0]
	}

	return &Exception{
		code:    code,
		message: message,
		error:   err,
	}
}
