package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ninajika/crud-go/api/server/controllers"
	"github.com/ninajika/crud-go/api/server/types"
)

func GetPostById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	result, err := controllers.GetPost(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreatePostById(c *gin.Context) {
	var jsonData types.CreatePostInput
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal error when parsing json",
		})
		return
	}

	err := controllers.CreatePost(jsonData.ID, jsonData.Title, jsonData.Body, jsonData.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal error when creating data",
		})
		return
	}
}

func UpdatePostById(c *gin.Context) {
	id := c.Param("id")
	var jsonData types.CreatePostInput
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal error when parsing json",
		})
		return
	}

	ok, err := controllers.UpdatePost(id, jsonData.Title, jsonData.Body, jsonData.Tags)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("internal error: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("sucessfully update  post with id: %s", id),
	})
}

func RemovePostById(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Deleting post with ID: %s", id)
	if !controllers.RemovePost(id) {
		log.Println("fail to delete post")
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}
	c.Status(204)
}
