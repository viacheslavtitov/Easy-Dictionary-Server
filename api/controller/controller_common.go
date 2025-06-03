package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ValidateUserIdInContext(context *gin.Context) (*int, bool) {
	userID, exists := context.Get("userID")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return nil, false
	}
	idStr, ok := userID.(string)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return nil, false
	}
	if idInt, err := strconv.Atoi(idStr); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return nil, false
	} else {
		return &idInt, true
	}
}

func ValidateDeleteByIdResult(context *gin.Context, id string, errorMessage string, deletedRows int64, err error) bool {
	if err != nil || deletedRows < 1 {
		zap.S().Error(errorMessage + " " + id)
		zap.S().Error(err)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			return false
		} else {
			context.JSON(http.StatusInternalServerError, errorMessage+" "+id)
			return false
		}
	}
	return true
}
