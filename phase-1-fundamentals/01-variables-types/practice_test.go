package variablestypes

import (
	"testing"
)


func TestConvertFahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
		hasError bool
	}{
		{32.0, "0.0°C", false},
		{212.0, "100.0°C", false},
		{-40.0, "-40.0°C", false},
		{98.6, "37.0°C", false},
		{-500.0, "", true},
	}

	for _, tt := range tests {
		result, err := ConvertFahrenheitToCelsius(tt.input)
		if (err != nil) != tt.hasError {
			t.Errorf("ConvertFahrenheitToCelsius(%f) error presence expected %v, got %v", tt.input, tt.hasError, err)
		}
		if result != tt.expected {
			t.Errorf("ConvertFahrenheitToCelsius(%f) = %q; expected %q", tt.input, result, tt.expected)
		}
	}
}

func TestCalculateCircleArea(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		hasError bool
	}{
		{"10.0", 314.159, false},
		{"0.0", 0.0, false},
		{"-5.0", 0.0, true},
		{"invalid", 0.0, true},
	}

	for _, tt := range tests {
		result, err := CalculateCircleArea(tt.input)
		if (err != nil) != tt.hasError {
			t.Errorf("CalculateCircleArea(%q) error presence expected %v, got %v", tt.input, tt.hasError, err)
		}
		if !tt.hasError && (result < tt.expected-0.001 || result > tt.expected+0.001) {
			t.Errorf("CalculateCircleArea(%q) = %f; expected %f", tt.input, result, tt.expected)
		}
	}
}

func TestBuildUserProfile(t *testing.T) {
	res := BuildUserProfile("Alice", "Admin", "Active")
	expected := "Welcome Alice to GoMastery (Role: Admin, Status: Active)"
	if res != expected {
		t.Errorf("BuildUserProfile() = %q; expected %q", res, expected)
	}
}
