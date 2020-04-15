package types

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username   string `json:"username"`
	Email      string `json:"username"`
	Typeofuser string `json:"username"`
	jwt.StandardClaims
}
