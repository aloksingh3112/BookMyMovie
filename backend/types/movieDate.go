package types

type MovieDate struct {
	MovieID   string   `json:"movieID"`
	TheatreID []uint   `json:"theatreIDS"`
	Dates     []string `json:"dates"`
	Times     []string `json:"times"`
}
