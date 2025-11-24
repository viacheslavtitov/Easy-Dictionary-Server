package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ValidateUserIdInContext(context *gin.Context) (*int, *string, bool) {
	userUUID, existsUUID := context.Get("userUUID")
	userID, existsID := context.Get("userID")
	zap.S().Debugf("Get data from context user id %d, user uuid %s", userID, userUUID)
	if !existsUUID || !existsID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return nil, nil, false
	}
	uuidStr, okUUID := userUUID.(string)
	if !okUUID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user UUID"})
		return nil, nil, false
	}
	idInt, okID := userID.(int)
	if !okID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return nil, nil, false
	} else {
		return &idInt, &uuidStr, true
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

func ParseQueryInt(c *gin.Context, param string) (int, error) {
	val := c.Query(param)
	return strconv.Atoi(val)
}

func ParseQueryIntArray(c *gin.Context, param string) (*[]int, error) {
	var ids []int
	raw := c.Query(param)
	zap.S().Infow("raw for params is: ", param, raw)
	if raw != "" {
		for _, s := range strings.Split(raw, ",") {
			v, err := strconv.Atoi(strings.TrimSpace(s))
			if err != nil {
				return nil, fmt.Errorf("Param %s is not int", param)
			}
			ids = append(ids, v)
		}
	}
	return &ids, nil
}

func ParseDateTime(c *gin.Context, param string) (*time.Time, error) {
	fromStr := c.Query(param)
	zap.S().Infow("fromStr for params is: ", param, fromStr)
	if fromStr != "" {
		t, err := time.Parse(time.DateOnly, fromStr)
		if err != nil {
			return nil, fmt.Errorf("Param %s is not time or wrong format %s", param, time.DateOnly)
		}
		return &t, nil
	}
	return nil, nil
}
