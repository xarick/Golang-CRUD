package models

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type UserCrUp struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
