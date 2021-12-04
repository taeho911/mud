package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Salt struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Salt     []byte             `bson:"salt"`
	Maketime time.Time          `bson:"maketime"`
}

func (salt *Salt) NotNullFields() []interface{} {
	return []interface{}{
		salt.Username,
		salt.Salt,
	}
}

func (salt *Salt) IndexFields() []mongo.IndexModel {
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

func (salt *Salt) SetMaketime() {
	salt.Maketime = time.Now()
}
