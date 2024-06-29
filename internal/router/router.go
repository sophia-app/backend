package router

import "github.com/gin-gonic/gin"

// Initialize initializes the router.
func Initialize() {
	router := gin.Default()

	initializeRoutes(router)

	router.Run(":8080")
}
