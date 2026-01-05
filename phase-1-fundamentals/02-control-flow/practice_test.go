package controlflow

import (
	"testing"
)

func TestSumUpTo(t *testing.T) {
	if val := SumUpTo(10); val != 55 {
		t.Errorf("expected SumUpTo(10) to be 55, got %d", val)
	}
}

func TestGetGradeName(t *testing.T) {
	tests := []struct {
		score    int
		expected string
	}{
		{105, "Superb"},
		{90, "Excellent"},
		{85, "Very Good"},
		{72, "Good"},
		{60, "Pass"},
		{50, "Fail"},
	}

	for _, tt := range tests {
		res := GetGradeName(tt.score)
		if res != tt.expected {
			t.Errorf("GetGradeName(%d) = %q; expected %q", tt.score, res, tt.expected)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	res := FizzBuzz(1, 15)
	expected := []string{
		"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz",
	}

	if len(res) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(res))
	}

	for i := range res {
		if res[i] != expected[i] {
			t.Errorf("at index %d: expected %q, got %q", i, expected[i], res[i])
		}
	}
}
