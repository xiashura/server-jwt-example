package jwt

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/xiashura/server-jwt-example/model"
)

//Client модель клиента
type Client model.Client

//Valid Проверка токена на валидность
func (client Client) Valid() error {
	tokenString := client.Token.Key
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil
	})
	return err
}

//Generate создание токена
func (client Client) Generate() (string, error) {

	claims := jwt.MapClaims{}
	claims["Authorized"] = client.Token.Authorized
	claims["User"] = client.User
	claims["exp"] = time.Now().Add(time.Hour*24).Unix() - client.Token.Time.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

//Registration регистрация пользователся
func Registration(client Client) model.Token {
	token, err := client.Generate()
	if err != nil {
		return model.Token{}
	}
	return model.Token{
		Key:        token,
		Authorized: false,
		Time:       time.Now().Add(time.Hour * 24),
	}
}

//Authentication авторизация пользователя
func Authentication(client Client) model.Client {
	if err := client.Valid(); err != nil {
		return model.Client{}
	}
	_, err := client.Generate()
	if err != nil {
		return model.Client{}
	}
	return model.Client{
		User: client.User,
		Token: model.Token{
			Key:        client.Token.Key,
			Time:       client.Token.Time,
			Authorized: true,
		},
	}
}

//Unauthenticated выход из приложения
func Unauthenticated(client Client) model.Client {
	if err := client.Valid(); err != nil {
		return model.Client{}
	}
	_, err := client.Generate()
	if err != nil {
		return model.Client{}
	}
	return model.Client{
		User: client.User,
		Token: model.Token{
			Key:        client.Token.Key,
			Time:       client.Token.Time,
			Authorized: false,
		},
	}
}

//Expired проверка на время токена
func Expired(client Client) model.Token {
	if client.Token.Time.Unix() < time.Now().Add(time.Second).Unix() {
		token, err := client.Generate()
		if err != nil {
			return client.Token
		}
		return model.Token{
			Authorized: client.Token.Authorized,
			Time:       time.Now().Add(time.Hour * 24),
			Key:        token,
		}
	}
	return client.Token
}
