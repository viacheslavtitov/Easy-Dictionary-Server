package controller

import (
	"net/http"
	"time"

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

func (authController *AuthController) Login(c *gin.Context) {
	zap.S().Info("POST Login")
	var request domain.AuthRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}

	user, err := authController.AuthUseCase.GetUserByEmail(c, request.Email)
	if err != nil || user == nil {
		zap.S().Error("User with " + request.Email + " not found")
		c.JSON(http.StatusForbidden, domain.ErrorResponse{Message: "User not found with the " + request.Email + " email"})
		return
	} else {
		zap.S().Debugf("User found %s", request.Email)
	}

	// if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
	// 	c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
	// 	return
	// }

	accessToken, err := authController.AuthUseCase.CreateAccessToken(user, authController.Env.AppEnv, authController.Env.JwtSecret, time.Duration(authController.Env.JwtExpTimeMinutes)*time.Minute)
	if err != nil {
		zap.S().Error("Failed to create access token for user UID " + user.UID)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	authResponse := domain.AuthResponse{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, authResponse)
}
