package models

import (
	"github.com/jinzhu/gorm"
)

type Movie struct {
	gorm.Model
	Title    string `json:"title" gorm:"not null"`
	Language string `json:"language" gorm:"not null;"`
	Genre    string `json:"genre" gorm:"not null;"`
	Director string `json:"director" gorm:"not null"`
	StarCast string `json:"starcast" gorm:"not null"`
	Year     string `json:"year" gorm:"not null"`
	Duration string `json:"duration" gorm:"not null"`
	Poster   string `json:"poster"`
	ImdbID   string `json:"imdbID"`
	Dates    []Date `json:"dates" gorm:"many2many:movie_dates;"`
	UserID   uint
}
