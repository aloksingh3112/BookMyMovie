package controller

import (
	"fmt"

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
	}
	db.Where("username=?", inputUser.Username).First(&user)

	movie := models.Movie{Title: inputMovie.Title, MovieID: inputMovie.MovieID, Language: inputMovie.Language,
		Genre: inputMovie.Genre, Director: inputMovie.Director, StarCast: inputMovie.StarCast,
		Year: inputMovie.Year, Duration: inputMovie.Duration, UserID: user.ID}
	db.Debug().Save(&movie)

}
