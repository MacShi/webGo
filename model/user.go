package model

type UserLogin struct {
	Id   int `json:"id"`
	UserName string	`json:"user_name"`
	Pwd string `json:"pwd"`
}