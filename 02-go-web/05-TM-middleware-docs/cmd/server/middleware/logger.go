package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(
			c.Request.Method,
			c.Request.URL,
			time.Now().Format("02/01/2006 15:04:05"),
			c.Request.ContentLength, "bytes",
		)

		c.Next()
	}
}
