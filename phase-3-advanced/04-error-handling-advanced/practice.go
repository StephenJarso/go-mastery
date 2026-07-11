package errors_advanced

import (
	"fmt"
	"strings"
	"time"
)

// PRACTICE EXERCISE #1: Custom ValidationError
// Define a custom error type called ValidationError. It should hold a list of
// validation error messages for specific fields (e.g., map[string]string or similar).
// Implement the error interface so it prints all field validation errors clearly.

type ValidationError struct {
	Errors map[string]string
}

func (ve *ValidationError) Error() string {
	var sb strings.Builder
	sb.WriteString("validation failed:")
	for field, errMsg := range ve.Errors {
		sb.WriteString(fmt.Sprintf(" [%s: %s]", field, errMsg))
	}
	return sb.String()
}

// PRACTICE EXERCISE #2: Retry Pattern with Error Wrapping
// Implement a function called RetryOperation that accepts a function f to run,
// and a maximum number of retries. It should call f, and if f returns an error,
// wait 5 milliseconds and try again, up to maxRetries.
// If all attempts fail, return a wrapped error containing the final error.

func RetryOperation(f func() error, maxRetries int) error {
	var err error
	for i := 0; i <= maxRetries; i++ {
		err = f()
		if err == nil {
			return nil
		}
		if i < maxRetries {
			time.Sleep(5 * time.Millisecond)
		}
	}
	return fmt.Errorf("operation failed after %d retries: %w", maxRetries, err)
}

// PRACTICE EXERCISE #3: Panic Recovery Runner
// Implement a function called SafeRun that accepts a function f to execute.
// It should run f and, if f panics, recover from it and return the panic as an error.
// If f executes successfully without panicking, return nil.

func SafeRun(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	f()
	return nil
}
