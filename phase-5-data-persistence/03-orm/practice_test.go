package orm

import (
	"testing"
	"gorm.io/gorm"
	"github.com/glebarez/sqlite"
)


func TestORM(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	err = db.AutoMigrate(&Customer{}, &Order{})
	if err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	cust, err := CreateCustomerWithOrder(db, "Alice", 250.75)
	if err != nil || cust.ID == 0 {
		t.Fatalf("failed to create customer and order: %v", err)
	}

	retrieved, err := GetCustomerWithOrders(db, cust.ID)
	if err != nil {
		t.Fatalf("failed to get customer: %v", err)
	}

	if len(retrieved.Orders) != 1 || retrieved.Orders[0].Amount != 250.75 {
		t.Errorf("association load failed: %+v", retrieved)
	}
}
