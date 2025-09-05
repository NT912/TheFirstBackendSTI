package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lay JWT tu cookie
		tokenString, err := c.Cookie("jwt_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Status":  "Error",
				"Message": "Missing auth token",
			})
			c.Abort()
			return
		}

		// Parse and verify token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Kiem tra thuat toan ký
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			// Lay secret tu env
			secret := os.Getenv("JWT_SECRET")
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Status":  "Error",
				"Message": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Lay claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Gan vao context de controller lay
			c.Set("user_id", claims["user_id"])
			c.Set("email", claims["email"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Status":  "Error",
				"Message": "Invalid token claims",
			})
			c.Abort()
			return
		}

		// Cho request di tiep
		c.Next()
	}
}
