package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TheatreRoute(r *gin.Engine) {
	theatre := r.Group("/theatre")
	{
		theatre.GET("/all", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "hello"})
		})
	}

}
