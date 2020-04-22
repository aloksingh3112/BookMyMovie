package types

import "github.com/aloksingh3112/BookMyMovie/models"

type Booking struct {
	UserID    uint          `json:"userId"`
	MovieID   uint          `json:"movieId"`
	DateID    uint          `json:"dateId"`
	TheatreID uint          `json:"theatreId"`
	TimeID    uint          `json:"timeId"`
	Seat      []models.Seat `json:"seats"`
}
