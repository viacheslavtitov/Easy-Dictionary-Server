package controller

import (
	"context"
	"net/http"
	"time"

	middleware "easy-dictionary-server/api/middleware"
	domainAuth "easy-dictionary-server/domain"
	domainUser "easy-dictionary-server/domain/user"
	internalenv "easy-dictionary-server/internalenv"
	validatorutil "easy-dictionary-server/internalenv/validator"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct {
	AuthUseCase domainAuth.AuthUseCase
	UserUseCase domainUser.UserUseCase
	Env         *internalenv.Env
}

// SignIn godoc
// @Summary      Sign in
// @Description  Do sign in user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param   input body domainAuth.AuthRequest true "User data"
// @Success      200  {string}  Barear Access Token
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/signin [post]
func (authController *AuthController) Login(c *gin.Context) {
	zap.S().Info("POST Login")
	var request domainAuth.AuthRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}
	if request.Provider == "email" {
		if request.Password == "" {
			zap.S().Error("Password is empty")
			c.JSON(http.StatusBadRequest, domainAuth.ErrorResponse{Message: "Password is required"})
			return
		}
		if request.Email == "" {
			zap.S().Error("Email is empty")
			c.JSON(http.StatusBadRequest, domainAuth.ErrorResponse{Message: "Email is required"})
			return
		}
		if !middleware.ValidatePassword(request.Password) {
			zap.S().Error("Password is not valid")
			c.JSON(http.StatusBadRequest, domainAuth.ErrorResponse{Message: "Password validation failed"})
			return
		}
	}

	user, userId, err := authController.AuthUseCase.GetUserByEmail(c, request.Email)
	if err != nil || (user == nil || userId == nil) {
		zap.S().Errorf("User with %s %s not found", request.Provider, request.Email)
		zap.S().Error(err)
		c.JSON(http.StatusForbidden, domainAuth.ErrorResponse{Message: "User not found with the " + request.Email + " email"})
		return
	} else {
		zap.S().Debugf("User found %s %s %s", user.FirstName, user.LastName, user.UUID)
		if request.Provider == "email" {
			emailProvider := user.FindEmailProvider()
			if emailProvider == nil {
				c.JSON(http.StatusBadRequest, domainAuth.ErrorResponse{Message: "User is not registered by " + request.Email + " email"})
				return
			}
			if !middleware.ComparePassword(request.Password, emailProvider.HashedPassword) {
				c.JSON(http.StatusBadRequest, domainAuth.ErrorResponse{Message: "Password validation failed"})
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
			c.JSON(http.StatusInternalServerError, domainAuth.ErrorResponse{Message: err.Error()})
			return
		}

		refreshToken, err := getOrCreateRefreshToken(c, *user, authController.UserUseCase, authController.AuthUseCase, *authController.Env)
		if err != nil {
			zap.S().Errorf("Failed to create refresh token for user %s", user.UUID)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, domainAuth.ErrorResponse{Message: err.Error()})
			return
		}
		authResponse := domainAuth.AuthResponse{
			AccessToken:     accessToken,
			RefreshToken:    refreshToken.Token,
			RefreshTokenExp: refreshToken.ExpiresAt,
		}
		c.JSON(http.StatusOK, authResponse)
	}
}

