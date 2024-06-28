package middleware

import (
	"github.com/gin-gonic/gin"
)

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		status := c.Writer.Status()
		if status == 404 {
			c.HTML(404, "404.html", gin.H{})
			return
		}
	}
}
