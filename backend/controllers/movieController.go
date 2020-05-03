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
	var movieModel models.Movie
	var inputMovie types.Movie

	err := c.ShouldBindJSON(&inputMovie)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Something went wrong", "statusCode": 400})
		return
	}
	db.Where("username=?", inputUser.Username).First(&user)
	error := db.Where("imdb_id=?", inputMovie.ImdbID).First(&movieModel).Error
	fmt.Println(movieModel, error)
	if error == nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Movie is already present", "statusCode": 400})
		return
	}
	movie := models.Movie{Title: inputMovie.Title, Language: inputMovie.Language,
		Genre: inputMovie.Genre, Director: inputMovie.Director, StarCast: inputMovie.StarCast,
		Year: inputMovie.Year, Duration: inputMovie.Duration, ImdbID: inputMovie.ImdbID, Poster: inputMovie.Poster, UserID: user.ID}
	err = db.Save(&movie).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Something went wrong", "statusCode": 500})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie, "message": "Movie created successfully", "statusCode": 200})
	//db.Model(&user).Association("Movie").Delete(&movie)
	//db.Delete(&movie)

}

// func GetMovies(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)
// 	inputUser := utils.DecodeToken(c)
// 	var user models.User
// 	err := db.Where("username =?", inputUser.Username).Preload("Movie").First(&user).Error
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Something went wrong", "statusCode": 500})
// 		return
// 	}
// 	if len(user.Movie) == 0 {
// 		c.JSON(http.StatusOK, gin.H{"data": user.Movie, "message": "No Data Found", "statusCode": 200})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": user.Movie, "message": "Fetch Data Successfully", "statusCode": 200})

// }

func GetMovies(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var movie []models.Movie
	err := db.Find(&movie).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Something went wrong", "statusCode": 500})
		return
	}
	if len(movie) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": movie, "message": "No Data Found", "statusCode": 200})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie, "message": "Fetch Data Successfully", "statusCode": 200})

}

func UpdateMovies(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var inputMovie types.Movie
	var movie models.Movie
	c.ShouldBindJSON(&inputMovie)
	err := db.Where("id=?", c.Param("id")).First(&movie).Error

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
	err := db.Where("id=?", c.Param("id")).First(&movie).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "No data found", "statusCode": 400})
		return
	}

	db.Model(&user).Association("Movie").Delete(&movie)
	err = db.Unscoped().Delete(&movie).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Some thing went wrong", "statusCode": 500})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie, "message": "Data deleted successfully", "statusCode": 200})

}

func MapMovieWithTheatre(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//var movieInput types.Movie
	var movieDateInput types.MovieDate

	var movie models.Movie
	//var movieDate models.MovieDate
	var dateData []models.Date

	c.ShouldBindJSON(&movieDateInput)

	db.Model(&models.MovieDate{}).AddForeignKey("movie_id", "movies(id)", "CASCADE", "CASCADE")
	db.Model(&models.MovieDate{}).AddForeignKey("date_id", "dates(id)", "CASCADE", "CASCADE")
	db.Model(&models.DateTheatre{}).AddForeignKey("date_id", "dates(id)", "CASCADE", "CASCADE")
	db.Model(&models.DateTheatre{}).AddForeignKey("theatre_id", "theatres(id)", "CASCADE", "CASCADE")
	db.Model(&models.TheatreTime{}).AddForeignKey("time_id", "times(id)", "CASCADE", "CASCADE")
	db.Model(&models.TheatreTime{}).AddForeignKey("theatre_id", "theatres(id)", "CASCADE", "CASCADE")
	db.Model(&models.MovieTheatre{}).AddForeignKey("movie_id", "movies(id)", "CASCADE", "CASCADE ")
	db.Model(&models.MovieTheatre{}).AddForeignKey("theatre_id", "theatres(id)", "CASCADE", "CASCADE ")
	db.Model(&models.MovieTime{}).AddForeignKey("movie_id", "movies(id)", "CASCADE", "CASCADE ")
	db.Model(&models.MovieTime{}).AddForeignKey("time_id", "times(id)", "CASCADE", "CASCADE ")

	db.Model(&models.Seat{}).AddForeignKey("time_id", "times(id)", "CASCADE", "CASCADE")
	db.Where("id=?", movieDateInput.MovieID).First(&movie)
	fmt.Println(movieDateInput.MovieID, movie)
	for _, date := range movieDateInput.Dates {
		var t models.Date
		err := db.Where("date=?", date).First(&t).Error
		if err != nil {
			tx := models.Date{
				Date: date,
			}
			db.Save(&tx)
			dateData = append(dateData, tx)

		} else {
			dateData = append(dateData, t)
		}

	}

	db.Model(&movie).Association("Dates").Append(&dateData)
	db.Save(&movie)
	db.Model(&movie).Preload("Dates").First(&movie)

	var TheatreData []models.Theatre
	for _, id := range movieDateInput.TheatreID {
		var t models.Theatre
		db.Where("id=?", id).First(&t)
		TheatreData = append(TheatreData, t)

	}
	db.Model(&movie).Association("Theatres").Append(&TheatreData)
	db.Save(&movie)

	for _, date := range movieDateInput.Dates {
		var TheatreData []models.Theatre
		var dt models.Date
		db.Where("date=?", date).First(&dt)
		for _, id := range movieDateInput.TheatreID {
			var t models.Theatre
			db.Where("id=?", id).First(&t)
			TheatreData = append(TheatreData, t)

		}
		db.Model(&dt).Association("Theatre").Append(&TheatreData)
		db.Save(&dt)

	}

	for _, id := range movieDateInput.TheatreID {
		var times []models.Time
		var dt models.Theatre
		db.Where("id=?", id).First(&dt)
		for _, time := range movieDateInput.Times {
			var t models.Time
			err := db.Where("time=?", time).First(&t).Error
			if err != nil {
				tx := models.Time{
					Time: time,
				}
				db.Save(&tx)
				times = append(times, tx)

			} else {
				times = append(times, t)
			}

		}
		db.Model(&dt).Association("Time").Append(&times)
		db.Save(&dt)
	}

	var timesData []models.Time
	for _, id := range movieDateInput.Times {
		var t models.Time
		db.Where("time=?", id).First(&t)
		timesData = append(timesData, t)

	}
	db.Model(&movie).Association("Times").Append(&timesData)
	db.Save(&movie)

	// var dt models.Date
	c.JSON(http.StatusOK, gin.H{"data": nil, "message": "Map theatre with Movie Successfully", "statusCode": 200})

}

func GetSeatMapping(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//decodedUser := utils.DecodeToken(c)
	db.Model(&models.Seat{}).AddForeignKey("seat_map_id", "seat_maps(id)", "CASCADE", "CASCADE")
	var seatMapInput types.SeatMap
	var seatMap []models.SeatMap
	var seatData []models.SeatMap
	// var user models.User
	c.ShouldBindJSON(&seatMapInput)
	//db.Where("username?=", decodedUser.Username).First(&user)
	db.Where("movie_id=?", seatMapInput.MovieID).Where("date_id=?", seatMapInput.DateID).Where("theatre_id=?", seatMapInput.TheatreID).Where("time_id=?", seatMapInput.TimeID).Find(&seatMap)
	fmt.Println(seatMap)
	for _, data := range seatMap {
		var seatMap models.SeatMap
		db.Where("id=?", data.ID).Preload("Seat").Find(&seatMap)
		seatData = append(seatData, seatMap)

	}
	c.JSON(http.StatusOK, gin.H{"data": seatData, "message": "Seat Data", "statusCode": 200})

}
