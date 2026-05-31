package dbfundamentals

import (
	"testing"
	"database/sql"
	_ "github.com/glebarez/go-sqlite"
)


func TestDatabaseFundamentals(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE accounts (id INTEGER PRIMARY KEY AUTOINCREMENT, owner TEXT, balance REAL)")
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	id1, err := CreateAccount(db, "Alice", 100.0)
	id2, err := CreateAccount(db, "Bob", 50.0)
	if err != nil || id1 == 0 || id2 == 0 {
		t.Fatalf("failed to create accounts: %v", err)
	}

	err = TransferMoney(db, id1, id2, 30.0)
	if err != nil {
		t.Fatalf("transfer failed: %v", err)
	}

	var bal1, bal2 float64
	db.QueryRow("SELECT balance FROM accounts WHERE id = ?", id1).Scan(&bal1)
	db.QueryRow("SELECT balance FROM accounts WHERE id = ?", id2).Scan(&bal2)

	if bal1 != 70.0 || bal2 != 80.0 {
		t.Errorf("wrong balances: bal1=%f, bal2=%f", bal1, bal2)
	}
}
