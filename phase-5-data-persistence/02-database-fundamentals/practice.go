package dbfundamentals

import (
	"database/sql"
	"errors"
)


var _ = sql.Open
var _ = errors.New

type Account struct {
	ID      int
	Owner   string
	Balance float64
}

// Exercise 1: Create Account
// Insert new account and return its primary key ID.
func CreateAccount(db *sql.DB, owner string, balance float64) (int, error) {
	// TODO: Implement
	return 0, nil
}

// Exercise 2: Transfer Transactional Money
// Move amount from one account to another within a Transaction. Balance cannot be negative.
func TransferMoney(db *sql.DB, fromID, toID int, amount float64) error {
	// TODO: Implement
	return nil
}
