package main

import (
	"github.com/aloksingh3112/BookMyMovie/db"
	"github.com/aloksingh3112/BookMyMovie/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := db.Setup()
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3002"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))
	router.Use(CORS())
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

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
