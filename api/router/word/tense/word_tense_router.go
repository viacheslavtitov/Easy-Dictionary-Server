package router

import (
	controller "easy-dictionary-server/api/controller/word/tense"
	middleware "easy-dictionary-server/api/middleware"
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"
	repository "easy-dictionary-server/repository/word/tense"
	usecase "easy-dictionary-server/usecase/word/tense"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewWordRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up word tag route")
	rl := repository.NewWordTagRepository(database)
	lc := &controller.WordTenseController{
		WordTenseUseCase: usecase.NewWordTenseUsecase(rl, timeout),
	}
	wordTagGroup := group.Group("", middleware.JWTMiddleware(env, middleware.Client.VALUE))
	{
		wordTagGroup.POST("api/word/tense/create", lc.Create)
		wordTagGroup.POST("api/word/tense/edit", lc.Edit)
		wordTagGroup.GET("api/word/tense/all", lc.GetAllForWord)
		wordTagGroup.DELETE("api/word/tense/:id", lc.Delete)
	}
}
