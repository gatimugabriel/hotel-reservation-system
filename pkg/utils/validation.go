package utils

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidEmail    = errors.New("invalid email format")
	ErrInvalidPassword = errors.New("password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, and one number")
	ErrInvalidName     = errors.New("name must contain only letters and be between 2 and 50 characters")
)

func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return ErrInvalidPassword
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

	if !hasUpper || !hasLower || !hasNumber {
		return ErrInvalidPassword
	}

	return nil
}

func ValidateName(name string) error {
	name = strings.TrimSpace(name)
	if len(name) < 2 || len(name) > 50 {
		return ErrInvalidName
	}

	if !regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(name) {
		return ErrInvalidName
	}

	return nil
}

func SanitizeString(input string) string {
	// Remove any HTML tags
	htmlTagRegex := regexp.MustCompile(`<[^>]*>`)
	sanitized := htmlTagRegex.ReplaceAllString(input, "")

	// Trim spaces
	sanitized = strings.TrimSpace(sanitized)

	return sanitized
}