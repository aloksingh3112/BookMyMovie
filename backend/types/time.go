package types

type Time struct {
	Time  string `json:"time"`
	Seats []Seat `json:"seats"`
}
