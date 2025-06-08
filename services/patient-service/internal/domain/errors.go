// Custom error types
// internal/domain/errors.go
package domain

import "errors"

var (
	// Patient errors
	ErrPatientNotFound      = errors.New("patient not found")
	ErrPatientAlreadyExists = errors.New("patient already exists")
	ErrInvalidPatientData   = errors.New("invalid patient data")

	// General errors
	ErrInvalidInput        = errors.New("invalid input")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrInternalServerError = errors.New("internal server error")
)

// CustomError untuk error yang lebih detail
type CustomError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(code, message, details string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Details: details,
	}
}
