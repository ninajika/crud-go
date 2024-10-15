package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ninajika/crud-go/api/server/middleware"
	"github.com/ninajika/crud-go/api/server/routes"
)

func main() {
	r := gin.Default()

	r.GET("/post/:id", routes.GetPostById)
	r.POST("/post/:id/update", routes.UpdatePostById)
	r.GET("/post/:id/remove", routes.RemovePostById)
	r.POST("/post/:id/create", routes.CreatePostById)

	r.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.BindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Hardcoded credentials
		if loginData.Username == "admin" && loginData.Password == "password" {
			// Generate single JWT token
			tokenString, err := middleware.GenerateToken(loginData.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
				return
			}

			// Set the token in a cookie
			c.SetCookie("jwt_token", tokenString, 24*3600, "/", "localhost", false, true)

			// Respond with success message
			c.JSON(http.StatusOK, gin.H{
				"message": "Logged in successfully",
				"token":   tokenString,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		}
	})

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/protected-route", func(c *gin.Context) {
		userID, _ := c.Get("userID")
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the protected route!",
			"user":    userID,
		})
	})
	r.Run("localhost:8080")
}
