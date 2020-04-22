package models

import "github.com/jinzhu/gorm"

type Booking struct {
	gorm.Model
	UserID    uint `json:"userId"`
	MovieID   uint `json:"movieId"`
	DateID    uint `json:"dateId"`
	TheatreID uint `json:"theatreId"`
	TimeID    uint `json:"timeId"`
	Seat      []Seat
}
