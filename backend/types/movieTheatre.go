package types

type MovieTheatre struct {
	MovieID   uint   `json:"movieID"`
	TheatreID []uint `json:"theatreIDS"`
}
