package variablestypes

import (
	"testing"
)

func TestShowDeclarations(t *testing.T) {
	output := ShowDeclarations()
	expected := "App: GoMastery, User: stephenjarso, Age: 30, Email: stephenjacob815@gmail.com"
	if output != expected {
		t.Errorf("ShowDeclarations() = %q; expected %q", output, expected)
	}
}

func TestConvertStringToInt(t *testing.T) {
	val, err := ConvertStringToInt("42")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 42 {
		t.Errorf("expected 42, got %d", val)
	}

	_, err = ConvertStringToInt("invalid")
	if err == nil {
		t.Error("expected error for non-numeric input")
	}
}

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
