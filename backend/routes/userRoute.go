package routes

import (
	controller "github.com/aloksingh3112/BookMyMovie/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	cinema := router.Group("/user")
	{
		cinema.GET("/all", controller.User)
	}
}
