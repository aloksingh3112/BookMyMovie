package models

import (
	"github.com/jinzhu/gorm"
)

type MovieTheatre struct {
	gorm.Model
	MovieID   uint
	TheatreID uint
	Dates     []DateTime `gorm:"ForeignKey:MovieTheatreID"`
}
