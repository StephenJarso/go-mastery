package pointers

import (
	"testing"
)


func TestSwap(t *testing.T) {
	x, y := 5, 10
	Swap(&x, &y)
	if x != 10 || y != 5 {
		t.Errorf("expected swap result 10 and 5, got %d and %d", x, y)
	}
}

func TestIncrementCounter(t *testing.T) {
	c := 5
	IncrementCounter(&c, 3)
	if c != 8 {
		t.Errorf("expected 8, got %d", c)
	}
}

func TestUpdateGrade(t *testing.T) {
	s := Student{Name: "Bob", Grade: 85.0}
	err := UpdateGrade(&s, 95.5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.Grade != 95.5 {
		t.Errorf("expected grade 95.5, got %f", s.Grade)
	}

	err = UpdateGrade(&s, 105.0)
	if err == nil {
		t.Error("expected error for grade > 100.0")
	}
}
