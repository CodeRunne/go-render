package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// set mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/", DefaultPage)
	router.GET("/ping", DefaultPong)
	router.Run(":3000")
}

func DefaultPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome to the default page",
	})
}

func DefaultPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
