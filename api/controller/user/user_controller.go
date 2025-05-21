package controller

import (
	middleware "easy-dictionary-server/api/middleware"
	"easy-dictionary-server/domain"
	userDomain "easy-dictionary-server/domain/user"
	internalenv "easy-dictionary-server/internalenv"
	validatorutil "easy-dictionary-server/internalenv/validator"
	"strconv"

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
		zap.S().Error(err)
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
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Email is required"})
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
		zap.S().Error(err)
		c.JSON(http.StatusForbidden, domain.ErrorResponse{Message: "User can't register with " + request.Email + " by provider " + request.Provider})
	} else {
		zap.S().Debugf("User created %s %s", request.FirstName, request.SecondName)
		c.JSON(http.StatusCreated, user)
	}
}

func (userController *UserController) Edit(c *gin.Context) {
	zap.S().Info("POST Edit")
	var request userDomain.EditUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}

	user, err := userController.UserUseCase.UpdateUser(c, request.ID, request.FirstName, request.SecondName)
	if err != nil || user == nil {
		zap.S().Error("Failed to update user with" + request.Email)
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "User can't update with " + request.Email})
	} else {
		zap.S().Debugf("User updated %s %s", request.FirstName, request.SecondName)
		c.JSON(http.StatusOK, user)
	}
}

func (userController *UserController) GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	zap.S().Infof("GET GetUserByID: %s", userID)
	if userIdInt, err := strconv.Atoi(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	} else {
		user, err := userController.UserUseCase.GetByID(c, userIdInt)
		if err != nil || user == nil {
			zap.S().Errorf("Failed to get user id %d", userIdInt)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to find user"})
			return
		} else {
			zap.S().Debugf("User found %s %s", user.FirstName, user.SecondName)
			c.JSON(http.StatusOK, user)
			return
		}
	}

}
