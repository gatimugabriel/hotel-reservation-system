package utils

import (
	"fmt"
	"net/http"
	"time"
)

// ParseDate converts a date string to time.Time.
// Expects date in format "2006-01-02" (YYYY-MM-DD)
func ParseDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, fmt.Errorf("date string is empty")
	}

	// Parse the date string
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format: %v", err)
	}

	return date, nil
}

// GetDateFromURL combines getting param and parsing date
func GetDateFromURL(r *http.Request, paramName string) (time.Time, error) {
	dateStr := GetParamFromURL(r, paramName)
	return ParseDate(dateStr)
}