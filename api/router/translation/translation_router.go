package router

import (
	controller "easy-dictionary-server/api/controller/translation"
	middleware "easy-dictionary-server/api/middleware"
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"
	repository "easy-dictionary-server/repository/translation"
	usecase "easy-dictionary-server/usecase/translation"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewTranslationRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up translation route")
	rl := repository.NewTranslationRepository(database)
	lc := &controller.TranslationController{
		TranslationUseCase: usecase.NewTranslationUsecase(rl, timeout),
	}
	transGroup := group.Group("", middleware.JWTMiddleware(env, middleware.Client.VALUE))
	{
		transGroup.POST("api/translation/create", lc.Create)
		transGroup.POST("api/translation/edit", lc.Edit)
		transGroup.GET("api/translation/all/:id", lc.GetAllForWord)
		transGroup.DELETE("api/translation/:id", lc.Delete)
	}
}
