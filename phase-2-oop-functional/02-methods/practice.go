package main

import (
	"fmt"
	"time"
)

// PRACTICE EXERCISES for Phase 2.2: Methods & Receivers

// Car is a struct for practice
type Car struct {
	Brand string
	Model string
	Year  int
	Miles float64
}

// Exercise 1: Value Receiver Methods
// Create methods that READ but don't MODIFY

// TODO: Create these value receiver methods:
// 1. GetInfo() string - returns formatted car info
// 2. GetAge() int - returns age of car (current year - model year)
// 3. NeedsService() bool - returns true if miles > 100000

// Solution for GetInfo (example)
func (c Car) GetInfo() string {
	return fmt.Sprintf("%d %s %s with %.0f miles", c.Year, c.Brand, c.Model, c.Miles)
}

// GetAge returns age of car (current year - model year)
func (c Car) GetAge() int {
	return time.Now().Year() - c.Year
}

// NeedsService returns true if miles > 100000
func (c Car) NeedsService() bool {
	return c.Miles > 100000
}

// Exercise 2: Pointer Receiver Methods
// Create methods that MODIFY the receiver

// TODO: Create these pointer receiver methods:
// 1. Drive(miles float64) - adds miles to the car
// 2. Service() - resets miles to 0
// 3. UpdateYear(year int) - updates the year

// Solution for Drive (example)
func (c *Car) Drive(miles float64) {
	if miles > 0 {
		c.Miles += miles
	}
}

// Service resets miles to 0
func (c *Car) Service() {
	c.Miles = 0
}

// UpdateYear updates the year of the car
func (c *Car) UpdateYear(year int) {
	c.Year = year
}

// Exercise 3: Method Chaining
// TODO: Implement a builder for creating Car objects

type CarBuilder struct {
	car *Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{car: &Car{}}
}

// TODO: Add these builder methods (return *CarBuilder for chaining):
// 1. WithBrand(brand string) *CarBuilder
// 2. WithModel(model string) *CarBuilder
// 3. WithYear(year int) *CarBuilder
// 4. WithMiles(miles float64) *CarBuilder
// 5. Build() *Car - returns final car

func (cb *CarBuilder) WithBrand(brand string) *CarBuilder {
	cb.car.Brand = brand
	return cb
}

// WithModel sets the car's model
func (cb *CarBuilder) WithModel(model string) *CarBuilder {
	cb.car.Model = model
	return cb
}

// WithYear sets the car's year
func (cb *CarBuilder) WithYear(year int) *CarBuilder {
	cb.car.Year = year
	return cb
}

// WithMiles sets the car's miles
func (cb *CarBuilder) WithMiles(miles float64) *CarBuilder {
	cb.car.Miles = miles
	return cb
}

func (cb *CarBuilder) Build() *Car {
	return cb.car
}

// Example usage
func PracticeExercises() {
	fmt.Println("=== Practice Exercises: Methods & Receivers ===")

	// Exercise 1: Value Receiver Methods
	fmt.Println("\n--- Exercise 1: Value Receiver Methods ---")
	myCar := Car{
		Brand: "Toyota",
		Model: "Camry",
		Year:  2015,
		Miles: 95000,
	}

	fmt.Printf("Car: %s\n", myCar.GetInfo())
	fmt.Printf("Car Age: %d years\n", myCar.GetAge())
	fmt.Printf("Needs service? %v\n", myCar.NeedsService())

	// Exercise 2: Pointer Receiver Methods
	fmt.Println("\n--- Exercise 2: Pointer Receiver Methods ---")
	myCar.Drive(6000)
	fmt.Printf("After driving 6000 miles: %s\n", myCar.GetInfo())
	fmt.Printf("Needs service after drive? %v\n", myCar.NeedsService())
	myCar.Service()
	fmt.Printf("After service: %s\n", myCar.GetInfo())
	myCar.UpdateYear(2018)
	fmt.Printf("After updating year: %s\n", myCar.GetInfo())

	// Exercise 3: Method Chaining
	fmt.Println("\n--- Exercise 3: Method Chaining ---")
	builtCar := NewCarBuilder().
		WithBrand("Honda").
		WithModel("Civic").
		WithYear(2020).
		WithMiles(15000).
		Build()

	fmt.Printf("Built car: %s\n", builtCar.GetInfo())
}

func main() {
	PracticeExercises()
	fmt.Println("\n=== Complete the TODOs above ===")
}
