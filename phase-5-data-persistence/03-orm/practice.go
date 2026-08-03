package orm

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name   string
	Orders []Order
}

type Order struct {
	gorm.Model
	CustomerID uint
	Amount     float64
}

// Exercise 1: Create Customer and Order
// Create a new Customer with associated Order using GORM.
func CreateCustomerWithOrder(db *gorm.DB, name string, orderAmount float64) (*Customer, error) {
	// TODO: Implement
	return nil, nil
}

// Exercise 2: Retrieve Customer with Eager Loading
// Retrieve Customer by ID with all associated Orders preloaded.
func GetCustomerWithOrders(db *gorm.DB, customerID uint) (*Customer, error) {
	// TODO: Implement
	return nil, nil
}
