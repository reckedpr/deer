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

		if path == "/favicon.ico" {
			return
		}

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		logger := zap.S()

		fields := []interface{}{
			"status", status,
			"method", c.Request.Method,
			"path", path,
			"latency", latency,
			"clientIP", c.ClientIP(),
		}

		if len(c.Errors) > 0 {
			fields = append(fields, "error", c.Errors.String())
		}

		switch {
		case status >= 500:
			logger.Errorw("HTTP request", fields...)
		case status >= 400:
			logger.Warnw("HTTP request", fields...)
		default:
			logger.Infow("HTTP request", fields...)
		}
	}
}

func NoCacheHeaders(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")
}
