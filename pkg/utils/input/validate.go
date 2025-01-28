package input

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationResponse struct {
	Errors []ValidationError `json:"errors"`
}

// GetValidationMessage returns a user-friendly message for validation tags
func GetValidationMessage(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Must be a valid email address"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	case "e164":
		return "Must be a valid phone number in E.164 format"
	case "passwd":
		return "Password must contain at least 8 characters, one uppercase letter, one lowercase letter, and one number"
	case "oneof":
		return "Must be one of the allowed values"
	default:
		return "Invalid value"
	}
}

// ValidateStruct validates a given struct and returns structured validation errors
func ValidateStruct(data interface{}) *ValidationResponse {
	validate := validator.New()

	// custom password validation
	_ = validate.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
		hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)

		return len(password) >= 8 && hasUpper && hasLower && hasNumber
	})

	err := validate.Struct(data)
	if err != nil {
		var errors []ValidationError

		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			tag := err.Tag()
			param := err.Param()

			message := GetValidationMessage(tag)
			if tag == "min" || tag == "max" {
				message = fmt.Sprintf("%s (%s characters)", message, param)
			}

			errors = append(errors, ValidationError{
				Field:   field,
				Message: message,
			})
		}

		return &ValidationResponse{Errors: errors}
	}

	return nil
}