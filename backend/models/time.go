package models

import (
	"github.com/jinzhu/gorm"
)

type Time struct {
	gorm.Model
	Time string `json:"time"`
}
