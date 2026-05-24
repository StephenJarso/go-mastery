package solutions

import (
	"errors"
	"fmt"
	"strconv"
)


func ConvertFahrenheitToCelsius(f float64) (string, error) {
	if f < -459.67 {
		return "", errors.New("temperature is below absolute zero")
	}
	celsius := (f - 32.0) * 5.0 / 9.0
	return fmt.Sprintf("%.1f°C", celsius), nil
}

func CalculateCircleArea(radiusStr string) (float64, error) {
	r, err := strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		return 0, err
	}
	if r < 0 {
		return 0, errors.New("radius cannot be negative")
	}
	const Pi = 3.14159
	return Pi * r * r, nil
}

func BuildUserProfile(name, role, status string) string {
	const SystemName = "GoMastery"
	return fmt.Sprintf("Welcome %s to %s (Role: %s, Status: %s)", name, SystemName, role, status)
}
