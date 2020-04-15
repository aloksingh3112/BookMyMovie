package models

import "github.com/jinzhu/gorm"

type Theatre struct {
	gorm.Model
	TheatreName string  `json:"theatrename" gorm:"not null"`
	Location    string  `json:"location" gorm:"not null"`
	City        string  `json:"city" gorm:"not null"`
	Movie       []Movie `gorm:"many2many:movie_theatre;"`
}