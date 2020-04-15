package main

import (
	"github.com/aloksingh3112/BookMyMovie/db"
	"github.com/aloksingh3112/BookMyMovie/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := db.Setup()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	routes.UserRoute(router)
	routes.MovieRoutes(router)
	routes.TheatreRoute(router)
	routes.BookingRoute(router)
	routes.AuthRoute(router)
	router.Run()
}
