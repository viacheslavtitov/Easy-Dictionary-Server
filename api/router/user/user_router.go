package route

import (
	controller "easy-dictionary-server/api/controller/user"
	// middleware "easy-dictionary-server/api/middleware"
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"
	"easy-dictionary-server/repository"
	usecase "easy-dictionary-server/usecase/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewUserRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up user route")
	ur := repository.NewUserRepository(database)
	ac := &controller.UserController{
		UserUseCase: usecase.NewUserUsecase(ur, timeout),
	}
	group.POST("api/signup", ac.Register)
	// group.POST("/edit", ac.Register, middleware.JWTMiddleware(env))
}
