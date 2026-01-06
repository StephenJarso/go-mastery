package functions

import (
	"testing"
)

func TestDivide(t *testing.T) {
	q, r, err := Divide(10, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if q != 3 || r != 1 {
		t.Errorf("expected quotient 3 and remainder 1, got q=%d, r=%d", q, r)
	}

	_, _, err = Divide(10, 0)
	if err == nil {
		t.Error("expected error dividing by zero")
	}
}

func TestCalculateStats(t *testing.T) {
	min, max := CalculateStats([]int{5, 2, 9, -1, 7})
	if min != -1 || max != 9 {
		t.Errorf("expected min -1 and max 9, got min=%d, max=%d", min, max)
	}
}

func TestMakeMultiplier(t *testing.T) {
	double := MakeMultiplier(2)
	triple := MakeMultiplier(3)

	if val := double(5); val != 10 {
		t.Errorf("expected double(5) = 10, got %d", val)
	}
	if val := triple(5); val != 15 {
		t.Errorf("expected triple(5) = 15, got %d", val)
	}
}

func TestFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Test case: filter even numbers
	even := Filter(nums, func(x int) bool {
		return x%2 == 0
	})

	expectedEven := []int{2, 4, 6, 8, 10}
	if len(even) != len(expectedEven) {
		t.Fatalf("expected length %d, got %d", len(expectedEven), len(even))
	}
	for i := range even {
		if even[i] != expectedEven[i] {
			t.Errorf("expected %d at index %d, got %d", expectedEven[i], i, even[i])
		}
	}
}
