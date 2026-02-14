package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinZapMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		if path == "/favicon.ico" {
			return
		}

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		zap.S().Infow("HTTP request",
			"status", status,
			"method", c.Request.Method,
			"path", path,
			"latency", latency.String(),
			"clientIP", c.ClientIP(),
			"error", c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)
	}
}

func NoCacheHeaders(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")
}
