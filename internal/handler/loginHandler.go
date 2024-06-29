package handler

import (
	"errors"
	"net/http"
	"os"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/sophia-app/backend/internal/schemas"
	"github.com/sophia-app/backend/pkg/hash"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @BasePath /api/v1

// @Summary Authenticate user
// @Description Authenticate user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login request object"
// @Success 200 {object} nil
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /login [post]
func Login(ctx *gin.Context) {
	request := LoginRequest{}

	if err := validateJsonRequest(ctx, &request); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{}

	err := db.Where("username = ?", request.Username).Or("email = ?", request.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle case where user is not found
			logger.Infof("no user found with the given username/email")
			sendError(ctx, http.StatusNotFound, "user not found")
		} else {
			// Handle other errors
			logger.Errorf("error getting user: %v", err.Error())
			sendError(ctx, http.StatusInternalServerError, "error getting user")
		}
		return
	}

	salt := os.Getenv("HASH_SALT")
	encrypter := hash.NewArgon2idHash(1, uint32(utf8.RuneCountInString(salt)), 64*1024, 32, 256)

	if err := encrypter.Compare([]byte(user.Password), []byte(request.Password), []byte(salt)); err != nil {
		logger.Errorf("error comparing password: %v", err.Error())
		sendError(ctx, http.StatusUnauthorized, "invalid password")
		return
	}

	sendSuccess(ctx, "login successful", nil)
}
