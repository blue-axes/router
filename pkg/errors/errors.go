package errors

import (
	"fmt"
	"strings"

	"github.com/blue-axes/router/entity/errno"
)

type (
	Error struct {
		code    errno.Errno
		message string
		wrapErr error
	}
)

func (e *Error) Error() string {
	return fmt.Sprintf("code:%d message:%s", e.code, e.message)
}

func (e *Error) Is(err error) bool {
	if e == nil && err == nil {
		return true
	}
	switch valErr := err.(type) {
	case *Error:
		if e.code == valErr.code && e.message == valErr.message {
			return true
		}
		return false
	default:
		return false
	}
}

func WithCode(code errno.Errno, msg ...string) error {
	e := &Error{
		code:    code,
		message: strings.Join(msg, ";"),
	}
	return e
}

func WrapWithCode(err error, code errno.Errno, msg ...string) error {
	e := &Error{
		code:    code,
		message: strings.Join(msg, ";"),
		wrapErr: err,
	}
	return e
}

func UnWrap(err error) error {
	switch valErr := err.(type) {
	case *Error:
		return valErr.wrapErr
	default:
		return err
	}
}
