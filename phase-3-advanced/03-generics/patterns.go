package generics

import "sync"

// 1. Generic Slice Operations (Map and Filter)

// Map applies a function to all elements in a slice and returns a slice containing the results.
func Map[T any, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

// Filter filters elements of a slice using a predicate function.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// 2. Generic Data Structure: Thread-Safe Stack
// Stacks are LIFO (Last In, First Out) data structures.

type Stack[T any] struct {
	mu    sync.Mutex
	items []T
}

// NewStack creates a new empty Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, item)
}

// Pop removes and returns the top item of the stack.
// Returns the zero value of T and false if empty.
func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var zero T
	if len(s.items) == 0 {
		return zero, false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index] // shrink slice
	return item, true
}

// Size returns the number of items in the stack.
func (s *Stack[T]) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items)
}
