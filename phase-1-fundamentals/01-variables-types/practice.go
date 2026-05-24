package variablestypes

import (
	"errors"
)


var _ = errors.New

// Exercise 1: Fahrenheit to Celsius
// Convert Fahrenheit to Celsius. Celsius = (Fahrenheit - 32) * 5/9
// Returns a string representation like "25.5°C" or an error if temperature is below absolute zero (-459.67°F).
func ConvertFahrenheitToCelsius(f float64) (string, error) {
	// TODO: Implement
	return "", nil
}

// Exercise 2: Circle Area Calculator
// Given a radius as a string, parse it to float64, calculate area using Pi = 3.14159 (as a constant),
// and return it. Return error if string is not a valid float or if radius is negative.
func CalculateCircleArea(radiusStr string) (float64, error) {
	// TODO: Implement
	return 0, nil
}

// Exercise 3: User Profile Builder
// Given user profile values, use the constant SystemName = "GoMastery" to format a welcome string.
// Return: "Welcome <Name> to <SystemName> (Role: <Role>, Status: <Status>)"
func BuildUserProfile(name, role, status string) string {
	// TODO: Implement
	return ""
}
