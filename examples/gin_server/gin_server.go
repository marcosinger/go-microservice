package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func aboutHandle(c *gin.Context) {
	c.JSON(200, gin.H{
		"timestamp": time.Now(),
		"message":   "Golang Brasil",
	})
}

func main() {
	router := gin.Default()
	router.GET("/gin/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/gin/meetup/about", aboutHandle)
	router.Run(":8081")
}
