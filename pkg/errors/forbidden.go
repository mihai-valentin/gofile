package errors

import "net/http"

type ForbiddenError struct {
	code    int
	message string
}

func NewForbiddenError(err error) *ForbiddenError {
	return &ForbiddenError{
		code:    http.StatusForbidden,
		message: err.Error(),
	}
}

func (e *ForbiddenError) Error() string {
	return e.message
}

func (e *ForbiddenError) GetCode() int {
	return e.code
}
