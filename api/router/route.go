package route

import (
	routeAuth "easy-dictionary-server/api/router/auth"
	routeDictionary "easy-dictionary-server/api/router/dictionary"
	routeLanguage "easy-dictionary-server/api/router/language"
	routeTranslation "easy-dictionary-server/api/router/translation"
	routeTranslationCategory "easy-dictionary-server/api/router/translation/category"
	routeUser "easy-dictionary-server/api/router/user"
	routeWord "easy-dictionary-server/api/router/word"
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Setup(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up routes with timeout sec ", timeout)
	//for admin
	routeUser.NewUserAdminRouter(timeout, group, database, env)
	//for clients
	routeAuth.NewAuthRouter(timeout, group, database, env)
	routeUser.NewUserClientRouter(timeout, group, database, env)
	routeLanguage.NewLanguageRouter(timeout, group, database, env)
	routeDictionary.NewDictionaryRouter(timeout, group, database, env)
	routeTranslationCategory.NewTranslationCategoryRouter(timeout, group, database, env)
	routeWord.NewWordRouter(timeout, group, database, env)
	routeTranslation.NewTranslationRouter(timeout, group, database, env)
}
