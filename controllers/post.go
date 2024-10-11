package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/ninajika/crud-go/utils"
)

type PostType struct {
	ID        int          `json:"id"`
	Title     string       `json:"title"`
	Body      string       `json:"body"`
	Tags      []string     `json:"tags"`
	Reactions PostReaction `json:"reactions"`
	Views     int          `json:"views"`
	UserId    int          `json:"userId"`
}

type PostReaction struct {
	Likes   int `json:"likes"`
	Disikes int `json:"dislikes"`
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	result, err := utils.ReadJson[PostType](fmt.Sprintf("dummies/%s/post.json", id))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreatePostById(c *gin.Context) {
	id := c.Param("id")
	var data PostType
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	dirPath := fmt.Sprintf("dummies/%s", id)

	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	filePath := filepath.Join(dirPath, "post.json")
	if _, err := os.Stat(filePath); err == nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("post.json already exists for ID: %s", id)})
		return
	}

	if err := utils.WriteJson(filePath, &data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, data)
}

func UpdatePostById(c *gin.Context) {
	id := c.Param("id")
	var data PostType
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	dirPath := fmt.Sprintf("dummies/%s", id)
	filePath := filepath.Join(dirPath, "post.json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(400, gin.H{"error": fmt.Sprintf("post.json does not exist for ID: %s", id)})
		return
	}

	if err := utils.WriteJson(filePath, &data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func RemovePostById(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Deleting post with ID: %s", id)
	err := os.Remove(fmt.Sprintf("dummies/%s/post.json", id))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}
