package variablestypes

import (
	"errors"
	"fmt"
)

// PRACTICE EXERCISE: Temperature Converter
// Write a function that converts Fahrenheit to Celsius and returns a formatted string.
// Formula: Celsius = (Fahrenheit - 32) * 5/9
// Returns a string representation like "25.5°C" or an error if temperature is below absolute zero (-459.67°F).

func ConvertFahrenheitToCelsius(f float64) (string, error) {
	if f < -459.67 {
		return "", errors.New("temperature is below absolute zero")
	}
	celsius := (f - 32.0) * 5.0 / 9.0
	return fmt.Sprintf("%.1f°C", celsius), nil
}
