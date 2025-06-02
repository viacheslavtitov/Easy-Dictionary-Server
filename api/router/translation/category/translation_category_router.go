package route

import (
	controller "easy-dictionary-server/api/controller/translation/category"
	middleware "easy-dictionary-server/api/middleware"
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"
	repository "easy-dictionary-server/repository/translation/category"
	usecase "easy-dictionary-server/usecase/translation/category"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewTranslationCategoryRouter(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up translation category route")
	rl := repository.NewTranslationCategoryRepository(database)
	lc := &controller.TranslationCategoryController{
		TranslationCategoryUseCase: usecase.NewTranslationCategoryUsecase(rl, timeout),
	}
	transCategoryGroup := group.Group("", middleware.JWTMiddleware(env, middleware.Client.VALUE))
	{
		transCategoryGroup.POST("api/translation/category/create", lc.Create)
		transCategoryGroup.POST("api/translation/category/edit", lc.Edit)
		transCategoryGroup.GET("api/translation/category/all", lc.GetAllForUser)
		transCategoryGroup.DELETE("api/translation/category/:id", lc.Delete)
	}
}
