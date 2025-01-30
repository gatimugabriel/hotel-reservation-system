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

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format: %v", err)
	}

	return date, nil
}

// ValidateCheckInDate ensures the check-in date is not in the past
func ValidateCheckInDate(checkInDate time.Time) error {
	today := time.Now()
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())

	if checkInDate.Before(today) {
		return fmt.Errorf("check-in date should start today onwards")
	}
	return nil
}

// ValidateCheckInCheckOutDates ensures checkout is after checkin
func ValidateCheckInCheckOutDates(checkIn, checkOut time.Time) error {
	if checkOut.Before(checkIn) {
		return fmt.Errorf("check-out date must be after check-in date")
	}
	return nil
}

// ParseAndValidateCheckInDate combines parsing and validation for check-in dates
func ParseAndValidateCheckInDate(dateStr string) (time.Time, error) {
	date, err := ParseDate(dateStr)
	if err != nil {
		return time.Time{}, err
	}

	if err := ValidateCheckInDate(date); err != nil {
		return time.Time{}, err
	}

	return date, nil
}

// GetDateFromURL combines getting param and parsing date
func GetDateFromURL(r *http.Request, paramName string) (time.Time, error) {
	dateStr := GetParamFromURL(r, paramName)
	//return ParseDate(dateStr)

	date, err := ParseDate(dateStr)
	if err != nil {
		return time.Time{}, err
	}

	if paramName == "check_in" {
		if err := ValidateCheckInDate(date); err != nil {
			return time.Time{}, err
		}
	}

	return date, nil
}