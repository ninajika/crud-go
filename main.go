package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ninajika/crud-go/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/post/:id", controllers.GetPostById)
	r.POST("/post/:id/update", controllers.UpdatePostById)
	r.GET("/post/:id/remove", controllers.RemovePostById)
	r.POST("/post/:id/create", controllers.CreatePostById)
	r.Run("localhost:8080")
}
