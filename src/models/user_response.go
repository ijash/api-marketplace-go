package models

type UserResponse struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	FullName string `json:"fullName"`
	Token    string `json:"token"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
