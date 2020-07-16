package model

import "time"

//User модель пользователя
type User struct {
	Email    string `json:"Email,omitempty"`
	Name     string `json:"Name,omitempty"`
	Password string `json:"Password,omitempty"`
}

//Client модель клиента
type Client struct {
	User  User  `json:"User,omitempty"`
	Token Token `json:"Token,omitempty"`
}

//Token модель токена
type Token struct {
	Time       time.Time `json:"Time,omitempty"`
	Authorized bool      `json:"Authorized,omitempty"`
	Key        string    `json:"Key,omitempty"`
}
