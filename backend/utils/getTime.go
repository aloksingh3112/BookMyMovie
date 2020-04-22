package utils

var alphabet []string = []string{"A", "B", "C", "D", "E", "F"}

// func GetTime(c *gin.Context, times []string) []models.Time {
// 	db := c.MustGet("db").(*gorm.DB)
// 	var timeData []models.Time
// 	for index, t := range times {
// 		seats := GenSeat(c, index)
// 		time := models.Time{
// 			Time:  t,
// 			Seats: seats,
// 		}
// 		db.Save(&time)
// 		timeData = append(timeData, time)

// 	}
// 	// fmt.Println("timedata", timeData)
// 	return timeData
// }

// func GenSeat(c *gin.Context, i int) []models.Seat {
// 	db := c.MustGet("db").(*gorm.DB)
// 	var seats []models.Seat
// 	var s = alphabet[i]
// 	for i := 0; i < 20; i++ {
// 		var seat models.Seat
// 		seat = models.Seat{
// 			SeatNumber: s + string(i),
// 			IsBooked:   false,
// 			Price:      "500",
// 		}

// 		db.Save(&seat)
// 		seats = append(seats, seat)
// 		fmt.Println("seats", seats)

// 	}
// 	return seats
// }
