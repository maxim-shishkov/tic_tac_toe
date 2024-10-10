package api

const (
	CodeNotFound       = "NOT_FOUND"
	CodeInternalError  = "INTERNAL_ERROR"
	CodeInvalidRequest = "INVALID_REQUEST"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Err     string `json:"err"`
}

func NotFound(msg string, err error) Error {
	return Error{
		Code:    CodeNotFound,
		Message: msg,
		Err:     err.Error(),
	}
}

func InternalError(msg string, err error) Error {
	return Error{
		Code:    CodeInternalError,
		Message: msg,
		Err:     err.Error(),
	}
}

func InvalidRequest(msg string, err error) Error {
	return Error{
		Code:    CodeInvalidRequest,
		Message: msg,
		Err:     err.Error(),
	}
}
