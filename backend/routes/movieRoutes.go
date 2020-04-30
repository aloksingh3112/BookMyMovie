package routes

import (
	controller "github.com/aloksingh3112/BookMyMovie/controllers"
	"github.com/aloksingh3112/BookMyMovie/middleware"

	"github.com/gin-gonic/gin"
)

func MovieRoutes(router *gin.Engine) {
	movie := router.Group("/movie")
	{
		movie.POST("/addMovie", middleware.VerifyToken, middleware.VerifyUser, controller.AddMovie)
		movie.GET("/getMovies", controller.GetMovies)
		movie.PUT("/updateMovie/:id", middleware.VerifyToken, middleware.VerifyUser, controller.UpdateMovies)
		movie.DELETE("/deleteMovie/:id", middleware.VerifyToken, middleware.VerifyUser, controller.DeleteMovie)
	}
	movieTheatreMap := router.Group("/map")
	{
		movieTheatreMap.POST("/movieTheatre", middleware.VerifyToken, middleware.VerifyUser, controller.MapMovieWithTheatre)
		movieTheatreMap.POST("/seatMap", middleware.VerifyToken, middleware.VerifyUser, controller.GetSeatMapping)

	}
}
