package route

import (
	"easy-dictionary-server/api/controller"
	"easy-dictionary-server/repository"
	"easy-dictionary-server/usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewAuthRouter(timeout int, group *gin.RouterGroup) {
	zap.S().Info("Set up auth route")
	ur := repository.NewUserRepository()
	ac := &controller.AuthController{
		AuthUseCase: usecase.NewAuthUsecase(ur, timeout),
	}
	group.POST("/auth", ac.Login)
	zap.S().Debug("added /auth post method")
}
