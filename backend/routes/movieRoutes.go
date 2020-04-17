package routes

import (
	controller "github.com/aloksingh3112/BookMyMovie/controllers"

	"github.com/gin-gonic/gin"
)

func MovieRoutes(router *gin.Engine) {
	movie := router.Group("/movie")
	{
		movie.POST("/addMovie", controller.AddMovie)
	}
}
