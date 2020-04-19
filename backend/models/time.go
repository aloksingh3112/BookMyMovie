package models

type Time struct {
	Time       string `json:"time"`
	Seats      []Seat `json:"seats" ForeignKey:"TimeID"`
	DateTimeID uint
}
