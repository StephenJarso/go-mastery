package dbfundamentals

import (
	"database/sql"
	"fmt"
	"time"

	// Side-effect import to register the pure-Go SQLite driver with database/sql.
	_ "github.com/glebarez/go-sqlite"
)

// Go's standard library provides a SQL database abstraction layer: "database/sql".
// It is driver-agnostic; you write SQL queries using the standard package,
// and load the driver of choice (PostgreSQL, MySQL, SQLite, etc.) via side-effect imports.

// ConnectDatabase establishes a connection pool to an SQLite database.
func ConnectDatabase(dsn string) (*sql.DB, error) {
	// 1. sql.Open initializes the database pool.
	// It doesn't actually connect to the database right away; it just configures the pool.
	// Driver name is "sqlite" (registered by github.com/glebarez/go-sqlite).
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// 2. Configure connection pool parameters.
	// Connection pooling is built directly into sql.DB and is fully thread-safe.
	
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	// For SQLite, having multiple concurrent write connections can cause locks,
	// so a low limit (or 1 for exclusive write locking) is typical, though read concurrency is fine.
	db.SetMaxOpenConns(10)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(5)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.SetConnMaxLifetime(1 * time.Hour)

	// 3. Ping tests the connection to verify that the DSN is valid and the server is reachable.
	// This is the actual point where a database connection is attempted.
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
