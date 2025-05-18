package route

import (
	"easy-dictionary-server/api/controller"
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"
	"easy-dictionary-server/repository"
	"easy-dictionary-server/usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewAuthRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up auth route")
	ur := repository.NewUserRepository(database)
	ac := &controller.AuthController{
		AuthUseCase: usecase.NewAuthUsecase(ur, timeout),
	}
	group.POST("/auth", ac.Login)
}
