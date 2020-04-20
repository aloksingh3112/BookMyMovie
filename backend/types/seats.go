package types

type Seat struct {
	SeatNumber string `json:"seatnumber"`
	IsBooked   bool   `json:"isbooked"`
	Price      string `json:"price"`
	TimeID     uint
}
