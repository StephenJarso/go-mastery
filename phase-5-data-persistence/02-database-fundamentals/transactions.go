package dbfundamentals

import (
	"database/sql"
	"fmt"
)

// Database Transactions ensure ACID compliance (Atomicity, Consistency, Isolation, Durability).
// If a business operation requires multiple SQL statements (e.g. subtracting balance from Account A
// and adding it to Account B), they must either ALL succeed or ALL fail.

// TransferStudentGradeTransaction runs an atomic transaction that moves a student
// to a new grade and logs the modification in a audit_log table.
func TransferStudentGradeTransaction(db *sql.DB, studentID int64, newGrade string) error {
	// 1. Start the transaction.
	// db.Begin returns a *sql.Tx (Transaction handle) which supports Exec, Query, etc.
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	// 2. CRITICAL TRANSACTION PATTERN: Defer Rollback.
	// If the transaction finishes successfully and tx.Commit() is called,
	// tx.Rollback() does nothing (returns sql.ErrTxDone, which is ignored).
	// If a panic or early return occurs due to an error, Rollback is automatically triggered,
	// reversing any database changes.
	defer tx.Rollback()

	// Create audit table if not exists (for demo integrity)
	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS audit_log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		student_id INTEGER,
		old_grade TEXT,
		new_grade TEXT
	);`)
	if err != nil {
		return fmt.Errorf("failed to create audit table: %w", err)
	}

	// Retrieve the student's old grade first (querying within the transaction)
	var oldGrade string
	err = tx.QueryRow("SELECT grade FROM students WHERE id = ?", studentID).Scan(&oldGrade)
	if err != nil {
		return fmt.Errorf("failed to fetch current grade: %w", err)
	}

	// Perform Action 1: Update the student's grade
	updateQuery := "UPDATE students SET grade = ? WHERE id = ?"
	_, err = tx.Exec(updateQuery, newGrade, studentID)
	if err != nil {
		return fmt.Errorf("failed to update student grade: %w", err)
	}

	// Perform Action 2: Write log to audit
	insertLogQuery := "INSERT INTO audit_log (student_id, old_grade, new_grade) VALUES (?, ?, ?)"
	_, err = tx.Exec(insertLogQuery, studentID, oldGrade, newGrade)
	if err != nil {
		return fmt.Errorf("failed to write audit log: %w", err)
	}

	// 3. Commit the transaction.
	// If this succeeds, changes are saved permanently.
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
