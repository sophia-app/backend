package utils

import (
	"github.com/sophia-app/backend/configs"
	"gorm.io/gorm"
)

var (
	logger *configs.Logger
	db *gorm.DB
)

// InitializeUtils initializes the application utils.
func InitializeUtils() {
	logger = configs.GetLogger("utils")
	db = configs.GetDatabase()
}
