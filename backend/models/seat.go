package models

import (
	"github.com/jinzhu/gorm"
)

type Seat struct {
	gorm.Model
	SeatNumber string `json:"seatnumber"`
	Price      string `json:"price"`
	Type       string `json:"type"`
	SeatMapID  uint
}
