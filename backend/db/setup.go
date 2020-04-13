package db

import (
	"github.com/aloksingh3112/BookMyMovie/models"

	"github.com/aloksingh3112/BookMyMovie/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Setup function to connect with db
func Setup() *gorm.DB {
	db, err := gorm.Open("mysql", config.DBPATH)

	if err != nil {
		panic("Unable to connect to db")
	}
	db.AutoMigrate(&models.Movie{})
	return db

}
