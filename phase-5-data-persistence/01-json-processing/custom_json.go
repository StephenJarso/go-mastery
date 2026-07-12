package jsonprocessing

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Sometimes, standard JSON serialization is not sufficient. E.g., custom time formats,
// API contracts with custom date formatting, or translating dynamic types.
// We can implement custom JSON behavior by implementing:
// - json.Marshaler interface:   MarshalJSON() ([]byte, error)
// - json.Unmarshaler interface: UnmarshalJSON([]byte) error

// CustomDate wraps time.Time to support "YYYY-MM-DD" JSON formatting instead of RFC3339.
type CustomDate struct {
	time.Time
}

const dateFormat = "2006-01-02"

// MarshalJSON implements the json.Marshaler interface.
// It must return a valid JSON value (e.g. double-quoted string).
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	// Format time as "YYYY-MM-DD"
	formatted := cd.Time.Format(dateFormat)
	// Surround with double quotes to make it a valid JSON string.
	jsonStr := fmt.Sprintf(`%q`, formatted)
	return []byte(jsonStr), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It parses a double-quoted "YYYY-MM-DD" string back to time.Time.
func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	// Remove outer quotes from the JSON string.
	str := string(data)
	str = strings.Trim(str, `"`)

	// Parse date string
	parsedTime, err := time.Parse(dateFormat, str)
	if err != nil {
		return fmt.Errorf("failed to parse CustomDate: %w", err)
	}

	cd.Time = parsedTime
	return nil
}

type Event struct {
	Name string     `json:"name"`
	Date CustomDate `json:"date"`
}
