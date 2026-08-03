package jsonprocessing

import (
	"encoding/json"
	"errors"
	"time"
)

var _ = json.Marshal
var _ = time.Now
var _ = errors.New

type Config struct {
	AppName string `json:"app_name"`
	Port    int    `json:"port"`
}

type PracticeCustomDate struct {
	time.Time
}

type PracticeCustomDateConfig struct {
	Created PracticeCustomDate `json:"created"`
}

// Exercise 1: Parse Config JSON
// Unmarshal data into Config struct. Return error if empty field.
func ParseConfig(data []byte) (Config, error) {
	// TODO: Implement
	return Config{}, nil
}

// Exercise 2: Custom JSON Date Unmarshaling
// Implement json.Unmarshaler on PracticeCustomDate to parse dates in format "02-01-2006".
func (cd *PracticeCustomDate) UnmarshalJSON(b []byte) error {
	// TODO: Implement
	return nil
}
