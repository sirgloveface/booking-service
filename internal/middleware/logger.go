package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		log.Printf("[GIN] %v | %3d | %13v | %s | %s\n",
			time.Now().Format("2006/01/02 - 15:04:05"),
			status,
			latency,
			method,
			path,
		)
	}
}
