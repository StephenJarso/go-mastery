package generics

import (
	"errors"
	"golang.org/x/exp/constraints"
)


var _ = errors.New

// Exercise 1: Generic Map Slice
// MapSlice transforms a slice of type T to U using function f.
func MapSlice[T any, U any](slice []T, f func(T) U) []U {
	// TODO: Implement
	return nil
}

// Exercise 2: Generic Stack
// A generic stack structure with Push, Pop, and Len methods.
type PracticeStack[T any] struct {
	items []T
}

func (s *PracticeStack[T]) Push(v T) {
	// TODO: Implement
}

func (s *PracticeStack[T]) Pop() (T, error) {
	// TODO: Implement
	var zero T
	return zero, nil
}

func (s *PracticeStack[T]) Len() int {
	// TODO: Implement
	return 0
}

// Exercise 3: Generic Min Finder
// Find the smallest element in a slice of Ordered constraint types.
func FindMin[T constraints.Ordered](slice []T) (T, error) {
	// TODO: Implement
	var zero T
	return zero, nil
}
