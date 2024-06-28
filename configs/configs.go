package configs

import (
	"gorm.io/gorm"
)

var (
	logger *Logger
	db     *gorm.DB
)

// GetLogger returns a logger instance
func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

// GetDatabase returns a database instance
func GetDatabase() *gorm.DB {
	return db
}

// InitializeDatabase initializes the database
func InitializeDatabase() error {
	var err error

	db, err = InitializePostgres()
	if err != nil {
		return err
	}

	return nil
}
