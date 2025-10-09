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
			"clientIP", GetRealIP(c),
			"error", c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)
	}
}

func GetRealIP(c *gin.Context) string {
	if ip := c.GetHeader("CF-Connecting-IP"); ip != "" {
		return ip
	}

	if ip := c.GetHeader("X-Forwarded-For"); ip != "" {
		return ip
	}

	return c.ClientIP()
}
