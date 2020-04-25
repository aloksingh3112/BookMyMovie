package types

type Movie struct {
	ID       string `json:"movieId"`
	Title    string `json:"title"`
	Language string `json:"language"`
	Genre    string `json:"genre"`
	Director string `json:"director"`
	StarCast string `json:"starcast"`
	Year     string `json:"year"`
	Duration string `json:"duration"`
	Poster   string `json:"poster"`
	Theatre  []uint `json:"theatreIds"`
}
