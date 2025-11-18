package router

import (
	controller "easy-dictionary-server/api/controller/word/tag"
	middleware "easy-dictionary-server/api/middleware"
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"
	repository "easy-dictionary-server/repository/word/tag"
	usecase "easy-dictionary-server/usecase/word/tag"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewWordRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up word tag route")
	rl := repository.NewWordTagRepository(database)
	lc := &controller.WordTagController{
		WordTagUseCase: usecase.NewWordTagUsecase(rl, timeout),
	}
	wordTagGroup := group.Group("", middleware.JWTMiddleware(env, middleware.Client.VALUE))
	{
		wordTagGroup.POST("api/word/tag/create", lc.Create)
		wordTagGroup.POST("api/word/tag/edit", lc.Edit)
		wordTagGroup.GET("api/word/tag/dictionary/all", lc.GetAllForDictionary)
		wordTagGroup.GET("api/word/tag/word/all", lc.GetAllForWord)
		wordTagGroup.DELETE("api/word/tag/:id", lc.Delete)
	}
}
