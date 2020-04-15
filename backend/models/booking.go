package models

import "github.com/jinzhu/gorm"

type Booking struct {
	gorm.Model
	Movie  Movie
	Theate Theatre
}
