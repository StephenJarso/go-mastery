package pointers

import (
	"testing"
)

func TestIncrement(t *testing.T) {
	x := 10
	
	// Value increment (no change)
	IncrementByValue(x)
	if x != 10 {
		t.Errorf("expected x to remain 10, got %d", x)
	}

	// Pointer increment (changes original value)
	IncrementByPointer(&x)
	if x != 11 {
		t.Errorf("expected x to be 11, got %d", x)
	}
}

func TestModifyFirstElement(t *testing.T) {
	arr := []int{1, 2, 3}
	ModifyFirstElement(arr, 99)
	if arr[0] != 99 {
		t.Errorf("expected first element to be 99, got %d", arr[0])
	}
}

func TestSwap(t *testing.T) {
	a, b := 100, 200
	err := Swap(&a, &b)
	if err != nil {
		t.Fatalf("unexpected error swapping: %v", err)
	}
	if a != 200 || b != 100 {
		t.Errorf("expected swap values to be 200 and 100, got a=%d, b=%d", a, b)
	}

	err = Swap(nil, &b)
	if err == nil {
		t.Error("expected error swapping nil pointer")
	}
}
