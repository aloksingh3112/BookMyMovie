package controller

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/aloksingh3112/BookMyMovie/models"
	"github.com/aloksingh3112/BookMyMovie/types"
	"github.com/aloksingh3112/BookMyMovie/utils"
	"github.com/gin-gonic/gin"
)

func AddMovie(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	inputUser := utils.DecodeToken(c)
	var user models.User

	var inputMovie types.Movie

	err := c.ShouldBindJSON(&inputMovie)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Something went wrong", "statusCode": 400})
		return
	}
	db.Where("username=?", inputUser.Username).First(&user)

	movie := models.Movie{Title: inputMovie.Title, MovieID: inputMovie.MovieID, Language: inputMovie.Language,
		Genre: inputMovie.Genre, Director: inputMovie.Director, StarCast: inputMovie.StarCast,
		Year: inputMovie.Year, Duration: inputMovie.Duration, UserID: user.ID}
	err = db.Save(&movie).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Something went wrong", "statusCode": 500})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie, "message": "Movie created successfully", "statusCode": 200})
	//db.Model(&user).Association("Movie").Delete(&movie)
	//db.Delete(&movie)

}

func GetMovies(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	inputUser := utils.DecodeToken(c)
	var user models.User
	err := db.Where("username =?", inputUser.Username).Preload("Movie").First(&user).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Something went wrong", "statusCode": 500})
		return
	}
	if len(user.Movie) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": user.Movie, "message": "No Data Found", "statusCode": 200})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user.Movie, "message": "Fetch Data Successfully", "statusCode": 200})

}

func UpdateMovies(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var inputMovie types.Movie
	var movie models.Movie
	c.ShouldBindJSON(&inputMovie)
	err := db.Where("movie_id=?", c.Param("id")).First(&movie).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Some thing went wrong", "statusCode": 500})
		return
	}
	err = db.Model(&movie).Update(&inputMovie).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Some thing went wrong", "statusCode": 400})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie, "message": "Data updated Successfully", "statusCode": 200})

}

func DeleteMovie(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	inputUser := utils.DecodeToken(c)
	var movie models.Movie
	var user models.User
	db.Where("username=?", inputUser.Username).First(&user)
	err := db.Where("movie_id=?", c.Param("id")).First(&movie).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "No data found", "statusCode": 400})
		return
	}

	db.Model(&user).Association("Movie").Delete(&movie)
	err = db.Delete(&movie).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Some thing went wrong", "statusCode": 500})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie, "message": "Data deleted successfully", "statusCode": 200})

}
