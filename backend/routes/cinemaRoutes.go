package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CinemaRoute(router *gin.Engine) {
	cinema := router.Group("cinema")
	{
		cinema.GET("all", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "cinema"})
		})
	}
}
