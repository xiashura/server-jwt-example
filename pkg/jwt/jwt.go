package jwt

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/xiashura/server-jwt-example/model"
)

//Client модель клиента
type Client struct {
	User  model.User
	Token model.Token
}

//Valid Проверка токена на валидность
func (client Client) Valid() error {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tokenString := client.Token.Key
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWTPASS")), nil
	})
	return err
}

//Generate создание токена
func (client Client) Generate() (string, error) {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	claims := jwt.MapClaims{}
	claims["Authorized"] = client.User.Authorized
	claims["User"] = client.User
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWTPASS")))
}

//Registration регистрация пользователся
func Registration(client Client) model.Token {
	token, err := client.Generate()
	if err != nil {
		return model.Token{}
	}
	return model.Token{
		Key: token,
	}
}

//Authentication авторизация пользователя
func Authentication(client Client) model.User {
	if err := client.Valid(); err != nil {
		return model.User{}
	}
	_, err := client.Generate()
	if err != nil {
		return model.User{}
	}

	///aaaa
	return model.User{}
}

//Unauthenticated выход из приложения
func Unauthenticated(client Client) model.User {
	if err := client.Valid(); err != nil {
		return model.User{}
	}
	_, err := client.Generate()
	if err != nil {
		return model.User{}
	}
	return model.User{}
}

func Expired(client Client) model.Token {
	if client.User.Time.Unix() < time.Now().Add(time.Second).Unix() {
		token, err := client.Generate()
		if err != nil {
			return model.Token{}
		}
		return model.Token{
			Key: token,
		}
	}
	return client.Token
}
