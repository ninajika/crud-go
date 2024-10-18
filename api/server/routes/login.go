package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/ninajika/crud-go/api/server/types"
	"github.com/ninajika/crud-go/api/utils"
)

func LoginHandler(c *gin.Context) {
	var loginVals types.Login

	if err := c.ShouldBind(&loginVals); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if loginVals.Username != "admin" || loginVals.Password != "password" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		c.Abort()
		return
	}

	userID := uint(1)
	token, err := utils.GenerateToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.SetCookie("jwt_token", token, 3600, "/", "localhost", false, true)

	claims, err := utils.TokenValidate(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	expiration := claims.Claims.(jwt.MapClaims)["exp"].(float64)

	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"user_id":    userID,
		"expires_at": time.Unix(int64(expiration), 0).Format(time.RFC3339), // Format the expiration time
	})
}
