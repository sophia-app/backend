package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sophia-app/backend/internal/handler"

	docs "github.com/sophia-app/backend/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// initializeRoutes initializes the routes.
func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.POST("/login", handler.Login)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
