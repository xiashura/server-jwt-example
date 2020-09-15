package model

import (
	"time"
)

//User модель пользователя
type User struct {
	Email      string    `json:"Email,omitempty"`
	Name       string    `json:"Name,omitempty"`
	Password   string    `json:"Password,omitempty"`
	Authorized bool      `json:"Authorized,omitempty"`
	Time       time.Time `json:"Time,omitempty"`
}

//Token модель токена
type Token struct {
	Access  string `json:"Access,omitempty"`
	Refresh string `json:"Refresh,omitempty"`
}
