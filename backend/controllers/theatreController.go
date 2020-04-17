package controller

import (
	"net/http"

	"github.com/aloksingh3112/BookMyMovie/models"
	"github.com/aloksingh3112/BookMyMovie/types"
	"github.com/aloksingh3112/BookMyMovie/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateTheatre(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userInput := utils.DecodeToken(c)
	var inputTheatre types.Theatre
	var user models.User

	c.ShouldBindJSON(&inputTheatre)
	db.Where("username=?", userInput.Username).First(&user)
	// db.Model(&user).AddForeignKey("user_id", "user(id)", "CASCADE", "RESTRICT")
	theatre := models.Theatre{
		Location:    inputTheatre.Location,
		City:        inputTheatre.City,
		TheatreName: inputTheatre.TheatreName,
		UserID:      user.ID,
	}
	err := db.Save(&theatre).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Something went wrong", "statusCode": 400})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": theatre, "message": "Theatre created successfully", "statusCode": 200})

	//db.Debug().Where("username=?", userInput.Username).Preload("Theatre").First(&user)
	//fmt.Println(user)
}
