package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "gg"})
	})

	route.Run()
}
