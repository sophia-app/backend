package main

import (
	"github.com/joho/godotenv"
	"github.com/sophia-app/backend/configs"
	"github.com/sophia-app/backend/internal/router"
)

var (
	logger *configs.Logger
)

func main() {
	logger = configs.GetLogger("main")

	err := godotenv.Load()
	if err != nil {
		logger.Errorf("failed to load env file: %v", err)
	}

	err = configs.InitializeDatabase()
	if err != nil {
		logger.Errorf("database initialization failed: %v", err)
		return
	}

	router.Initialize()
}
