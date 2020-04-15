package db

import (
	"github.com/aloksingh3112/BookMyMovie/models"

	"github.com/aloksingh3112/BookMyMovie/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//Setup function to connect with db
func Setup() *gorm.DB {
	db, err := gorm.Open("mysql", config.DBPATH)

	if err != nil {
		panic("Unable to connect to db")
	}
	db.AutoMigrate(&models.Movie{}, &models.Theatre{}, &models.Booking{}, &models.Seat{}, &models.User{})
	return db

}