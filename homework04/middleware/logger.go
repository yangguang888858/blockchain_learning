package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path
		c.Next()
		latency := time.Since(now)
		status := c.Writer.Status()
		fmt.Printf("[%s] %s %d %s\n", method, path, status, latency)
	}
}
