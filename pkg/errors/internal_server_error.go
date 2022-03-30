package errors

import "net/http"

type InternalServerError struct {
	code    int
	message string
}

func NewInternalServerError(err error) *InternalServerError {
	return &InternalServerError{
		code:    http.StatusInternalServerError,
		message: err.Error(),
	}
}

func (e *InternalServerError) Error() string {
	return e.message
}

func (e *InternalServerError) GetCode() int {
	return e.code
}
