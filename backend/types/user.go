package types

type User struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	TypeOfUser string `json:"typeofuser"`
}
