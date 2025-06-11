package controller

import (
	"net/http"
	"time"

	middleware "easy-dictionary-server/api/middleware"
	"easy-dictionary-server/domain"
	internalenv "easy-dictionary-server/internalenv"
	validatorutil "easy-dictionary-server/internalenv/validator"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct {
	AuthUseCase domain.AuthUseCase
	Env         *internalenv.Env
}

// SignIn godoc
// @Summary      Sign in
// @Description  Do sign in user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {string}  Barear Access Token
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/signin [post]
func (authController *AuthController) Login(c *gin.Context) {
	zap.S().Info("POST Login")
	var request domain.AuthRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}
	if request.Provider == "email" {
		if request.Password == "" {
			zap.S().Error("Password is empty")
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Password is required"})
			return
		}
		if request.Email == "" {
			zap.S().Error("Email is empty")
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Email is required"})
			return
		}
		if !middleware.ValidatePassword(request.Password) {
			zap.S().Error("Password is not valid")
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Password validation failed"})
			return
		}
	}

	user, userId, err := authController.AuthUseCase.GetUserByEmail(c, request.Email)
	if err != nil || (user == nil || userId == nil) {
		zap.S().Errorf("User with %s %s not found", request.Provider, request.Email)
		zap.S().Error(err)
		c.JSON(http.StatusForbidden, domain.ErrorResponse{Message: "User not found with the " + request.Email + " email"})
		return
	} else {
		zap.S().Debugf("User found %s %s %s", user.FirstName, user.LastName, user.UUID)
		if request.Provider == "email" {
			emailProvider := user.FindEmailProvider()
			if emailProvider == nil {
				c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "User is not registered by " + request.Email + " email"})
				return
			}
			if !middleware.ComparePassword(request.Password, emailProvider.HashedPassword) {
				c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Password validation failed"})
				return
			} else {
				zap.S().Debug("User validated by password")
			}
		} else {
			//implement in the feature
			zap.S().Debug("User validated by provider")
		}
		accessToken, err := authController.AuthUseCase.CreateAccessToken(user, authController.Env.AppEnv, authController.Env.JwtSecret, user.Role,
			time.Duration(authController.Env.JwtExpTimeMinutes)*time.Minute, *userId)
		if err != nil {
			zap.S().Errorf("Failed to create access token for user %s", user.UUID)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		authResponse := domain.AuthResponse{
			AccessToken: accessToken,
		}
		c.JSON(http.StatusOK, authResponse)
	}
}
