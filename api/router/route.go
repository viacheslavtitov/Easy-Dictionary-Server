package route

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Setup(timeout int, group *gin.RouterGroup) {
	zap.S().Info("Set up routes with timeout sec ", timeout)
	NewAuthRouter(timeout, group)
}
