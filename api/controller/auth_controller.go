package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viacheslavtitov/easy-dictionary-server/domain"
	"go.uber.org/zap"
)

type AuthController struct {
	AuthUseCase domain.AuthUseCase
}

func (authController *AuthController) Login(c *gin.Context) {
	zap.S().Info("POST Login")
	var request domain.AuthRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := authController.AuthUseCase.GetUserByEmail(c, request.Email)
	if err != nil || user == nil {
		zap.S().Error("User with " + request.Email + " not found")
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
		return
	} else {
		zap.S().Debugf("User found %s %d", user.Email, user.ID)
	}

	// if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
	// 	c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
	// 	return
	// }

	accessToken, err := authController.AuthUseCase.CreateAccessToken(user, user.UID)
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
