package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MoneyAutoInput struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username  string             `bson:"username" json:"username"`
	Inputdate string             `bson:"inputdate" json:"inputdate"`
	Amount    float64            `bson:"amount" json:"amount"`
	Summary   string             `bson:"summary" json:"summary"`
	Tags      []string           `bson:"tags" json:"tags"`
	Maketime  time.Time          `bson:"maketime" json:"maketime"`
}

func (moneyAutoInput *MoneyAutoInput) NotNullFields() []interface{} {
	return []interface{}{
		moneyAutoInput.Username,
		moneyAutoInput.Inputdate,
		moneyAutoInput.Amount,
	}
}

func (moneyAutoInput *MoneyAutoInput) IndexFields() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.D{
				primitive.E{Key: "username", Value: 1},
				primitive.E{Key: "inputdate", Value: 1},
			},
		},
	}
}

func (moneyAutoInput *MoneyAutoInput) SetMaketime() {
	moneyAutoInput.Maketime = time.Now()
}
