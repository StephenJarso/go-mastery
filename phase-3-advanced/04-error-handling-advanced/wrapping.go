package errors_advanced

import (
	"errors"
	"fmt"
	"time"
)

// In Go, errors are values. Go 1.13 introduced error wrapping, which allows
// adding context to an error while preserving the original error type and value.

// 1. Sentinel Errors
// Sentinel errors are package-level variables representing specific error conditions.
var (
	ErrNotFound = errors.New("resource not found")
	ErrTimeout  = errors.New("operation timed out")
)

// 2. Custom Error Types
// A custom error is any type that implements the error interface:
//   type error interface { Error() string }

// DatabaseError represents a custom error that provides database-specific context.
type DatabaseError struct {
	Query     string
	Timestamp time.Time
	Err       error // The underlying cause
}

// Error implements the error interface.
func (e *DatabaseError) Error() string {
	return fmt.Sprintf("database error: query %q failed: %v (at %s)", 
		e.Query, e.Err, e.Timestamp.Format(time.RFC3339))
}

// Unwrap returns the underlying error. This is required for errors.Is and errors.As to work!
func (e *DatabaseError) Unwrap() error {
	return e.Err
}

// FetchUser simulates a database fetch operation that might return wrapped errors.
func FetchUser(id int) error {
	// Simulate query failure
	underlyingErr := ErrNotFound

	// Wrap the sentinel error inside a DatabaseError custom struct.
	return &DatabaseError{
		Query:     fmt.Sprintf("SELECT * FROM users WHERE id = %d", id),
		Timestamp: time.Now(),
		Err:       underlyingErr,
	}
}

// PerformOperation wraps an error with format string %w.
func PerformOperation() error {
	err := FetchUser(42)
	if err != nil {
		// Using %w allows unwrapping later. Using %v does NOT support unwrapping.
		return fmt.Errorf("failed to perform operation: %w", err)
	}
	return nil
}
