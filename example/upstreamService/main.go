package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/v1/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"posts": "Get Post"})
	})

	router.POST("/api/v1/posts", func(c *gin.Context) {
		jsonData, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println(string(jsonData))
		c.JSON(http.StatusOK, gin.H{"new post": string(jsonData)})
	})
	router.Run(":8082")
}
