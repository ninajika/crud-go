package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/ninajika/crud-go/api/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("jwt_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, token not found"})
			c.Abort()
			return
		}

		token, err := utils.TokenValidate(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			userID := uint(claims["user_id"].(float64))
			exp := int64(claims["exp"].(float64))

			c.Set("userID", userID)
			c.Set("exp", time.Unix(exp, 0))
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}
