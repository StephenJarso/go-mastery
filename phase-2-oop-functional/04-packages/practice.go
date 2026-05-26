package phase2packages

import (
	"strings"
	"strconv"
	"time"
	"errors"
)


var _ = strings.ToLower
var _ = strconv.Itoa
var _ = time.Now
var _ = errors.New

// Exercise 1: String URL Sanitizer
// Convert title to lowercase, replace spaces with hyphens, remove non-alphanumeric chars.
func FormatURLPath(title string) string {
	// TODO: Implement
	return ""
}

// Exercise 2: Parse and Sum CSV
// Given a CSV line like "10,20,30", parse elements and return sum.
func ParseAndSumCSV(line string) (int, error) {
	// TODO: Implement
	return 0, nil
}

// Exercise 3: Age in Days
// Calculate age in days between birthdate (YYYY-MM-DD) and current date.
func DaysSinceBirth(birthdate string, now time.Time) (int, error) {
	// TODO: Implement
	return 0, nil
}
