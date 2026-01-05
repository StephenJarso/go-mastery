package testing_basics

import (
	"errors"
	"testing"
)

func TestSumOfEvens(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty slice", []int{}, 0},
		{"no evens", []int{1, 3, 5}, 0},
		{"all evens", []int{2, 4, 6}, 12},
		{"mixed numbers", []int{1, 2, 3, 4, 5, 6}, 12},
		{"negative evens", []int{-2, -4, 3}, -6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SumOfEvens(tc.input)
			if got != tc.expected {
				t.Errorf("SumOfEvens(%v) = %d; expected %d", tc.input, got, tc.expected)
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name          string
		email         string
		expectedError error
	}{
		{
			name:          "Valid Email",
			email:         "user@example.com",
			expectedError: nil,
		},
		{
			name:          "Empty Email",
			email:         "",
			expectedError: ErrEmailEmpty,
		},
		{
			name:          "Missing @ Sign",
			email:         "userexample.com",
			expectedError: ErrEmailMissingAt,
		},
		{
			name:          "Multiple @ Signs",
			email:         "user@extra@example.com",
			expectedError: ErrEmailMultipleAt,
		},
		{
			name:          "Empty Local Part",
			email:         "@example.com",
			expectedError: ErrEmailEmptyLocal,
		},
		{
			name:          "Missing Domain Dot",
			email:         "user@example",
			expectedError: ErrEmailInvalidDom,
		},
		{
			name:          "Dot At Domain Start",
			email:         "user@.com",
			expectedError: ErrEmailInvalidDom,
		},
		{
			name:          "Dot At Domain End",
			email:         "user@example.",
			expectedError: ErrEmailInvalidDom,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateEmail(tc.email)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("ValidateEmail(%q) error = %v; expected %v", tc.email, err, tc.expectedError)
			}
		})
	}
}
