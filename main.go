package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//web server setup
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"hello": "world"})
    })

	router.GET("/welcome", func(c *gin.Context) {
        c.JSON(http.StatusOK, "Here is the welcome page of the app!")
		})

	router.Run(":8080")
}
