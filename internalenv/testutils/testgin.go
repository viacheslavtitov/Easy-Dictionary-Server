package testutils

import (
	"context"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func GetTestGinContext() context.Context {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return c
}
