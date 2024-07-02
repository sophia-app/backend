package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sophia-app/backend/internal/handlers/common"
	"github.com/sophia-app/backend/internal/schemas"
	"github.com/sophia-app/backend/internal/utils"
	myjwt "github.com/sophia-app/backend/pkg/jwt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	if r.Username == "" {

		return common.ErrParamIsRequired("username", "string")
	}

	if r.Password == "" {
		return common.ErrParamIsRequired("password", "string")
	}

	return nil
}

type LoginResponse struct {
	Message string   `json:"message"`
	Data    JWTToken `json:"data"`
}

type JWTToken struct {
	Token string `json:"token"`
}

// @BasePath /api/v1

// @Summary Authenticate user
// @Description Authenticate user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login request object"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 401 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /login [post]
func Login(ctx *gin.Context) {
	request := LoginRequest{}

	if err := common.ValidateJsonRequest(ctx, &request); err != nil {
		logger.Errorf("validation error: %v", err)
		common.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err)
		common.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{}

	// Get user by username or email
	err := db.Where("username = ?", request.Username).Or("email = ?", request.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Infof("no user found with the given username/email")
			common.SendError(ctx, http.StatusNotFound, "user not found")
		} else {
			logger.Errorf("error getting user: %v", err)
			common.SendError(ctx, http.StatusInternalServerError, "error getting user")
		}
		return
	}

	encrypter := utils.GetHashEncrypter()

	// Compare password
	if err := encrypter.Compare(user.Password, utils.GetHashSalt(), request.Password); err != nil {
		logger.Errorf("error comparing password: %v", err.Error())
		common.SendError(ctx, http.StatusUnauthorized, "invalid password")
		return
	}

	payload := utils.GetJWTPayload(user)
	secret := utils.GetJWTSecret()

	// Create JWT token
	token, err := myjwt.CreateJWT(payload, secret)
	if err != nil {
		logger.Errorf("error creating jwt: %v", err.Error())
		common.SendError(ctx, http.StatusInternalServerError, "error generating token")
	}

	jwtToken := JWTToken{
		Token: token,
	}

	common.SendSuccess(ctx, "login successful", jwtToken)
}
