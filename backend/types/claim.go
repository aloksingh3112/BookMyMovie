package types

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Typeofuser string `json:"typeofuser"`
	jwt.StandardClaims
}
