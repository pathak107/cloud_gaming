package apierrors

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	Code    int
	Err     error
	ErrMsg  string
	Context string //just to provide some extra information about the origin of the error
}

func (e *ApiError) Error() string {
	return e.ErrMsg
}

func (e *ApiError) ErrorOriginal() error {
	return fmt.Errorf("origin: %v , err: %v", e.Context, e.Err)
}

func NewServerError(err error, info string) error {
	return &ApiError{
		Code:    http.StatusInternalServerError,
		Err:     err,
		ErrMsg:  "Some unexpected internal error occured",
		Context: info,
	}
}

func New(err error, code int, msg string, info string) error {
	return &ApiError{
		Code:    code,
		Err:     err,
		ErrMsg:  msg,
		Context: info,
	}
}

func NewBadRequestError(err error, info string, msg string) error {
	return &ApiError{
		Code:    http.StatusBadRequest,
		Err:     err,
		ErrMsg:  msg,
		Context: info,
	}
}

func NewResourceNotFoundError(err error, info string, msg string) error {
	return &ApiError{
		Code:    http.StatusNotFound,
		Err:     err,
		ErrMsg:  msg,
		Context: info,
	}
}

func NewUnauthorizedError(err error, info string) error {
	return &ApiError{
		Code:    http.StatusUnauthorized,
		Err:     err,
		ErrMsg:  "Unauthorized access",
		Context: info,
	}
}
