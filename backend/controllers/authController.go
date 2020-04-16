package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aloksingh3112/BookMyMovie/config"

	"github.com/aloksingh3112/BookMyMovie/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/aloksingh3112/BookMyMovie/types"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func genHashSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

func compareHashPassword(hashPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashPassword, password)
}

func genToken(username string, email string, typeofuser string) (string, error) {
	expirationTime := time.Now().Add(2000 * time.Minute)
	claim := types.Claims{
		Username:   username,
		Email:      email,
		Typeofuser: typeofuser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(config.SECRETKEY)

	if err != nil {
		fmt.Println("Some thing went wrong", err)
		return "", err
	}
	return tokenString, err

}

func Signup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var data types.User
	err := c.ShouldBindJSON(&data)
	var user models.User
	if err != nil {
		fmt.Println(err.Error())
	}

	error := db.Where("username =? AND email=?", data.Username, data.Email).First(&user).Error
	if error == nil {
		fmt.Println(error)
		return
	}

	b := []byte(data.Password)
	password := genHashSalt(b)
	fmt.Println(password, data.Username)
	userData := models.User{Username: data.Username, Password: password, TypeOfUser: data.TypeOfUser, Email: data.Email, Name: data.Name}
	db.Create(&userData)
	c.JSON(http.StatusOK, gin.H{"data": userData})

}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var data types.User
	var user models.User
	err := c.ShouldBindJSON(&data)

	if err != nil {
		fmt.Println(err)
		return
	}
	error := db.Where("username=?", data.Username).First(&user).Error

	if error != nil {
		fmt.Println(error)
		return
	}

	isValidPass := compareHashPassword([]byte(user.Password), []byte(data.Password))

	if isValidPass != nil {
		fmt.Println("Password is wrong")
		return
	}

	token, _ := genToken(user.Username, user.Email, user.TypeOfUser)
	resData := map[string]string{
		"username":   user.Username,
		"typeofuser": user.TypeOfUser,
		"token":      token,
	}
	c.JSON(http.StatusOK, gin.H{"data": resData, "statusCode": 200, "message": "user is login successfully"})
}
