package response

type CodeResponse struct {
	code int
}

func NewCodeResponse(code int) *CodeResponse {
	return &CodeResponse{code}
}

func (r *CodeResponse) GetCode() int {
	return r.code
}

func (r *CodeResponse) GetText() string {
	return ""
}
