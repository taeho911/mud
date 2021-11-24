package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// 입력되지 않은 필드를 무시하고 싶을 경우: omitempty
// bson:_id를 omitempty하지 않을 경우 mongodb에 insert시 000000... 의 디폴트 _id가 생성되므로 주의
type Acc struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Owner    string             `bson:"onwer,omitempty"`
	Index    int                `bson:"index,omitempty"`
	Title    string             `bson:"title,omitempty"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
	Location string             `bson:"location,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Memo     string             `bson:"memo,omitempty"`
	Alias    []string           `bson:"alias,omitempty"`
	Deleted  bool               `bson:"deleted,omitempty"`
}
