package tokens

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/xiashura/server-jwt-example/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
		return err
	}
	// Проверка Access токена
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(client.Token.Access, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_PASS")), nil
	})

	//Проверка Refresh токена
	rclaims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(client.Token.Refresh, rclaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_PASS")), nil
	})

	return err

}

//ValidHash get hash resfresh token in stirng
func (client Client) ValidHash() string {
	hmacSecret := []byte(os.Getenv("JWT_PASS"))
	token, err := jwt.Parse(client.Token.Refresh, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		a := claims["GUID"]
		return fmt.Sprintf("%v", a)
	}

	return ""
}

//Generate Access и Refresh создание токена
func (client *Client) Generate(id primitive.ObjectID) (model.Token, error) {

	err := godotenv.Load()

	if err != nil {
		return model.Token{}, err
	}
	// Create Access
	claims := jwt.MapClaims{}
	claims["Authorized"] = client.User.Authorized
	claims["Email"] = client.User.Email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_PASS")))

	if err != nil {
		return model.Token{}, err
	}

	// Create Refresh
	rclaims := jwt.MapClaims{}

	rclaims["Authorized"] = client.User.Authorized
	rclaims["GUID"] = id
	rclaims["sub"] = 1
	rclaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, rclaims)

	rt, err := refreshToken.SignedString([]byte(os.Getenv("JWT_PASS")))

	if err != nil {
		return model.Token{}, err
	}
	client.Token.Access = t
	client.Token.Refresh = rt
	return model.Token{
		Access:  t,
		Refresh: rt,
	}, nil
}
