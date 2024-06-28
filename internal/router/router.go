package router

import "github.com/gin-gonic/gin"

// Initialize initializes the router
func Initialize() {
	r := gin.Default()

	initializeRoutes(r)

	r.Run(":8080")
}
