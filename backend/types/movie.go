package types

type Movie struct {
	MovieID  uint   `json:"movieID"`
	Title    string `json:"title"`
	Language string `json:"language"`
	Genre    string `json:"genre"`
	Director string `json:"director"`
	StarCast string `json:"starcast"`
	Year     string `json:"year"`
	Duration string `json:"duration"`
}
