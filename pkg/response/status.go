package response

type StatusResponse struct {
	code    int
	message string
}

func NewStatusResponse(code int, message string) *StatusResponse {
	return &StatusResponse{
		code:    code,
		message: message,
	}
}

func (r *StatusResponse) GetCode() int {
	return r.code
}

func (r *StatusResponse) GetText() string {
	return r.message
}
