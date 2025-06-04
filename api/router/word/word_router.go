package router

import (
	controller "easy-dictionary-server/api/controller/word"
	middleware "easy-dictionary-server/api/middleware"
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"
	repository "easy-dictionary-server/repository/word"
	usecase "easy-dictionary-server/usecase/word"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewWordRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up word route")
	rl := repository.NewWordRepository(database)
	lc := &controller.WordController{
		WordUseCase: usecase.NewWordUsecase(rl, timeout),
	}
	transCategoryGroup := group.Group("", middleware.JWTMiddleware(env, middleware.Client.VALUE))
	{
		transCategoryGroup.POST("api/word/create", lc.Create)
		transCategoryGroup.POST("api/word/edit", lc.Edit)
		transCategoryGroup.GET("api/word/all", lc.GetAllForDictionary)
		transCategoryGroup.GET("api/word/search", lc.SearchForDictionary)
		transCategoryGroup.DELETE("api/word/:id", lc.Delete)
	}
}
