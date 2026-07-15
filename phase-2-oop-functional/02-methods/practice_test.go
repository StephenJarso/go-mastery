package main

import (
	"testing"
	"time"
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

// TestCarGetAge tests the GetAge value receiver method
func TestCarGetAge(t *testing.T) {
	car := Car{Year: 2020}
	expectedAge := time.Now().Year() - 2020
	if car.GetAge() != expectedAge {
		t.Errorf("expected age %d, got %d", expectedAge, car.GetAge())
	}
}

// TestCarNeedsService tests the NeedsService value receiver method
func TestCarNeedsService(t *testing.T) {
	car1 := Car{Miles: 95000}
	car2 := Car{Miles: 100001}

	if car1.NeedsService() {
		t.Error("expected false for 95000 miles")
	}
	if !car2.NeedsService() {
		t.Error("expected true for 100001 miles")
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

// TestCarService tests the Service pointer receiver method
func TestCarService(t *testing.T) {
	car := &Car{Miles: 15000}
	car.Service()
	if car.Miles != 0 {
		t.Errorf("expected 0 miles, got %.0f", car.Miles)
	}
}

// TestCarUpdateYear tests the UpdateYear pointer receiver method
func TestCarUpdateYear(t *testing.T) {
	car := &Car{Year: 2020}
	car.UpdateYear(2022)
	if car.Year != 2022 {
		t.Errorf("expected year 2022, got %d", car.Year)
	}
}

// TestCarBuilderChaining tests method chaining
func TestCarBuilderChaining(t *testing.T) {
	car := NewCarBuilder().
		WithBrand("Ford").
		WithModel("Mustang").
		WithYear(2021).
		WithMiles(5000).
		Build()

	if car.Brand != "Ford" {
		t.Errorf("expected brand 'Ford', got %q", car.Brand)
	}
	if car.Model != "Mustang" {
		t.Errorf("expected model 'Mustang', got %q", car.Model)
	}
	if car.Year != 2021 {
		t.Errorf("expected year 2021, got %d", car.Year)
	}
	if car.Miles != 5000 {
		t.Errorf("expected 5000 miles, got %.0f", car.Miles)
	}
}

