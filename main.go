package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ninajika/crud-go/api/server/routes"
)

func main() {
	r := gin.Default()

	r.GET("/post/:id", routes.GetPostById)
	r.POST("/post/:id/update", routes.UpdatePostById)
	r.GET("/post/:id/remove", routes.RemovePostById)
	r.POST("/post/:id/create", routes.CreatePostById)
	r.Run("localhost:8080")
}
