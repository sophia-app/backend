package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sophia-app/backend/internal/handlers"
	"github.com/sophia-app/backend/internal/handlers/auth"

	docs "github.com/sophia-app/backend/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// initializeRoutes initializes the routes.
func initializeRoutes(router *gin.Engine) {
	handlers.InitializeHandlers()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.POST("/login", auth.Login)
		v1.POST("/register", auth.Register)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
