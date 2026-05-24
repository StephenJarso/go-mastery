package functions

import (
	"testing"
)


func TestFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	res := Filter(nums, func(n int) bool { return n%2 == 1 })
	expected := []int{1, 3, 5}
	if len(res) != len(expected) {
		t.Fatalf("expected 3 elements, got %d", len(res))
	}
	for i, v := range res {
		if v != expected[i] {
			t.Errorf("at index %d: got %d, expected %d", i, v, expected[i])
		}
	}
}

func TestNewCalculator(t *testing.T) {
	add, err := NewCalculator("+")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res := add(3, 4); res != 7 {
		t.Errorf("expected 7, got %d", res)
	}

	_, err = NewCalculator("invalid")
	if err == nil {
		t.Error("expected error for invalid op")
	}
}

func TestSafeDivide(t *testing.T) {
	q, r, err := SafeDivide(10, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if q != 3 || r != 1 {
		t.Errorf("expected q=3, r=1, got q=%d, r=%d", q, r)
	}

	_, _, err = SafeDivide(10, 0)
	if err == nil {
		t.Error("expected error dividing by zero")
	}
}
