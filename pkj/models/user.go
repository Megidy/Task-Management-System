package models

type User struct {
	Id       int    `json:"id"`
	Username int    `json:"username"`
	Password int    `json:"password"`
	Role     string `json:"role"`
}
