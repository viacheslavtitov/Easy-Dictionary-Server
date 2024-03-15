package route

import (
	"github.com/gin-gonic/gin"
	"github.com/viacheslavtitov/easy-dictionary-server/api/controller"
	"github.com/viacheslavtitov/easy-dictionary-server/repository"
	"github.com/viacheslavtitov/easy-dictionary-server/usecase"
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
