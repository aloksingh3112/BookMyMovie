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

// func GetTheatres(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)
// 	inputUser := utils.DecodeToken(c)
// 	var user models.User
// 	err := db.Where("username =?", inputUser.Username).Preload("Theatre").First(&user).Error
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Something went wrong", "statusCode": 500})
// 		return
// 	}
// 	if len(user.Movie) == 0 {
// 		c.JSON(http.StatusOK, gin.H{"data": user.Theatre, "message": "No Data Found", "statusCode": 200})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": user.Theatre, "message": "Fetch Data Successfully", "statusCode": 200})

// }

func GetTheatres(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//inputUser := utils.DecodeToken(c)
	var theatre []models.Theatre
	err := db.Find(&theatre).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Something went wrong", "statusCode": 500})
		return
	}
	if len(theatre) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": theatre, "message": "No Data Found", "statusCode": 200})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": theatre, "message": "Fetch Data Successfully", "statusCode": 200})

}

func UpdateTheatre(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var inputTheatre types.Theatre
	var theatre models.Theatre
	c.ShouldBindJSON(&inputTheatre)
	err := db.Where("id=?", c.Param("id")).First(&theatre).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "No data found", "statusCode": 400})
		return
	}
	err = db.Model(&theatre).Update(&inputTheatre).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Some thing went wrong", "statusCode": 500})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": theatre, "message": "Data updated Successfully", "statusCode": 200})

}

func DeleteTheatre(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	inputUser := utils.DecodeToken(c)
	var theatre models.Theatre
	var user models.User
	db.Where("username=?", inputUser.Username).First(&user)
	err := db.Where("id=?", c.Param("id")).First(&theatre).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "No data found", "statusCode": 400})
		return
	}

	db.Model(&user).Association("Theatre").Delete(&theatre)
	err = db.Unscoped().Delete(&theatre).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Some thing went wrong", "statusCode": 500})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": theatre, "message": "Data deleted successfully", "statusCode": 200})

}
