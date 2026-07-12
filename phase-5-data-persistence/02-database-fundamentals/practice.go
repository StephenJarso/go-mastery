package dbfundamentals

import (
	"database/sql"
	"errors"
	"fmt"
)

// PRACTICE EXERCISE #1 & #2: Account Manager and Transactional Fund Transfer
// Implement a simple banking ledger database manager.
//
// DB Schema requirements:
// Table name: accounts
// Columns:
//  - id (INTEGER PRIMARY KEY AUTOINCREMENT)
//  - owner (TEXT NOT NULL)
//  - balance (REAL NOT NULL CHECK(balance >= 0)) // Check constraint prevents negative balances!
//
// Implement functions for creating the table, creating accounts, getting balances,
// and transferring funds safely inside an SQL transaction.

var (
	ErrInsufficientFunds = errors.New("insufficient funds for transfer")
	ErrAccountNotFound   = errors.New("account not found")
)

type AccountManager struct {
	db *sql.DB
}

func NewAccountManager(db *sql.DB) *AccountManager {
	return &AccountManager{db: db}
}

// CreateAccountsTable creates the accounts table with a check constraint.
func (am *AccountManager) CreateAccountsTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS accounts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		owner TEXT NOT NULL,
		balance REAL NOT NULL CHECK(balance >= 0)
	);`
	_, err := am.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create accounts table: %w", err)
	}
	return nil
}

// CreateAccount registers a new bank account with an initial deposit.
func (am *AccountManager) CreateAccount(owner string, initialBalance float64) (int64, error) {
	if initialBalance < 0 {
		return 0, errors.New("initial balance cannot be negative")
	}

	query := `INSERT INTO accounts (owner, balance) VALUES (?, ?)`
	res, err := am.db.Exec(query, owner, initialBalance)
	if err != nil {
		return 0, fmt.Errorf("failed to insert account: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get account ID: %w", err)
	}

	return id, nil
}

// GetBalance returns the balance of a specific account.
func (am *AccountManager) GetBalance(accountID int64) (float64, error) {
	query := `SELECT balance FROM accounts WHERE id = ?`
	var balance float64

	err := am.db.QueryRow(query, accountID).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, ErrAccountNotFound
		}
		return 0, fmt.Errorf("query balance failed: %w", err)
	}

	return balance, nil
}

// TransferFunds transfers money from one account to another atomically inside a transaction.
func (am *AccountManager) TransferFunds(fromID, toID int64, amount float64) error {
	if amount <= 0 {
		return errors.New("transfer amount must be greater than zero")
	}

	// 1. Begin the transaction
	tx, err := am.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transfer transaction: %w", err)
	}
	// 2. Defer rollback for safety
	defer tx.Rollback()

	// 3. Fetch sender balance inside transaction (applying lock if necessary, but SQLite handles it)
	var fromBalance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = ?", fromID).Scan(&fromBalance)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("sender %w", ErrAccountNotFound)
		}
		return err
	}

	// 4. Verify funds
	if fromBalance < amount {
		return ErrInsufficientFunds
	}

	// 5. Check receiver existence
	var toBalance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = ?", toID).Scan(&toBalance)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("receiver %w", ErrAccountNotFound)
		}
		return err
	}

	// 6. Subtract from sender
	_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromID)
	if err != nil {
		return fmt.Errorf("failed to subtract funds: %w", err)
	}

	// 7. Add to receiver
	_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toID)
	if err != nil {
		return fmt.Errorf("failed to add funds: %w", err)
	}

	// 8. Commit!
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transfer: %w", err)
	}

	return nil
}
