package agent

import (
	"context"
	"taeho/mud/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ACCOUNT_COLL string = "account"
)

func accountCreateIndexes() ([]string, error) {
	var account model.Account
	return createIndexes(ACCOUNT_COLL, account.IndexFields())
}

func AccountInsertOne(ctx context.Context, entity model.Account) (model.Account, error) {
	result, err := insertOne(ACCOUNT_COLL, &entity, ctx, nil)
	if err != nil {
		return entity, err
	}
	entity.ID = result.InsertedID.(primitive.ObjectID)
	return entity, nil
}

func AccountFindAll(ctx context.Context) ([]model.Account, error) {
	filter := bson.M{"deleted": false}
	// mongodb의 소트에는 1(asc), -1(desc)이 존재한다.
	option := options.Find().SetSort(bson.D{primitive.E{Key: "index", Value: 1}})
	var entity []model.Account
	if err := find(ACCOUNT_COLL, &entity, ctx, filter, option); err != nil {
		return nil, err
	}
	return entity, nil
}
