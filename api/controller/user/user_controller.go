package controller

import (
	controllerCommon "easy-dictionary-server/api/controller"
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

// Register godoc
// @Summary Register user
// @Description Create new user with email and password or choice provider
// @Tags users
// @Accept  json
// @Produce  json
// @Param   input body userDomain.RegisterUserRequest true "User data"
// @Success 201 {object} userDomain.User
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/signup [post]
func (userController *UserController) Register(c *gin.Context, role string) {
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

	user, err := userController.UserUseCase.RegisterUser(c, request.FirstName, request.LastName, role, request.Email, request.Provider, passwordHash, request.ProviderToken)
	if err != nil || user == nil {
		zap.S().Error("Failed to register user with" + request.Email + " by provider " + request.Provider)
		zap.S().Error(err)
		c.JSON(http.StatusForbidden, domain.ErrorResponse{Message: "User can't register with " + request.Email + " by provider " + request.Provider})
	} else {
		zap.S().Debugf("User created %s %s", request.FirstName, request.LastName)
		c.JSON(http.StatusCreated, user)
	}
}

// Edit godoc
// @Summary      Edit user
// @Description  Update user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param input body userDomain.EditUserRequest true "New data for user"
// @Success      200  {object}  userDomain.User
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/users/edit [post]
func (userController *UserController) Edit(c *gin.Context) {
	zap.S().Info("POST Edit")
	if userId, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		var request userDomain.EditUserRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			zap.S().Error(err)
			validationErrors := validatorutil.FormatValidationError(err)
			c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
			return
		}

		user, err := userController.UserUseCase.UpdateUser(c, *userId, request.UUID, request.FirstName, request.LastName)
		if err != nil || user == nil {
			zap.S().Error("Failed to update user with" + request.Email)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "User can't update with " + request.Email})
		} else {
			zap.S().Debugf("User updated %s %s", request.FirstName, request.LastName)
			c.JSON(http.StatusOK, user)
		}
	}
}

// GetUserByUUID godoc
// @Summary      Get user by id (uuid)
// @Description  Get user by id (uuid)
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id path string true "UUID user"
// @Success      200  {object}  userDomain.User
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/users/:id [get]
func (userController *UserController) GetUserByUUID(c *gin.Context) {
	userUUID := c.Param("id")
	zap.S().Infof("GET GetUserByUUID: %s", userUUID)
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	user, err := userController.UserUseCase.GetByUUID(c, userUUID)
	if err != nil || user == nil {
		zap.S().Errorf("Failed to get user by uuid %s", userUUID)
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to find user"})
		return
	} else {
		zap.S().Debugf("User found %s %s", user.FirstName, user.LastName)
		c.JSON(http.StatusOK, user)
		return
	}
}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Get all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}  userDomain.User
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/users/all [get]
func (userController *UserController) GetAllUsers(c *gin.Context) {
	zap.S().Infof("GET GetAllUsers")
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	users, err := userController.UserUseCase.GetAllUsers(c)
	if err != nil || users == nil {
		zap.S().Errorf("Failed to get users")
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to find users"})
		return
	} else {
		zap.S().Debugf("User found %d", len(users))
		c.JSON(http.StatusOK, users)
		return
	}
}
