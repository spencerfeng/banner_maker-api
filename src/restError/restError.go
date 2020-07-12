package resterror

import (
	"fmt"
	"net/http"
)

// RestError ...
type RestError interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restError struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e restError) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s", e.ErrMessage, e.ErrStatus, e.ErrError)
}

func (e restError) Message() string {
	return e.ErrMessage
}

func (e restError) Status() int {
	return e.ErrStatus
}

func (e restError) Causes() []interface{} {
	return e.ErrCauses
}

// NewRestError ...
func NewRestError(message string, status int, err string) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
	}
}

// NewBadRequestError ...
func NewBadRequestError(message string) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

// NewInternalServerError ...
func NewInternalServerError(message string, err error) RestError {
	result := restError{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}
