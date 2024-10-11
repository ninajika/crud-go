package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ninajika/crud-go/utils"
)

func main() {
	r := gin.Default()
	r.GET("/readpost/:id", func(c *gin.Context) {
		data, err := utils.GetDump(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Post %s not found", c.Param("id"))})
			return
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run("localhost:8080")
}
