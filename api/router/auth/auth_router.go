package route

import (
	controllerAuth "easy-dictionary-server/api/controller/auth"
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"
	repositoryUser "easy-dictionary-server/repository/user"
	usecase "easy-dictionary-server/usecase/auth"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewAuthRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up auth route")
	ur := repositoryUser.NewUserRepository(database)
	ac := &controllerAuth.AuthController{
		AuthUseCase: usecase.NewAuthUsecase(ur, timeout),
		Env:         env,
	}
	clientGroup := group.Group("")
	clientGroup.POST("api/signin", ac.Login)
}