func getOrCreateRefreshToken(context context.Context, user domainUser.User, userUseCase domainUser.UserUseCase, authUseCase domainAuth.AuthUseCase, env internalenv.Env) (*domainUser.RefreshToken, error) {
	zap.S().Debug("Try to get or create refresh token")
	refreshTokenModel, err := userUseCase.GetRefreshTokenByUserUUID(context, user.UUID)
	if err != nil {
		zap.S().Errorf("Failed to get refresh token for user %s", user.UUID)
		zap.S().Error(err)
		return nil, err
	} else if refreshTokenModel == nil {
		refreshToken, expiresAt, createdAt, err := authUseCase.CreateRefreshToken(context, user.UUID, time.Duration(env.RefreshJwtExpTimeMinutes)*time.Minute)
		if err != nil {
			zap.S().Errorf("Failed to create refresh token for user %s", user.UUID)
			zap.S().Error(err)
			return nil, err
		}
		return &domainUser.RefreshToken{
			Token:     refreshToken,
			ExpiresAt: expiresAt,
			CreatedAt: *createdAt,
		}, nil
	} else if refreshTokenModel.ExpiresAt.Before(time.Now()) {
		zap.S().Debugf("Refresh token is expired %s but now %s", refreshTokenModel.ExpiresAt, time.Now())
		rows, err := userUseCase.DeleteRefreshToken(context, refreshTokenModel.UserUUID)
		if err != nil {
			zap.S().Errorf("User %s doesn't have any refresh tokens", refreshTokenModel.UserUUID)
			zap.S().Error(err)
		}
		zap.S().Debugf("Deleted %d tokens", rows)

		refreshToken, expiresAt, createdAt, err := authUseCase.CreateRefreshToken(context, user.UUID, time.Duration(env.RefreshJwtExpTimeMinutes)*time.Minute)
		if err != nil {
			zap.S().Errorf("Failed to create refresh token for user %s", user.UUID)
			zap.S().Error(err)
			return nil, err
		}
		return &domainUser.RefreshToken{
			Token:     refreshToken,
			ExpiresAt: expiresAt,
			CreatedAt: *createdAt,
		}, nil
	}
	return refreshTokenModel, nil
}

// SignIn godoc
// @Summary      Refresh Token
// @Description  Do refresh access token by refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param   input body domainAuth.RefreshTokenRequest true "Refresh token"
// @Success      200  {string}  Barear Access Token
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/refresh [post]
func (authController *AuthController) RefreshToken(c *gin.Context) {
	zap.S().Info("POST RefreshToken")
	var request domainAuth.RefreshTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}

	refreshTokenModel, err := authController.UserUseCase.GetRefreshToken(c, request.RefreshToken)
	if err != nil {
		zap.S().Errorf("Refresh token %s not found", request.RefreshToken)
		zap.S().Error(err)
		c.JSON(http.StatusForbidden, domainAuth.ErrorResponse{Message: "Refresh token " + request.RefreshToken + " not found"})
		return
	} else {
		zap.S().Debugf("Refresh token found %s", refreshTokenModel.Token)
		if refreshTokenModel.ExpiresAt.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, domainAuth.ErrorResponse{Message: "Invalid or expired refresh token"})
			return
		}
		user, userId, err := authController.UserUseCase.GetByUUID(c, refreshTokenModel.UserUUID)
		if err != nil {
			zap.S().Errorf("User %s not found", refreshTokenModel.UserUUID)
			zap.S().Error(err)
			c.JSON(http.StatusForbidden, domainAuth.ErrorResponse{Message: "User " + refreshTokenModel.UserUUID + " not found"})
			return
		}
		accessToken, err := authController.AuthUseCase.CreateAccessToken(user, authController.Env.AppEnv, authController.Env.JwtSecret, user.Role,
			time.Duration(authController.Env.JwtExpTimeMinutes)*time.Minute, *userId)
		if err != nil {
			zap.S().Errorf("Failed to create access token for user %s", user.UUID)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, domainAuth.ErrorResponse{Message: err.Error()})
			return
		}
		rows, err := authController.UserUseCase.DeleteRefreshToken(c, refreshTokenModel.UserUUID)
		if err != nil {
			zap.S().Errorf("User %s doesn't have any refresh tokens", refreshTokenModel.UserUUID)
			zap.S().Error(err)
			c.JSON(http.StatusForbidden, domainAuth.ErrorResponse{Message: "User " + refreshTokenModel.UserUUID + " not found"})
			return
		}
		zap.S().Debugf("Deleted %d tokens", rows)

		refreshToken, expiresAt, _, err := authController.AuthUseCase.CreateRefreshToken(c, user.UUID, time.Duration(authController.Env.RefreshJwtExpTimeMinutes)*time.Minute)
		if err != nil {
			zap.S().Errorf("Failed to create refresh token for user %s", user.UUID)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, domainAuth.ErrorResponse{Message: err.Error()})
			return
		}
		authResponse := domainAuth.AuthResponse{
			AccessToken:     accessToken,
			RefreshToken:    refreshToken,
			RefreshTokenExp: expiresAt,
		}
		c.JSON(http.StatusOK, authResponse)
	}
}
