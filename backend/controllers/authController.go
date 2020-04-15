package controller

import (
	"fmt"
	"net/http"

	"github.com/aloksingh3112/BookMyMovie/models"
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
	dbData := db.Where("username=?", data.Username).First(&user)
	fmt.Println(dbData.Value, dbData.Error)
	if dbData.Error != nil {
		fmt.Println(dbData.Error)
		return
	}

}
