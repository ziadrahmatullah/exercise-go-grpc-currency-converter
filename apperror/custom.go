package apperror

import (
	"google.golang.org/grpc/codes"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorRes struct {
	Message string `json:"message"`
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func (ce *CustomError) Error() string {
	return ce.Message
}

func (ce *CustomError) ToErrorRes() ErrorRes {
	return ErrorRes{
		Message: ce.Message,
	}
}

var (
	ErrUnknownTarget = NewCustomError(int(codes.InvalidArgument), "unknown target")
	ErrUnknownSource = NewCustomError(int(codes.InvalidArgument), "unknown source")
)
