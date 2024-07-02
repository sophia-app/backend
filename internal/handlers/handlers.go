package handlers

import "github.com/sophia-app/backend/internal/handlers/auth"

// InitializeHandlers initializes the application handlers.
func InitializeHandlers() {
	auth.InitializeHandler()
}
