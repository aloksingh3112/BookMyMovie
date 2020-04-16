package middleware

import (
	"net/http"

	"github.com/aloksingh3112/BookMyMovie/config"
	"github.com/aloksingh3112/BookMyMovie/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyToken(c *gin.Context) {

	header := c.Request.Header["Authorization"][0]

	claims := &types.Claims{}
	_, err := jwt.ParseWithClaims(header, claims, func(toekn *jwt.Token) (interface{}, error) {
		return config.SECRETKEY, nil
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": nil, "message": "you are not authenticated", "statusCode": 400})
		c.Abort()
		return
	}

	c.Next()

}
