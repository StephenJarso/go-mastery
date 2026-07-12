package dbfundamentals

import (
	"database/sql"
	"testing"

	_ "github.com/glebarez/go-sqlite"
)

func TestAccountManager(t *testing.T) {
	// Create an in-memory SQLite database for test runs
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	mgr := NewAccountManager(db)

	// 1. Create table
	if err := mgr.CreateAccountsTable(); err != nil {
		t.Fatalf("failed to create accounts table: %v", err)
	}

	// 2. Create accounts
	accA, err := mgr.CreateAccount("Stephen", 500.00)
	if err != nil {
		t.Fatalf("failed to create account A: %v", err)
	}

	accB, err := mgr.CreateAccount("Jacob", 100.00)
	if err != nil {
		t.Fatalf("failed to create account B: %v", err)
	}

	// 3. Successful fund transfer
	err = mgr.TransferFunds(accA, accB, 150.00)
	if err != nil {
		t.Fatalf("TransferFunds failed: %v", err)
	}

	balA, _ := mgr.GetBalance(accA)
	balB, _ := mgr.GetBalance(accB)

	if balA != 350.00 {
		t.Errorf("expected A's balance to be 350.00, got %f", balA)
	}
	if balB != 250.00 {
		t.Errorf("expected B's balance to be 250.00, got %f", balB)
	}

	// 4. Failed transfer (insufficient funds)
	err = mgr.TransferFunds(accA, accB, 1000.00)
	if err != ErrInsufficientFunds {
		t.Errorf("expected ErrInsufficientFunds, got %v", err)
	}

	// Verify balances did not change after rollback
	balA, _ = mgr.GetBalance(accA)
	if balA != 350.00 {
		t.Errorf("expected balance to remain 350.00, got %f", balA)
	}

	// 5. Non-existent account transfer
	err = mgr.TransferFunds(999, accB, 50.00)
	if err == nil {
		t.Error("expected transfer from invalid account to fail")
	}
}
