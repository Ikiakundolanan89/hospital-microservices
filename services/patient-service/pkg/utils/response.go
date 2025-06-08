// Response helpers
// pkg/utils/response.go
package utils

import (
	"patient-service/internal/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// SuccessResponse returns a success response
func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(dto.SuccessResponse{
		Message: message,
		Data:    data,
	})
}

// ErrorResponse returns an error response
func ErrorResponse(c *fiber.Ctx, status int, code, message, details string) error {
	return c.Status(status).JSON(dto.ErrorResponse{
		Error: dto.ErrorDetail{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

// ValidationErrorResponse returns validation error response
func ValidationErrorResponse(c *fiber.Ctx, err error) error {
	var errors []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors = append(errors, formatValidationError(e))
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
		Error: dto.ErrorDetail{
			Code:    "VALIDATION_ERROR",
			Message: "Validation failed",
			Details: joinErrors(errors),
		},
	})
}

func formatValidationError(e validator.FieldError) string {
	field := e.Field()
	tag := e.Tag()

	switch tag {
	case "required":
		return field + " is required"
	case "min":
		return field + " must be at least " + e.Param()
	case "max":
		return field + " must be at most " + e.Param()
	case "len":
		return field + " must be exactly " + e.Param() + " characters"
	case "email":
		return field + " must be a valid email"
	case "oneof":
		return field + " must be one of: " + e.Param()
	default:
		return field + " is invalid"
	}
}

func joinErrors(errors []string) string {
	result := ""
	for i, err := range errors {
		if i > 0 {
			result += "; "
		}
		result += err
	}
	return result
}
