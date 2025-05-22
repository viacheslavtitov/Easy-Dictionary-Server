package route

import (
	controller "easy-dictionary-server/api/controller/dictionary"
	middleware "easy-dictionary-server/api/middleware"
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"
	repositoryDictionary "easy-dictionary-server/repository/dictionary"
	usecase "easy-dictionary-server/usecase/dictionary"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewDictionaryRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up dictionary route")
	rd := repositoryDictionary.NewDictionaryRepository(database)
	dc := &controller.DictionaryController{
		DictionaryUseCase: usecase.NewDictionaryUsecase(rd, timeout),
	}
	dictGroup := group.Group("", middleware.JWTMiddleware(env, middleware.Client.VALUE))
	{
		dictGroup.POST("api/dictionary/create", dc.Create)
		dictGroup.POST("api/dictionary/edit", dc.Edit)
		dictGroup.GET("api/dictionary/all", dc.GetAllForUser)
		dictGroup.DELETE("api/dictionary/:id", dc.Delete)
	}
}
