package router

import (
	controller "easy-dictionary-server/api/controller/quiz"
	middleware "easy-dictionary-server/api/middleware"
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"
	repository "easy-dictionary-server/repository/quiz"
	usecase "easy-dictionary-server/usecase/quiz"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewQuizRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up quiz route")
	rd := repository.NewQuizRepository(database)
	dc := &controller.QuizController{
		QuizUseCase: usecase.NewQuizUsecase(rd, timeout),
	}
	dictGroup := group.Group("", middleware.JWTMiddleware(env, middleware.Client.VALUE))
	{
		dictGroup.POST("api/quiz/create", dc.Create)
		dictGroup.POST("api/quiz/edit", dc.Edit)
		dictGroup.POST("api/quiz/word/add", dc.AddWord)
		dictGroup.GET("api/quiz/all", dc.GetAllQuizes)
		dictGroup.GET("api/quiz/:id", dc.GetQuizById)
		dictGroup.DELETE("api/quiz/:id", dc.Delete)
		dictGroup.DELETE("api/quiz/word/:id", dc.DeleteWord)
	}
}
