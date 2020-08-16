package auth

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/xiashura/server-jwt-example/pkg/api"
	"golang.org/x/net/context"
)

type Client struct {
	reqU api.UserRequest
	reqJ api.JwtRequest
	res  api.JwtResponse
	pd   api.UnimplementedAuthServer
}

//Valid Проверка токена на валидность
func (client *Client) Valid(s string) error {
	tokenString := s
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil
	})
	return err
}

//Generate создание токена
func (client *Client) Generate(e string) (string, error) {
	claims := jwt.MapClaims{}
	claims["Authorized"] = client.res.Authorized
	claims["Email"] = e
	claims["exp"] = time.Now().Add(time.Hour*24).Unix() - client.res.Time
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func (c *Client) Registration(ctx context.Context, r *api.UserRequest) (*api.JwtResponse, error) {
	c.res.Authorized = true
	token, err := c.Generate(r.GetEmail())
	if err != nil {
		log.Fatal(err)
	}
	return &api.JwtResponse{
		Time:       time.Now().Add(time.Hour * 24).Unix(),
		Authorized: true,
		Key:        token,
	}, nil
}

func (c *Client) Authentication(ctx context.Context, r *api.JwtRequest) (*api.JwtResponse, error) {
	if err := c.Valid(r.GetKey()); err != nil {
		return &api.JwtResponse{
			Time:       0,
			Authorized: false,
			Key:        "",
		}, nil
	}
	return &api.JwtResponse{
		Time:       r.GetTime(),
		Authorized: r.GetAuthorized(),
		Key:        r.GetKey(),
	}, nil
}

/*
func (c *Client) Unauthenticated(ctx context.Context, r *api.JwtRequest) (*api.JwtResponse, error) {
	if r.GetTime() < time.Now().Add(time.Second).Unix() {
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


	return &api.JwtResponse{Time: 23, Authorized: true, Key: "qweqe"}, nil
}

func (c *Client) Expired(ctx context.Context, r *api.JwtRequest) (*api.JwtResponse, error) {

	return &api.JwtResponse{Time: 23, Authorized: true, Key: "qweqe"}, nil
}
*/
//Client модель клиента
/*
type Client model.Client


//Registration регистрация пользователся
func Registration(client Client) model.Token {
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
*/
