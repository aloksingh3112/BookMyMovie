package models

import (
	"github.com/jinzhu/gorm"
)

type DateTime struct {
	gorm.Model
	Date    string  `json:"date"`
	Times   []Time  `json:"timing" gorm:"ForeignKey:DateTimeID"`
	Movie   Movie   `gorm:"ForeignKey:DateTimeID"`
	Theatre Theatre `gorm:"ForeignKey:DateTimeID"`
}
