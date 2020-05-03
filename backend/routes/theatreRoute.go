package routes

import (
	controller "github.com/aloksingh3112/BookMyMovie/controllers"
	"github.com/aloksingh3112/BookMyMovie/middleware"

	"github.com/gin-gonic/gin"
)

func TheatreRoute(r *gin.Engine) {
	theatre := r.Group("/theatre")
	{
		theatre.POST("/addTheatre", middleware.VerifyToken, middleware.VerifyUser, controller.CreateTheatre)
		theatre.GET("/getTheatres", controller.GetTheatres)
		theatre.PUT("/updateTheatre/:id", middleware.VerifyToken, middleware.VerifyUser, controller.UpdateTheatre)
		theatre.DELETE("/deleteTheatre/:id", middleware.VerifyToken, middleware.VerifyUser, controller.DeleteTheatre)

	}
	theatreMap := r.Group("/map")
	{
		theatreMap.GET("/mapDate/:id", controller.GetTheatreMovieMap)
	}

}
