package simply

import (
	"github.com/xssnick/goeasy"
)

func ErrBadRequest(text string) *goeasy.BasicError {
	return goeasy.NewBasicError(goeasy.ErrCodeBadRequest, text)
}

func ErrAccessDenied(text string) *goeasy.BasicError {
	return goeasy.NewBasicError(goeasy.ErrCodeAccessDenied, text)
}

func ErrUnauthorized(text string) *goeasy.BasicError {
	return goeasy.NewBasicError(goeasy.ErrCodeUnauthorized, text)
}

func ErrInternal(text string) *goeasy.BasicError {
	return goeasy.NewBasicError(goeasy.ErrCodeInternal, text)
}

func ErrNotFound(text string) *goeasy.BasicError {
	return goeasy.NewBasicError(goeasy.ErrCodeNotFound, text)
}