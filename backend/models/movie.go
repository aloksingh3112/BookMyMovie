package models

type Movie struct {
	MovieID  uint      `gorm:"primary_key;unique"`
	Title    string    `json:"title" gorm:"not null"`
	Language string    `json:"language" gorm:"not null;"`
	Genre    string    `json:"genre" gorm:"not null;"`
	Director string    `json:"director" gorm:"not null"`
	StarCast string    `json:"starcast" gorm:"not null"`
	Year     string    `json:"year" gorm:"not null"`
	Duration string    `json:"duration" gorm:"not null"`
	Theatre  []Theatre `json:"theatre" gorm:"many2many:movie_theatre;"`
}
