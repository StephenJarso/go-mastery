package solutions

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

func CreateCustomerWithOrder(db *gorm.DB, name string, orderAmount float64) (*Customer, error) {
	cust := &Customer{
		Name: name,
		Orders: []Order{
			{Amount: orderAmount},
		},
	}
	err := db.Create(cust).Error
	if err != nil {
		return nil, err
	}
	return cust, nil
}

func GetCustomerWithOrders(db *gorm.DB, customerID uint) (*Customer, error) {
	var cust Customer
	err := db.Preload("Orders").First(&cust, customerID).Error
	if err != nil {
		return nil, err
	}
	return &cust, nil
}
