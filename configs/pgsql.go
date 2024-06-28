package configs

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitializePostgres initializes a connection to a PostgreSQL database
func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("failed to connect to the database: %v", err)
		return nil, err
	}

	// Migrate schemas
	// err = db.AutoMigrate(&schemas.Company{})
	// if err != nil {
	// 	logger.Errorf("Failed to migrate schema: %v", err)
	// 	return nil, err
	// }

	return db, nil
}
