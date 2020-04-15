package models

import (
	"github.com/jinzhu/gorm"
)

type Seat struct {
	gorm.Model
	SeatNumber string `json:"seatnumber"`
	IsBooked   bool   `json:"isbooked" goom:"default:false"`
}
