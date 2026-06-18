package main

import (
	"testing"
)

// TestCarGetInfo tests the GetInfo value receiver method
func TestCarGetInfo(t *testing.T) {
	car := Car{
		Brand: "Toyota",
		Model: "Camry",
		Year:  2015,
		Miles: 95000,
	}

	info := car.GetInfo()
	if info != "2015 Toyota Camry with 95000 miles" {
		t.Errorf("expected '2015 Toyota Camry with 95000 miles', got %q", info)
	}
}

// TestCarDrive tests the Drive pointer receiver method
func TestCarDrive(t *testing.T) {
	car := &Car{
		Brand: "Honda",
		Model: "Civic",
		Year:  2020,
		Miles: 10000,
	}

	car.Drive(5000)

	if car.Miles != 15000 {
		t.Errorf("expected 15000 miles, got %.0f", car.Miles)
	}
}

// TestCarBuilderChaining tests method chaining
func TestCarBuilderChaining(t *testing.T) {
	car := NewCarBuilder().
		WithBrand("Ford").
		Build()

	if car.Brand != "Ford" {
		t.Errorf("expected brand 'Ford', got %q", car.Brand)
	}
}
