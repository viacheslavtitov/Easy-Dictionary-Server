package route

import (
	database "easy-dictionary-server/db"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Setup(timeout int, group *gin.RouterGroup, database *database.Database) {
	zap.S().Info("Set up routes with timeout sec ", timeout)
	NewAuthRouter(timeout, group, database)
}
