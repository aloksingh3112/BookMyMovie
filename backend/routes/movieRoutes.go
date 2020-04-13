package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MovieRoutes(router *gin.Engine) {
	movie := router.Group("/movie")
	{
		movie.GET("/all", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "movie"})
		})
	}
}
