package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Maketime time.Time          `bson:"maketime"`
}

func (user *User) NotNullFields() []interface{} {
	return []interface{}{
		user.Username,
		user.Password,
	}
}

func (user *User) IndexFields() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys:    bson.M{"username": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{"maketime": -1},
		},
	}
}

func (user *User) SetMaketime() {
	user.Maketime = time.Now()
}
