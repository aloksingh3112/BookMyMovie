package utils

import (
	"fmt"

	"github.com/aloksingh3112/BookMyMovie/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var alphabet []string = []string{"A", "B", "C", "D", "E", "F"}

func GetTime(c *gin.Context, times []string) []models.Time {
	db := c.MustGet("db").(*gorm.DB)
	var timeData []models.Time
	for index, t := range times {
		seats := GenSeat(c, index)
		time := models.Time{
			Time:  t,
			Seats: seats,
		}
		db.Save(&time)
		timeData = append(timeData, time)

	}
	// fmt.Println("timedata", timeData)
	return timeData
}

func GenSeat(c *gin.Context, i int) []models.Seat {
	db := c.MustGet("db").(*gorm.DB)
	var seats []models.Seat
	var s = alphabet[i]
	for i := 0; i < 4; i++ {
		var seat models.Seat
		seat = models.Seat{
			SeatNumber: s + string(i),
			IsBooked:   false,
			Price:      "500",
		}

		db.Save(&seat)
		seats = append(seats, seat)
		fmt.Println("seats", seats)

	}
	return seats
}
