package api

type Error struct {
	Code string `json:"code"`
	Msg  string `json:msg`
	Err  string `json:"err"`
}

func (e *Error) Error() string {
	return e.Msg
}

const (
	CodeNotFound   = "not_found"
	CodeBadRequest = "bad_request"
	CodeInternal   = "internal_error"
)

func NotFound(msg string) *Error {
	return &Error{
		Code: CodeNotFound,
		Msg:  msg,
	}
}

func BadRequest(msg string, err error) *Error {
	return &Error{
		Code: CodeBadRequest,
		Msg:  msg,
		Err:  err.Error(),
	}
}

func InternalError(msg string, err error) *Error {
	return &Error{
		Code: CodeBadRequest,
		Msg:  msg,
		Err:  err.Error(),
	}
}
