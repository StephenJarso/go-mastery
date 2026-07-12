package orm

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestOrderTracker(t *testing.T) {
	// Initialize in-memory GORM database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	// Migrate models
	err = db.AutoMigrate(&Customer{}, &Order{}, &OrderItem{})
	if err != nil {
		t.Fatalf("migration failed: %v", err)
	}

	// Create Customer
	cust := Customer{Name: "StephenJarso", Email: "stephen@example.com"}
	db.Create(&cust)

	// Case 1: Valid Order
	items := []OrderItem{
		{ProductName: "MacBook Pro", Quantity: 1, UnitPrice: 1999.99},
		{ProductName: "USB-C Hub", Quantity: 2, UnitPrice: 49.50},
	}

	order, err := CreateOrder(db, cust.ID, items)
	if err != nil {
		t.Fatalf("CreateOrder failed: %v", err)
	}

	expectedTotal := 1999.99 + (2 * 49.50)
	if order.TotalAmount != expectedTotal {
		t.Errorf("expected total amount %f, got %f", expectedTotal, order.TotalAmount)
	}

	// Verify items saved in DB with correct foreign key
	var savedItems []OrderItem
	db.Where("order_id = ?", order.ID).Find(&savedItems)
	if len(savedItems) != 2 {
		t.Errorf("expected 2 saved items, got %d", len(savedItems))
	}

	// Case 2: Customer not found
	_, err = CreateOrder(db, 999, items)
	if err == nil {
		t.Error("expected placing order for non-existent customer to fail")
	}

	// Case 3: Invalid Item values (negative price)
	badItems := []OrderItem{
		{ProductName: "Freebie", Quantity: 1, UnitPrice: -10.00},
	}
	_, err = CreateOrder(db, cust.ID, badItems)
	if err == nil {
		t.Error("expected validation to fail for negative prices")
	}
}
