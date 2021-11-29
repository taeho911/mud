package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// 입력되지 않은 필드를 무시하고 싶을 경우: omitempty
// bson:_id를 omitempty하지 않을 경우 mongodb에 insert시 000000... 의 디폴트 _id가 생성되므로 주의
type Account struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Owner    string             `bson:"onwer,omitempty"`
	Title    string             `bson:"title,omitempty"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
	Location string             `bson:"location,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Memo     string             `bson:"memo,omitempty"`
	Alias    []string           `bson:"alias,omitempty"`
	Maketime time.Time          `bson:"maketime,omitempty"`
}

func (account *Account) NotNullFields() []interface{} {
	return []interface{}{
		account.Owner,
		account.Title,
		account.Username,
		account.Password,
	}
}

func (account *Account) IndexFields() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.M{"Owner": 1},
		},
		{
			Keys: bson.M{"maketime": -1},
		},
	}
}

func (account *Account) SetMaketime() {
	account.Maketime = time.Now()
}
