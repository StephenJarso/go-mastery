package generics

import (
	"testing"
)


func TestMapSlice(t *testing.T) {
	s := []int{1, 2, 3}
	res := MapSlice(s, func(x int) string { return string(rune('A' + x - 1)) })
	if len(res) != 3 || res[0] != "A" || res[1] != "B" || res[2] != "C" {
		t.Errorf("MapSlice failed: %v", res)
	}
}

func TestGenericStack(t *testing.T) {
	var s PracticeStack[int]
	s.Push(1)
	s.Push(2)
	val, err := s.Pop()
	if err != nil || val != 2 {
		t.Errorf("Pop failed: %v, %v", val, err)
	}
	if s.Len() != 1 {
		t.Errorf("len failed: %d", s.Len())
	}
}

func TestFindMin(t *testing.T) {
	s := []float64{3.14, 1.5, 2.71}
	val, err := FindMin(s)
	if err != nil || val != 1.5 {
		t.Errorf("FindMin failed: %v, %v", val, err)
	}
}
