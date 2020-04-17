package utils

import (
	"github.com/aloksingh3112/BookMyMovie/config"
	"github.com/aloksingh3112/BookMyMovie/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func DecodeToken(c *gin.Context) types.Claims {
	header := c.Request.Header["Authorization"][0]
	tkn, _ := jwt.ParseWithClaims(header, &types.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.SECRETKEY, nil
	})

	token, _ := tkn.Claims.(*types.Claims)
	return *token
}
