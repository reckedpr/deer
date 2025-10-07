package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func ReturnError(c *gin.Context, code int, msg string, err error) {
	if err != nil {
		log.Printf("[ %d ] %s: %v", code, msg, err)
	}
	c.JSON(code, gin.H{"error": msg})
	c.Abort()
}
