package controller

import (
	middleware "easy-dictionary-server/api/middleware"
	"easy-dictionary-server/domain"
	userDomain "easy-dictionary-server/domain/user"
	internalenv "easy-dictionary-server/internalenv"
	validatorutil "easy-dictionary-server/internalenv/validator"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	UserUseCase userDomain.UserUseCase
	Env         *internalenv.Env
}

func (userController *UserController) Register(c *gin.Context) {
	zap.S().Info("POST Register")
	var request userDomain.RegisterUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}
	var passwordHash string
	if request.Provider == "email" {
		if request.Password == "" {
			zap.S().Error("Password is empty")
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Password is required"})
			return
		}
		if request.Email == "" {
			zap.S().Error("Email is empty")
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Provider token is required"})
			return
		}
		if !middleware.ValidatePassword(request.Password) {
			zap.S().Error("Password is not valid")
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Password validation failed"})
			return
		}
		pHash, err := middleware.GeneratePasswordHash(request.Password)
		if err != nil {
			zap.S().Error("Failed to generate password hash")
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Password validation failed"})
			return
		} else {
			passwordHash = pHash
		}
	}

	user, err := userController.UserUseCase.RegisterUser(c, request.FirstName, request.SecondName, request.Email, request.Provider, passwordHash, request.ProviderToken)
	if err != nil || user == nil {
		zap.S().Error("Failed to register user with" + request.Email + " by provider " + request.Provider)
		c.JSON(http.StatusForbidden, domain.ErrorResponse{Message: "User can't register with " + request.Email + " by provider " + request.Provider})
		return
	} else {
		zap.S().Debugf("User created %s %s", request.FirstName, request.SecondName)
		c.JSON(http.StatusCreated, user)
		return
	}
}
