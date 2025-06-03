package router

import (
	controller "easy-dictionary-server/api/controller/user"
	middleware "easy-dictionary-server/api/middleware"
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"
	repositoryUser "easy-dictionary-server/repository/user"
	usecase "easy-dictionary-server/usecase/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewUserClientRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up user route for client")
	ur := repositoryUser.NewUserRepository(database)
	ac := &controller.UserController{
		UserUseCase: usecase.NewUserUsecase(ur, timeout),
	}
	clientGroup := group.Group("", middleware.JWTMiddleware(env, middleware.Client.VALUE))
	{
		clientGroup.POST("api/users/edit", ac.Edit)
		clientGroup.GET("api/users/:id", ac.GetUserByID)
	}
	group.POST("api/signup", func(c *gin.Context) {
		ac.Register(c, middleware.Client.VALUE)
	})
}

func NewUserAdminRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up user route for admin")
	ur := repositoryUser.NewUserRepository(database)
	ac := &controller.UserController{
		UserUseCase: usecase.NewUserUsecase(ur, timeout),
	}
	adminGroup := group.Group("", middleware.JWTMiddleware(env, middleware.Admin.VALUE))
	{
		adminGroup.GET("api/users/all", ac.GetAllUsers)
	}
}
