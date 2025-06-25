package errors

import (
	"fmt"
	"net/http"
)

// AppError adalah custom error type untuk aplikasi
type AppError struct {
	Code    int    // HTTP status code
	Message string // Error message
	Err     error  // Original error (optional)
}

// Error implements error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the wrapped error
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewNotFoundError creates a not found error
func NewNotFoundError(message string, err ...error) *AppError {
	var originalErr error
	if len(err) > 0 {
		originalErr = err[0]
	}
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
		Err:     originalErr,
	}
}

// NewBadRequestError creates a bad request error
func NewBadRequestError(message string, err ...error) *AppError {
	var originalErr error
	if len(err) > 0 {
		originalErr = err[0]
	}
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
		Err:     originalErr,
	}
}

// NewInternalServerError creates an internal server error
func NewInternalServerError(message string, err ...error) *AppError {
	var originalErr error
	if len(err) > 0 {
		originalErr = err[0]
	}
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
		Err:     originalErr,
	}
}

// Function lain untuk tipe error yang Anda butuhkan...

func NewUnauthorizedError(message string, err ...error) *AppError {
	var originalErr error
	if len(err) > 0 {
		originalErr = err[0]
	}
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
		Err:     originalErr,
	}
}
