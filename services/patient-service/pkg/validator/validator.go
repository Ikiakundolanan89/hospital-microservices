// Input validation

// pkg/validator/validator.go
package validator

import (
	"github.com/go-playground/validator/v10"
)

// New creates a new validator instance with custom validations
func New() *validator.Validate {
	validate := validator.New()

	// Register custom validations here if needed
	// Example:
	// validate.RegisterValidation("customTag", customValidationFunc)

	return validate
}
