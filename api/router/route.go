package route

import (
	database "easy-dictionary-server/db"
	internalenv "easy-dictionary-server/internalenv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Setup(timeout int, group *gin.RouterGroup, database *database.Database, env *internalenv.Env) {
	zap.S().Info("Set up routes with timeout sec ", timeout)
	NewAuthRouter(timeout, group, database, env)
}
