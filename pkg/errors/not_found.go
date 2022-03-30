package errors

import "net/http"

type NotFoundError struct {
	code    int
	message string
}

func NewNotFoundError(err error) *NotFoundError {
	return &NotFoundError{
		code:    http.StatusNotFound,
		message: err.Error(),
	}
}

func (e *NotFoundError) Error() string {
	return e.message
}

func (e *NotFoundError) GetCode() int {
	return e.code
}
