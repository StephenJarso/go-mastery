package orm

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// PRACTICE EXERCISE: E-commerce Order Tracker
// Define Customer, Order, and OrderItem GORM schemas, and implement transactional order creation.
//
// Associations:
// - Customer has many Orders.
// - Order belongs to Customer and has many OrderItems.
// - OrderItem belongs to Order.

type Customer struct {
	gorm.Model
	Name   string  `gorm:"not null"`
	Email  string  `gorm:"uniqueIndex;not null"`
	Orders []Order `gorm:"foreignKey:CustomerID"`
}

type Order struct {
	gorm.Model
	CustomerID  uint
	OrderDate   time.Time
	TotalAmount float64
	Items       []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID     uint
	ProductName string  `gorm:"not null"`
	Quantity    int     `gorm:"not null"`
	UnitPrice   float64 `gorm:"not null"`
}

// CreateOrder places an order for a customer and calculates totals atomically inside a transaction.
func CreateOrder(db *gorm.DB, customerID uint, items []OrderItem) (*Order, error) {
	if len(items) == 0 {
		return nil, errors.New("cannot place an order with zero items")
	}

	var order Order

	// Start GORM transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		// 1. Verify Customer exists
		var c Customer
		if err := tx.First(&c, customerID).Error; err != nil {
			return fmt.Errorf("customer not found: %w", err)
		}

		// 2. Calculate TotalAmount and configure items
		var total float64
		for i := range items {
			if items[i].Quantity <= 0 || items[i].UnitPrice <= 0 {
				return errors.New("invalid item quantity or unit price")
			}
			total += float64(items[i].Quantity) * items[i].UnitPrice
		}

		order = Order{
			CustomerID:  customerID,
			OrderDate:   time.Now(),
			TotalAmount: total,
			Items:       items,
		}

		// 3. Save order (GORM automatically inserts OrderItems and sets their OrderID foreign key!)
		if err := tx.Create(&order).Error; err != nil {
			return fmt.Errorf("failed to save order: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &order, nil
}
