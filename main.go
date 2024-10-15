package main

import (
	"log"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/ninajika/crud-go/api/server/middleware"
	"github.com/ninajika/crud-go/api/server/routes"
)

func main() {
	r := gin.Default()

	authMiddleware, err := middleware.InitJWTMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.GET("/post/:id", routes.GetPostById)
	r.POST("/post/:id/update", routes.UpdatePostById)
	r.GET("/post/:id/remove", routes.RemovePostById)
	r.POST("/post/:id/create", routes.CreatePostById)

	handlerMiddleware(authMiddleware)
	registerRoutes(r, authMiddleware)

	r.Run("localhost:8080")
}

func registerRoutes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	// Unprotected routes
	r.POST("/login", authMiddleware.LoginHandler)
	r.NoRoute(authMiddleware.MiddlewareFunc(), handleNoRoute())

	// Protected routes
	protected := r.Group("/api")
	protected.Use(authMiddleware.MiddlewareFunc())
	protected.GET("/protected-route", func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		user, _ := c.Get("id")
		c.JSON(http.StatusOK, gin.H{
			"userID":   claims["id"],
			"userName": user.(*middleware.User).UserName,
			"message":  "Welcome to the protected route!",
		})
	})

	protected.GET("/refresh_token", authMiddleware.RefreshHandler)
}

func handlerMiddleware(authMiddleware *jwt.GinJWTMiddleware) {
	if err := authMiddleware.MiddlewareInit(); err != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + err.Error())
	}
}

func handleNoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Page not found",
		})
	}
}
