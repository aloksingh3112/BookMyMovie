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

	routes.CinemaRoute(router)
	routes.MovieRoutes(router)
	router.Run()
}
