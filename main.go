package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ninajika/crud-go/api/server/middleware"
	"github.com/ninajika/crud-go/api/server/routes"
)

func main() {
	r := gin.Default()

	registerRoutes(r)

	r.Run("localhost:8080")
}

func registerRoutes(r *gin.Engine) {
	r.POST("/login", routes.LoginHandler)

	r.NoRoute(handleNoRoute())

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/test", func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"userID":  userID,
			"message": "Welcome to the testing route!",
		})
	})

	protected.GET("/post/:id", routes.GetPostById)
	protected.PUT("/post/:id", routes.UpdatePostById)
	protected.DELETE("/post/:id", routes.RemovePostById)
	protected.POST("/post", routes.CreatePostById)
}

func handleNoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Page not found",
		})
	}
}
