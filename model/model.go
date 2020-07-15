package model

//User модель пользователя в приложение
type User struct {
	Email    string `json:"Email,omitempty"`
	Name     string `json:"Name,omitempty"`
	Password string `json:"Password,omitempty"`
}
