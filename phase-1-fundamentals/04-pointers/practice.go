package pointers

import (
	"errors"
)


var _ = errors.New

type Student struct {
	Name  string
	Grade float64
}

// Exercise 1: Safe Swap
// Swap the values of two integer pointers in place.
func Swap(a, b *int) {
	// TODO: Implement
}

// Exercise 2: Increment Counter
// Increment the integer value pointed to by val by the specified amount.
func IncrementCounter(val *int, amount int) {
	// TODO: Implement
}

// Exercise 3: Update Student Grade
// Update the student's grade through a pointer. Returns error if newGrade is not between 0.0 and 100.0.
func UpdateGrade(s *Student, newGrade float64) error {
	// TODO: Implement
	return nil
}
