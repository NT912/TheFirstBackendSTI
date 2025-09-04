package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// TODO: verify JWT token
		c.Next()
	}
}
