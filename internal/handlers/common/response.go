package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Message string `json:"message"`
}

// SendError sends an error response.
func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
	})
}

// SendSuccess sends a success response.
func SendSuccess(ctx *gin.Context, message string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}
