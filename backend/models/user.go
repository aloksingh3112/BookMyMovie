package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username   string    `json:"username" gorm:"not null;unique"`
	Password   string    `json:"password" gorm:"not null"`
	Email      string    `json:"email" gorm:"unique;notnull"`
	Name       string    `json:"name" gorm:"not_null"`
	TypeOfUser string    `json:"typeofuser" gorm:"default:'USER'"`
	Movie      []Movie   `json:"movie"`
	Theatre    []Theatre `json:"theatre"`
	Booking    []Booking `gorm:"foreignkey:ID"`
}
