package models

import (
	"github.com/jinzhu/gorm"
)

type Date struct {
	gorm.Model
	Date    string    `json:"date"`
	Theatre []Theatre `gorm:"many2many:date_theatres;"`
}
