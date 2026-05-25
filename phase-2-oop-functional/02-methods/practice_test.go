package methods

import (
	"testing"
)


func TestCarBuilder(t *testing.T) {
	car := NewCarBuilder().Brand("Tesla").Model("Model 3").Speed(150).Build()
	if car.Brand != "Tesla" || car.Model != "Model 3" || car.Speed != 150 {
		t.Errorf("car builder failed: %+v", car)
	}
}

func TestCounterMethods(t *testing.T) {
	var c Counter
	c.Increment()
	c.Increment()
	if c.GetValue() != 2 {
		t.Errorf("expected count 2, got %d", c.GetValue())
	}
}

func TestIntSliceSum(t *testing.T) {
	s := IntSlice{1, 2, 3, 4, 5}
	if s.Sum() != 15 {
		t.Errorf("expected sum 15, got %d", s.Sum())
	}
}
