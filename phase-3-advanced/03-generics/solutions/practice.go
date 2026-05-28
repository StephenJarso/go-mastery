package solutions

import (
	"errors"
	"golang.org/x/exp/constraints"
)


func MapSlice[T any, U any](slice []T, f func(T) U) []U {
	res := make([]U, len(slice))
	for i, v := range slice {
		res[i] = f(v)
	}
	return res
}

type PracticeStack[T any] struct {
	items []T
}

func (s *PracticeStack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *PracticeStack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	idx := len(s.items) - 1
	item := s.items[idx]
	s.items = s.items[:idx]
	return item, nil
}

func (s *PracticeStack[T]) Len() int {
	return len(s.items)
}

func FindMin[T constraints.Ordered](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, errors.New("empty slice")
	}
	min := slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
	}
	return min, nil
}
