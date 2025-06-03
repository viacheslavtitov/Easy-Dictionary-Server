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
	transCategoryGroup := group.Group("", middleware.JWTMiddleware(env, middleware.Client.VALUE))
	{
		transCategoryGroup.POST("api/word/tag/create", lc.Create)
		transCategoryGroup.POST("api/word/tag/edit", lc.Edit)
		transCategoryGroup.GET("api/word/tag/all", lc.GetAllForDictionary)
		transCategoryGroup.DELETE("api/word/tag/:id", lc.Delete)
	}
}
