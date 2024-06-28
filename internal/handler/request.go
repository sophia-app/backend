package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// validateJsonRequest validates the request body of a JSON request
func validateJsonRequest(ctx *gin.Context, request interface{}) error {
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return fmt.Errorf("malformed request body")
	}

	return nil
}

// errParamIsRequired returns an error indicating that a parameter is required
func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}
