package routes

import (
	controller "github.com/aloksingh3112/BookMyMovie/controllers"
	"github.com/gin-gonic/gin"
)

func TimeRoute(r *gin.Engine) {
	time := r.Group("time")
	{
		time.GET("/getTime/:id", controller.GetTime)
	}
}
