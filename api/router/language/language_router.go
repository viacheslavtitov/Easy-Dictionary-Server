package route

import (
	controller "easy-dictionary-server/api/controller/language"
	middleware "easy-dictionary-server/api/middleware"
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"
	repositoryLanguage "easy-dictionary-server/repository/language"
	usecase "easy-dictionary-server/usecase/language"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewLanguageRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up language route")
	rl := repositoryLanguage.NewLanguageRepository(database)
	lc := &controller.LanguageController{
		LanguageUseCase: usecase.NewLanguageUsecase(rl, timeout),
	}
	group.POST("api/languages/create", lc.Create, middleware.JWTMiddleware(env))
	group.POST("api/languages/edit", lc.Edit, middleware.JWTMiddleware(env))
	group.GET("api/languages/all", lc.GetAllForUser, middleware.JWTMiddleware(env))
	group.DELETE("api/languages/:id", lc.Delete, middleware.JWTMiddleware(env))
}
