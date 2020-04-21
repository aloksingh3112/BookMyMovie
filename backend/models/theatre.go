package models

import "github.com/jinzhu/gorm"

type Theatre struct {
	gorm.Model
	TheatreName string `json:"theatrename" gorm:"not null"`
	Location    string `json:"location" gorm:"not null"`
	City        string `json:"city" gorm:"not null"`
	UserID      uint
	Time        []Time `gorm:"many2many:theatre_times;"`
}
