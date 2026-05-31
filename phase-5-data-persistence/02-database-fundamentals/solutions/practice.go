package solutions

import (
	"database/sql"
	"errors"
)


type Account struct {
	ID      int
	Owner   string
	Balance float64
}

func CreateAccount(db *sql.DB, owner string, balance float64) (int, error) {
	res, err := db.Exec("INSERT INTO accounts (owner, balance) VALUES (?, ?)", owner, balance)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func TransferMoney(db *sql.DB, fromID, toID int, amount float64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var fromBalance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = ?", fromID).Scan(&fromBalance)
	if err != nil {
		return err
	}

	if fromBalance < amount {
		return errors.New("insufficient balance")
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
