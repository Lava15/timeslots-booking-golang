package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	isAdmin  bool   `json:"is_admin"`
}
