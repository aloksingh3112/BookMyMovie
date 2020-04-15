package routes

import (
	controller "github.com/aloksingh3112/BookMyMovie/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoute(c *gin.Engine) {
	auth := c.Group("/auth")
	{
		auth.POST("/signup", controller.Signup)
	}
}
