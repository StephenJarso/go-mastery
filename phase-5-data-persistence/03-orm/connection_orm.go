package orm

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitORM initializes the GORM database handle, configures logging,
// and runs migrations to build the tables in SQLite.
func InitORM(dsn string) (*gorm.DB, error) {
	// Configure logging settings
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable color
		},
	)

	// 1. Open connection via the SQLite driver wrapper
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to GORM database: %w", err)
	}

	// 2. Configure connection pool (GORM uses standard database/sql connection pool under the hood)
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetConnMaxLifetime(1 * time.Hour)
	}

	// 3. Auto Migration: Creates tables, missing columns, constraints, and indexes.
	// It will NOT delete or modify existing columns (to avoid data loss).
	err = db.AutoMigrate(&User{}, &Profile{}, &Post{}, &Tag{})
	if err != nil {
		return nil, fmt.Errorf("auto migration failed: %w", err)
	}

	return db, nil
}
