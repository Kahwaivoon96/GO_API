package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware for Gin framework
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if strings.HasPrefix(c.Request.URL.Path, "/swagger/") {
			c.Next()
			return
		}

		apiKey := c.GetHeader("X-API-KEY")
		if apiKey != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
