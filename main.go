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

	registerRoutes(r, authMiddleware)
	handlerMiddleware(authMiddleware)

	r.Run("localhost:8080")
}

func registerRoutes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	r.POST("/login", authMiddleware.LoginHandler)
	r.NoRoute(authMiddleware.MiddlewareFunc(), handleNoRoute())

	protected := r.Group("/api")
	protected.Use(authMiddleware.MiddlewareFunc())
	protected.GET("/test", func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		user, _ := c.Get("id")
		c.JSON(http.StatusOK, gin.H{
			"userID":   claims["id"],
			"userName": user.(*middleware.User).UserName,
			"message":  "Welcome to the testing route!",
		})
	})

	protected.GET("/refresh_token", authMiddleware.RefreshHandler)

	protected.GET("/post/:id", routes.GetPostById)
	protected.PUT("/post/:id", routes.UpdatePostById)
	protected.DELETE("/post/:id", routes.RemovePostById)
	protected.POST("/post", routes.CreatePostById)
}

func handlerMiddleware(authMiddleware *jwt.GinJWTMiddleware) {
	if err := authMiddleware.MiddlewareInit(); err != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + err.Error())
	}
}

func handleNoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		if claims["id"] != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "Page not found",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized access",
			})
		}
	}
}
