package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ninajika/crud-go/models"
	"github.com/ninajika/crud-go/utils"
)

func GetPostById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	result, err := utils.GetDump(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreatePostById(c *gin.Context) {
	id := c.Param("id")
	var data models.PostType
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := utils.CreateJson(id, &data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, data)
}

func UpdatePostById(c *gin.Context) {
	id := c.Param("id")
	var data models.PostType
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := utils.UpdateJson(id, &data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func RemovePostById(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Deleting post with ID: %s", id)
	if err := utils.DeletePost(id); err != nil {
		log.Println(err)
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}
	c.Status(204)
}
