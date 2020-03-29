package goeasy

import (
	"fmt"
)

const (
	ErrCodeBadRequest   = 400
	ErrCodeUnauthorized = 401
	ErrCodeAccessDenied = 403
	ErrCodeNotFound     = 404
	ErrCodeInternal     = 500
)

type Error interface {
	Code() int
	Error() string
}

type BasicError struct {
	StatusCode int    `json:"-"`
	ErrorText  string `json:"error"`
}

func (h *BasicError) Code() int {
	return h.StatusCode
}

func (h *BasicError) Error() string {
	return h.ErrorText
}

func NewBasicError(status int, format string, args ...interface{}) *BasicError {
	return &BasicError{StatusCode: status, ErrorText: fmt.Sprintf(format, args...)}
}
