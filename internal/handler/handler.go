package handler

import (
	"github.com/sophia-app/backend/configs"
	"gorm.io/gorm"
)

var (
	logger *configs.Logger
	db     *gorm.DB
)

// InitializeHandler initializes the handler
func InitializeHandler() {
	logger = configs.GetLogger("handler")
	db = configs.GetDatabase()
}
