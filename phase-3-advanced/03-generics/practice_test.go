package generics

import (
	"testing"
)

func TestFindMin(t *testing.T) {
	// 1. Int slice
	ints := []int{10, 5, 20, -3, 8}
	if got := FindMin(ints); got != -3 {
		t.Errorf("Expected -3, got %d", got)
	}

	// 2. Float slice
	floats := []float64{3.14, 1.59, 2.65, 0.58}
	if got := FindMin(floats); got != 0.58 {
		t.Errorf("Expected 0.58, got %f", got)
	}

	// 3. Empty slice
	var empty []int
	if got := FindMin(empty); got != 0 {
		t.Errorf("Expected 0, got %d", got)
	}
}

func TestPracticeSet(t *testing.T) {
	s := NewPracticeSet[string]()

	// 1. Add and Size
	s.Add("apple")
	s.Add("banana")
	s.Add("apple") // duplicate, should not add
	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}

	// 2. Has
	if !s.Has("apple") {
		t.Error("Expected set to contain 'apple'")
	}
	if s.Has("cherry") {
		t.Error("Expected set NOT to contain 'cherry'")
	}

	// 3. Remove
	s.Remove("apple")
	if s.Has("apple") {
		t.Error("Expected 'apple' to be removed")
	}
	if s.Size() != 1 {
		t.Errorf("Expected size 1, got %d", s.Size())
	}

	// 4. List
	list := s.List()
	if len(list) != 1 || list[0] != "banana" {
		t.Errorf("Expected list ['banana'], got %v", list)
	}
}
