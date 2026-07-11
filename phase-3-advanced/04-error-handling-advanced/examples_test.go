package errors_advanced

import (
	"errors"
	"testing"
)

func TestErrorWrappingAndUnwrapping(t *testing.T) {
	err := PerformOperation()
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	// 1. errors.Is checks if any error in the wrap chain matches the target.
	// Since PerformOperation wraps FetchUser, and FetchUser wraps ErrNotFound,
	// errors.Is(err, ErrNotFound) must be true.
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("Expected errors.Is(err, ErrNotFound) to be true")
	}

	// It should NOT match ErrTimeout
	if errors.Is(err, ErrTimeout) {
		t.Errorf("Expected errors.Is(err, ErrTimeout) to be false")
	}

	// 2. errors.As inspects the chain and finds the first error matching the target pointer type.
	// It extracts the error into that pointer.
	var dbErr *DatabaseError
	if !errors.As(err, &dbErr) {
		t.Fatalf("Expected errors.As(err, &dbErr) to be true")
	}

	// Verify we can access fields of the custom error type
	expectedQuery := "SELECT * FROM users WHERE id = 42"
	if dbErr.Query != expectedQuery {
		t.Errorf("Expected query %q, got %q", expectedQuery, dbErr.Query)
	}

	// Check underlying error inside DatabaseError
	if dbErr.Err != ErrNotFound {
		t.Errorf("Expected underlying error to be ErrNotFound, got %v", dbErr.Err)
	}
}

func TestPanicRecovery(t *testing.T) {
	// 1. Division by non-zero must succeed
	res, err := SafeDivision(10, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if res != 5 {
		t.Errorf("Expected result 5, got %d", res)
	}

	// 2. Division by zero must panic internally, but recover and return an error
	_, err = SafeDivision(10, 0)
	if err == nil {
		t.Fatal("Expected error from divide-by-zero, got nil")
	}

	t.Logf("Successfully recovered and returned error: %v", err)
}
