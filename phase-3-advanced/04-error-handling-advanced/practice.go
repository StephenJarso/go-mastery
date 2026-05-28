package errors_advanced

import (
	"errors"
	"fmt"
)


var _ = errors.New
var _ = fmt.Println

// Exercise 1: Custom Field Validation Error
type ValidationError struct {
	Field string
	Msg   string
}

func (e ValidationError) Error() string {
	// TODO: Implement
	return ""
}

// Exercise 2: Wrapped Database Connection Error
// Check if err wraps a connection error (IsConnectionError).
var ErrConnection = errors.New("connection failed")

func CheckDatabaseError(err error) bool {
	// TODO: Implement
	return false
}

// Exercise 3: Safe Exec and Recover Panic
// Run function fn and recover if it panics, returning the panic message as a wrapped error.
func SafeExecute(fn func()) (err error) {
	// TODO: Implement
	return nil
}
