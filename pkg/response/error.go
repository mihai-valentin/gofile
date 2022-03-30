package response

import (
	"fmt"
	"time"
)

type ErrorResponse struct {
	code    int
	message string
}

func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		code:    code,
		message: message,
	}
}

func (r *ErrorResponse) GetCode() int {
	return r.code
}

func (r *ErrorResponse) GetText() string {
	return fmt.Sprintf("[%d] [%s] %s", r.code, time.Now().String(), r.message)
}
