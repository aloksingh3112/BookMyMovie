package middleware

import (
	"fmt"
	"net/http"

	"github.com/aloksingh3112/BookMyMovie/models"

	"github.com/aloksingh3112/BookMyMovie/config"
	"github.com/aloksingh3112/BookMyMovie/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func VerifyUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	header := c.Request.Header["Authorization"][0]
	var user models.User
	tkn, _ := jwt.ParseWithClaims(header, &types.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.SECRETKEY, nil
	})

	token, _ := tkn.Claims.(*types.Claims)
	err := db.Where("username=?", token.Username).First(&user).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": nil, "message": "No such username exist", "statusCode": 404})
		c.Abort()
		return
	}

	if user.TypeOfUser != "ADMIN" {
		fmt.Println(user.TypeOfUser, token.Username)
		c.JSON(http.StatusOK, gin.H{"data": nil, "message": "You are not Authorized", "statusCode": 403})
		c.Abort()
		return
	}

	c.Next()
}
