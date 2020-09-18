package db

import (
	ctx "context"
	"encoding/base64"
	"fmt"

	"github.com/xiashura/server-jwt-example/internal/app/mycrypt"
	"github.com/xiashura/server-jwt-example/internal/app/tokens"
	"github.com/xiashura/server-jwt-example/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	User  model.User
	Token model.Token
}

//Auth_db подключение к базе данных
func (d Data) Auth_db() (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI("mongodb+srv://plexxxanov:aXQA8skfqNLjIj6I@cluster0.efpzs.gcp.mongodb.net/test?retryWrites=true&w=majority")

	client, err := mongo.Connect(ctx.TODO(), clientOptions)

	if err != nil {
		return &mongo.Client{}, nil
	}

	return client, nil
}

//Create_User_db
func (d *Data) Create_User_db() (model.Token, error) {

	client, err := d.Auth_db()

	if err != nil {
		return model.Token{}, err
	}

	con := client.Database("test").Collection("user")
	clinetjwt := tokens.Client{}

	id := d.User.ID
	token, err := clinetjwt.Generate(id)
	hashToken := base64.StdEncoding.EncodeToString([]byte(mycrypt.Mycrypt(token.Refresh)))

	d.User.HashTokens = append(d.User.HashTokens, hashToken)
	for _, el := range d.User.HashTokens {
		fmt.Println(el)
	}

	_, err = con.InsertOne(ctx.TODO(), d.User)

	if err != nil {
		return model.Token{}, err
	}

	return token, nil
}

//Append_token to hashTokens on user db
func (d *Data) Append_token(token string) {

	client, err := d.Auth_db()
	if err != nil {
		fmt.Println(err)
	}

	hashToken := base64.StdEncoding.EncodeToString([]byte(mycrypt.Mycrypt(token)))
	con := client.Database("test").Collection("user")

	res, err := con.UpdateOne(ctx.TODO(), d.User, bson.M{"$push": bson.M{
		"HashTokens": hashToken,
	}})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func (d Data) Find_User_db_by_id(id string) (model.User, error) {

	client, err := d.Auth_db()

	if err != nil {
		return model.User{}, err
	}

	user := &model.User{}
	con := client.Database("test").Collection("user")
	objID, _ := primitive.ObjectIDFromHex(id)

	res := con.FindOne(ctx.Background(), bson.M{"_id": objID})
	err = res.Decode(user)

	if err != nil {
		return model.User{}, err
	}

	return *user, nil
}
func (d *Data) Find_User_db() (model.User, error) {

	client, err := d.Auth_db()

	if err != nil {
		return model.User{}, err
	}

	user := &model.User{}
	con := client.Database("test").Collection("user")

	res := con.FindOne(ctx.Background(), d.User)
	err = res.Decode(user)

	if err != nil {
		return model.User{}, err
	}
	d.User = *user

	return *user, nil
}

//Delete_Refresh_Token
func (d Data) Delete_Refresh_Token() (*mongo.UpdateResult, error) {

	client, err := d.Auth_db()

	if err != nil {
		fmt.Println(err)
	}

	con := client.Database("test").Collection("user")

	hashToken := base64.StdEncoding.EncodeToString([]byte(mycrypt.Mycrypt(d.Token.Refresh)))

	return con.UpdateOne(ctx.TODO(), bson.M{"HashTokens": hashToken},
		bson.M{"$pull": bson.M{"HashTokens": hashToken}})

}

//Delete_AllRefresh_Token
func (d Data) Delete_AllRefresh_Token() (*[]mongo.UpdateResult, error) {

	client, err := d.Auth_db()

	if err != nil {
		fmt.Println(err)
	}

	con := client.Database("test").Collection("user")
	res, err := d.Find_User_db()
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}

	results := []mongo.UpdateResult{}
	for i := range res.HashTokens {
		ress, err := con.UpdateOne(ctx.TODO(), bson.M{"_id": d.User.ID},
			bson.M{"$pull": bson.M{"HashTokens": res.HashTokens[i]}})
		if err != nil {
			fmt.Println(err)
		}
		results = append(results, *ress)
	}
	return &results, err
}

///
///
///
///
