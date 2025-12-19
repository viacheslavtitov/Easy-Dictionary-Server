package router

import (
	controller "easy-dictionary-server/api/controller/tense"
	middleware "easy-dictionary-server/api/middleware"
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"
	repository "easy-dictionary-server/repository/dictionary"
	usecase "easy-dictionary-server/usecase/dictionary"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewTenseRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up tense route")
	rd := repository.NewTenseRepository(database)
	dc := &controller.TenseController{
		TenseUseCase: usecase.NewTenseUsecase(rd, timeout),
	}
	tenseGroup := group.Group("", middleware.JWTMiddleware(env, middleware.Client.VALUE))
	{
		tenseGroup.POST("api/tense/edit", dc.Edit)
		tenseGroup.DELETE("api/tense/:id", dc.Delete)
	}
}
