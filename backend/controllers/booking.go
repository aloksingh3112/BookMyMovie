package controller

import (
	"net/http"

	"github.com/aloksingh3112/BookMyMovie/models"
	"github.com/aloksingh3112/BookMyMovie/types"
	"github.com/aloksingh3112/BookMyMovie/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func BookMovie(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	decodeUser := utils.DecodeToken(c)

	db.Model(&models.Seat{}).AddForeignKey("seat_map_id", "seat_maps(id)", "CASCADE", "CASCADE")
	db.Model(&models.SeatMap{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")

	var seatMapInput types.SeatMap

	var user models.User
	c.ShouldBindJSON(&seatMapInput)
	db.Where("username=?", decodeUser.Username).First(&user)

	var data []models.Seat
	for i := 0; i < len(seatMapInput.Seat); i++ {
		data = append(data, seatMapInput.Seat[i])
	}
	seatData := models.SeatMap{
		UserID:    user.ID,
		MovieID:   seatMapInput.MovieID,
		DateID:    seatMapInput.DateID,
		TheatreID: seatMapInput.TheatreID,
		TimeID:    seatMapInput.TimeID,
		Seat:      data,
	}
	db.Model(&seatData).Association("Seat").Append(&data)
	db.Save(&seatData)
	db.Model(&user).Association("Booking").Append(&seatData)
	db.Save(&user)

	db.Model(&seatData).Preload("Booking").First(&seatData)
	c.JSON(http.StatusOK, gin.H{"data": seatData, "message": "Ticket booked successfully", "statusCode": 200})

}

func CancelTicket(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var seatMap models.SeatMap
	err := db.Where("id=?", c.Param("id")).First(&seatMap).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": err, "message": "Something went wrong", "statusCode": 500})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": seatMap, "message": "Ticket canceled successfully", "statusCode": 200})

}
