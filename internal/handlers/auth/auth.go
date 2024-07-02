package auth

import (
	"github.com/sophia-app/backend/configs"
	"gorm.io/gorm"
)

var (
	logger *configs.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = configs.GetLogger("authHandler")
	db = configs.GetDatabase()
}
