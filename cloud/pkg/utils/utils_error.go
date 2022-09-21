package utils

import "errors"

type ApiError struct {
	statusCode int
	Err        error
}

func (e *ApiError) Error() string {
	return e.Err.Error()
}

func (e *ApiError) StatusCode() int {
	return e.statusCode
}

func NewUnexpectedServerError() error {
	return &ApiError{
		statusCode: 500,
		Err:        errors.New("some unexpected error occurred"),
	}
}

func NewNotFoundError(err error) error {
	return &ApiError{
		statusCode: 404,
		Err:        err,
	}
}
