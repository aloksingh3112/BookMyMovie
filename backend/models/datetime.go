package models

import (
	"github.com/jinzhu/gorm"
)

type DateTime struct {
	gorm.Model
	Date           string `json:"date"`
	Time           []Time `json:"timing" gorm:"ForeignKey:DateTimeID"`
	MovieTheatreID uint
}
