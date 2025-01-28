package input

import (
	"regexp"
	"strings"
)

func SanitizeString(input string) string {
	// Remove any HTML tags
	htmlTagRegex := regexp.MustCompile(`<[^>]*>`)
	sanitized := htmlTagRegex.ReplaceAllString(input, "")

	// Trim spaces
	sanitized = strings.TrimSpace(sanitized)

	return sanitized
}