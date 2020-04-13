package models

type Movie struct {
	ID    string `json:"id" gorm:"primary_key"`
	TITLE string `json:"title"`
}
