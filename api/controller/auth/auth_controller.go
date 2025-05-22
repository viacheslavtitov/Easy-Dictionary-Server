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

	user, err := authController.AuthUseCase.GetUserByEmail(c, request.Email)
	if err != nil || user == nil {
		zap.S().Errorf("User with %s %s not found", request.Provider, request.Email)
		zap.S().Error(err)
		c.JSON(http.StatusForbidden, domain.ErrorResponse{Message: "User not found with the " + request.Email + " email"})
		return
	} else {
		zap.S().Debugf("User found %s %s %d", user.FirstName, user.LastName, user.ID)
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
	}
	accessToken, err := authController.AuthUseCase.CreateAccessToken(user, authController.Env.AppEnv, authController.Env.JwtSecret, user.Role, time.Duration(authController.Env.JwtExpTimeMinutes)*time.Minute)
	if err != nil {
		zap.S().Errorf("Failed to create access token for user %d", user.ID)
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	authResponse := domain.AuthResponse{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, authResponse)
}
