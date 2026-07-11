package generics

// PRACTICE EXERCISE #1: Generic Minimum Finder
// Implement a generic function called FindMin that accepts a slice of numeric types
// (integers or floats) and returns the minimum value in the slice.
// If the slice is empty, return the zero value of the type.

type OrderedNumber interface {
	int | int32 | int64 | float32 | float64
}

func FindMin[T OrderedNumber](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero
	}

	minVal := slice[0]
	for _, val := range slice {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

// PRACTICE EXERCISE #2: Generic Set Data Structure
// Implement a generic Set data structure called PracticeSet.
// A set is a collection of unique elements. It should support:
// - Add(item T)
// - Remove(item T)
// - Has(item T) bool
// - Size() int
// - List() []T (returns all elements in the set)

type PracticeSet[T comparable] struct {
	elements map[T]struct{}
}

func NewPracticeSet[T comparable]() *PracticeSet[T] {
	return &PracticeSet[T]{
		elements: make(map[T]struct{}),
	}
}

func (s *PracticeSet[T]) Add(item T) {
	s.elements[item] = struct{}{}
}

func (s *PracticeSet[T]) Remove(item T) {
	delete(s.elements, item)
}

func (s *PracticeSet[T]) Has(item T) bool {
	_, ok := s.elements[item]
	return ok
}

func (s *PracticeSet[T]) Size() int {
	return len(s.elements)
}

func (s *PracticeSet[T]) List() []T {
	list := make([]T, 0, len(s.elements))
	for key := range s.elements {
		list = append(list, key)
	}
	return list
}
