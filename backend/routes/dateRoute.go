package routes

import (
	controller "github.com/aloksingh3112/BookMyMovie/controllers"
	"github.com/gin-gonic/gin"
)

func DateRoute(r *gin.Engine) {
	date := r.Group("date")
	{
		date.GET("/getDate/:id", controller.GetDate)
	}
}
