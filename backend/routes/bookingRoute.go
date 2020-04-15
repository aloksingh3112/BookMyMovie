package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BookingRoute(r *gin.Engine) {
	booking := r.Group("/booking")
	{
		booking.GET("/all", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "booking"})
		})

	}
}
