package errors_advanced

import (
	"fmt"
)

// Panics and Recovery
//
// - A panic stops the ordinary flow of control and begins panicking.
//   It walks up the call stack, running deferred functions, and then crashes.
// - A recover call stops the panicking sequence and retrieves the value passed to panic.
//   It is only useful inside deferred functions.
//
// Guidelines:
// 1. Don't use panic for normal error handling. Return error values instead.
// 2. Use panic only for unrecoverable errors (e.g., failed startup configs, nil pointer dereferences).
// 3. If writing a library, recover from internal panics and return them as errors.

// SafeDivision divides two integers, recovering from a divide-by-zero panic.
func SafeDivision(dividend, divisor int) (result int, err error) {
	// Defer a function that checks for panic and recovers.
	defer func() {
		if r := recover(); r != nil {
			// Recover returns interface{} of the panic value.
			// We convert the panic into an error value.
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()

	// If divisor is 0, Go runtime will panic with a divide-by-zero error.
	result = dividend / divisor
	return result, nil
}

// RaisePanic demonstrates throwing a manual panic.
func RaisePanic() {
	panic("something went critically wrong")
}
