package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Money struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string             `bson:"username" json:"username"`
	Date     time.Time          `bson:"date" json:"date"`
	Amount   float64            `bson:"amount" json:"amount"`
	Summary  string             `bson:"summary" json:"summary"`
	Tags     []string           `bson:"tags" json:"tags"`
	Maketime time.Time          `bson:"maketime" json:"maketime"`
}

func (money *Money) NotNullFields() []interface{} {
	return []interface{}{
		money.Username,
		money.Date,
		money.Amount,
	}
}

func (money *Money) IndexFields() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.D{
				primitive.E{Key: "username", Value: 1},
				primitive.E{Key: "date", Value: -1},
				primitive.E{Key: "tags", Value: 1},
			},
		},
	}
}

func (money *Money) SetMaketime() {
	money.Maketime = time.Now()
}
