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

func NewUserRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env, role string) {
	zap.S().Info("Set up user route")
	ur := repositoryUser.NewUserRepository(database)
	ac := &controller.UserController{
		UserUseCase: usecase.NewUserUsecase(ur, timeout),
	}
	roleMiddleware := middleware.JWTMiddleware(env, role)
	group.Use(roleMiddleware)
	{
		group.GET("api/users/all", ac.GetAllUsers)
		group.POST("api/signup", func(c *gin.Context) {
			ac.Register(c, role)
		})

		group.POST("api/users/edit", ac.Edit)
		group.GET("api/users/:id", ac.GetUserByID)
	}
}
