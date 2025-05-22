package route

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
	roleMiddleware := middleware.JWTMiddleware(env, middleware.Client.VALUE)
	group.Use(roleMiddleware)
	{
		group.POST("api/signup", func(c *gin.Context) {
			ac.Register(c, middleware.Client.VALUE)
		})

		group.POST("api/users/edit", ac.Edit)
		group.GET("api/users/:id", ac.GetUserByID)
	}
}

func NewUserAdminRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up user route for admin")
	ur := repositoryUser.NewUserRepository(database)
	ac := &controller.UserController{
		UserUseCase: usecase.NewUserUsecase(ur, timeout),
	}
	roleMiddleware := middleware.JWTMiddleware(env, middleware.Admin.VALUE)
	group.Use(roleMiddleware)
	{
		group.GET("api/users/all", roleMiddleware, ac.GetAllUsers)
	}
}
