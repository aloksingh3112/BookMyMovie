package routes

import (
	controller "github.com/aloksingh3112/BookMyMovie/controllers"
	"github.com/aloksingh3112/BookMyMovie/middleware"

	"github.com/gin-gonic/gin"
)

func TheatreRoute(r *gin.Engine) {
	theatre := r.Group("/theatre")
	{
		theatre.POST("/addtheatre", middleware.VerifyToken, middleware.VerifyUser, controller.CreateTheatre)

	}

}
