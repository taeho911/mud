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
	Owner    string             `bson:"owner"`
	Title    string             `bson:"title"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Location string             `bson:"location"`
	Email    string             `bson:"email"`
	Memo     string             `bson:"memo"`
	Alias    []string           `bson:"alias"`
	Maketime time.Time          `bson:"maketime"`
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
