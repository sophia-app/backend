package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sophia-app/backend/internal/handlers/common"
	"github.com/sophia-app/backend/internal/schemas"
	"github.com/sophia-app/backend/internal/utils"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *RegisterRequest) Validate() error {
	if r.Name == "" {
		return common.ErrParamIsRequired("name", "string")
	}

	if r.Email == "" {
		return common.ErrParamIsRequired("email", "string")
	}

	if r.Username == "" {
		return common.ErrParamIsRequired("username", "string")
	}

	if r.Password == "" {
		return common.ErrParamIsRequired("password", "string")
	}

	return nil
}

type RegisterResponse struct {
	Message string               `json:"message"`
	Data    schemas.UserResponse `json:"data"`
}

// @BasePath /api/v1

// @Summary Register new user
// @Description Register new user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Register request object"
// @Success 200 {object} RegisterResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /register [post]
func Register(ctx *gin.Context) {
	request := RegisterRequest{}

	if err := common.ValidateJsonRequest(ctx, &request); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		common.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		common.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{}

	// Check if email or username already exists
	err := db.Where("email = ?", request.Email).Or("username = ?", request.Username).First(&user).Error
	if err == nil {
		if request.Email == user.Email {
			logger.Errorf("email already exists")
			common.SendError(ctx, http.StatusBadRequest, "email already exists")
			return
		} else {
			logger.Errorf("username already exists")
			common.SendError(ctx, http.StatusBadRequest, "username already exists")
			return
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Errorf("failed to check email or username: %v", err.Error())
		common.SendError(ctx, http.StatusInternalServerError, "failed to check email or username")
		return
	}

	encrypter := utils.GetHashEncrypter()

	// Generate password hash
	hash, err := encrypter.GenerateHash(request.Password, utils.GetHashSalt())
	if err != nil {
		logger.Errorf("error generating password hash: %v", err.Error())
		common.SendError(ctx, http.StatusInternalServerError, "error generating password hash")
		return
	}

	user = schemas.User{
		Name:     request.Name,
		Email:    request.Email,
		Username: request.Username,
		Password: hash.Hash,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("failed to create user: %v", err.Error())
		common.SendError(ctx, http.StatusInternalServerError, "failed to create user")
		return
	}

	common.SendSuccess(ctx, "user created successfully", user)
}
