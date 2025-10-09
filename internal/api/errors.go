package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ReturnError(c *gin.Context, code int, msg string, err error) {
	zap.S().Errorw("API Error",
		"path", c.FullPath(),
		"method", c.Request.Method,
		"status", code,
		"error", err,
		"msg", msg,
	)

	c.JSON(code, gin.H{"error": msg})
	c.Abort()
}
