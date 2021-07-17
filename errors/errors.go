// Package errors
// The error is normally JSON encoded.
package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Id      string `json:"id"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Error implements the error interface.
func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// New generates a custom error.
func New(id string, code int, status string, message string) error {
	return &Error{
		Id:      id,
		Code:    code,
		Status:  status,
		Message: message,
	}
}

// Parse tries to parse a JSON string into an error.
// If that fails, it will set the given string as the error message.
func Parse(message string) *Error {
	e := &Error{}
	err := json.Unmarshal([]byte(message), e)
	if err != nil {
		e.Message = message
	}
	return e
}

// BadRequest generates a 400 error.
func BadRequest(id string, format string, a ...interface{}) error {
	return &Error{
		Id:      id,
		Code:    http.StatusBadRequest,
		Status:  http.StatusText(http.StatusBadRequest),
		Message: fmt.Sprintf(format, a...),
	}
}

// Unauthorized generates a 401 error.
func Unauthorized(id string, format string, a ...interface{}) error {
	return &Error{
		Id:      id,
		Code:    http.StatusUnauthorized,
		Status:  http.StatusText(http.StatusUnauthorized),
		Message: fmt.Sprintf(format, a...),
	}
}

// Forbidden generates a 403 error.
func Forbidden(id string, format string, a ...interface{}) error {
	return &Error{
		Id:      id,
		Code:    http.StatusForbidden,
		Status:  http.StatusText(http.StatusForbidden),
		Message: fmt.Sprintf(format, a...),
	}
}

// NotFound generates a 404 error.
func NotFound(id string, format string, a ...interface{}) error {
	return &Error{
		Id:      id,
		Code:    http.StatusNotFound,
		Status:  http.StatusText(http.StatusNotFound),
		Message: fmt.Sprintf(format, a...),
	}
}

// MethodNotAllowed generates a 405 error.
func MethodNotAllowed(id string, format string, a ...interface{}) error {
	return &Error{
		Id:      id,
		Code:    http.StatusMethodNotAllowed,
		Status:  http.StatusText(http.StatusMethodNotAllowed),
		Message: fmt.Sprintf(format, a...),
	}
}

// RequestTimeout generates a 408 error.
func RequestTimeout(id string, format string, a ...interface{}) error {
	return &Error{
		Id:      id,
		Code:    http.StatusRequestTimeout,
		Status:  http.StatusText(http.StatusRequestTimeout),
		Message: fmt.Sprintf(format, a...),
	}
}

// Conflict generates a 409 error.
func Conflict(id string, format string, a ...interface{}) error {
	return &Error{
		Id:      id,
		Code:    http.StatusConflict,
		Status:  http.StatusText(http.StatusConflict),
		Message: fmt.Sprintf(format, a...),
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(id string, format string, a ...interface{}) error {
	return &Error{
		Id:      id,
		Code:    http.StatusInternalServerError,
		Status:  http.StatusText(http.StatusInternalServerError),
		Message: fmt.Sprintf(format, a...),
	}
}
