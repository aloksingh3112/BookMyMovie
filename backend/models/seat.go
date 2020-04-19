package models

import (
	"github.com/jinzhu/gorm"
)

type Seat struct {
	gorm.Model
	SeatNumber string `json:"seatnumber"`
	IsBooked   bool   `json:"isbooked" gorm:"default:false"`
	Price      string `json:"price"`
	TimeID     uint
}
