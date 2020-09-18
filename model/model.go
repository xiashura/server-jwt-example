package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User модель пользователя
type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email      string             `bson:"Email,omitempty" json:"Email,omitempty"`
	Name       string             `bson:"Name,omitempty" json:"Name,omitempty"`
	Password   string             `bson:"Password,omitempty" json:"Password,omitempty"`
	HashTokens []string           `bson:"HashTokens,omitempty" json:"Token,omitempty"`
	Authorized bool               `bson:"Authorized,omitempty" json:"Authorized,omitempty"`
}

//Token модель токена
type Token struct {
	Access  string `json:"Access,omitempty"`
	Refresh string `json:"Refresh,omitempty"`
}
