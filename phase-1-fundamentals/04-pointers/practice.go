package pointers

import (
	"errors"
)

// PRACTICE EXERCISE: Safe Swapper
// Write a function Swap that swaps the values of two integers using pointers.
// If either pointer is nil, return an error.

func Swap(a, b *int) error {
	if a == nil || b == nil {
		return errors.New("cannot swap nil pointers")
	}
	temp := *a
	*a = *b
	*b = temp
	return nil
}
