package routes

import (
	controller "github.com/aloksingh3112/BookMyMovie/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", controller.Signup)
		auth.POST("/login", controller.Login)
	}
}
