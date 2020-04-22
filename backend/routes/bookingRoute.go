package routes

import (
	controller "github.com/aloksingh3112/BookMyMovie/controllers"
	"github.com/aloksingh3112/BookMyMovie/middleware"

	"github.com/gin-gonic/gin"
)

func BookingRoute(r *gin.Engine) {
	booking := r.Group("/booking")
	{
		booking.POST("/bookMovie", middleware.VerifyToken, middleware.VerifyUser, controller.BookMovie)

	}
}
