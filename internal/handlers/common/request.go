package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ValidateJsonRequest validates the request body of a JSON request.
func ValidateJsonRequest(ctx *gin.Context, request interface{}) error {
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return fmt.Errorf("malformed request body")
	}

	return nil
}

// ErrParamIsRequired returns an error indicating that a parameter is required.
func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}
