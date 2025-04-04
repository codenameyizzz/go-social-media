package middleware

import (
	"go-socialmedia/controllers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header kosong"})
			c.Abort()
			return
		}

		// Format: Bearer <token>
		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Format token salah"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString[1], &controllers.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("secretkey"), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token parsing gagal"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*controllers.Claims); ok && token.Valid {
			c.Set("userID", claims.UserID)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		c.Next()
	}
}
