package models

type User struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
}
