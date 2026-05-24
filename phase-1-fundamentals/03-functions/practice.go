package functions

import (
	"errors"
)


var _ = errors.New

// Exercise 1: Collection Filter
// Filter a slice of integers based on a predicate callback.
func Filter(numbers []int, predicate func(int) bool) []int {
	// TODO: Implement
	return nil
}

// Exercise 2: Calculator Closure Factory
// Returns a closure func(int, int) int based on the operation "+", "-", "*".
// Returns error if operation is unsupported.
func NewCalculator(op string) (func(int, int) int, error) {
	// TODO: Implement
	return nil, nil
}

// Exercise 3: Safe Division
// Returns quotient and remainder of dividend / divisor.
// Uses named return parameters. Returns error if divisor is zero.
func SafeDivide(dividend, divisor int) (quotient int, remainder int, err error) {
	// TODO: Implement
	return 0, 0, nil
}
