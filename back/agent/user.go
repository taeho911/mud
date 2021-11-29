package agent

import (
	"context"
	"taeho/mud/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	userCollname string = "user"
)

func UserFindByUsername(ctx context.Context, username string) (*model.User, error) {
	filter := bson.M{"username": username}
	var entity model.User
	if err := findOne(userCollname, &entity, ctx, filter, nil); err != nil {
		return &entity, err
	}
	return &entity, nil
}

func UserInsertOne(ctx context.Context, entity *model.User) (*model.User, error) {
	result, err := insertOne(userCollname, entity, ctx, nil)
	if err != nil {
		return entity, err
	}
	entity.ID = result.InsertedID.(primitive.ObjectID)
	return entity, nil
}

func UserDeleteByID(ctx context.Context, id primitive.ObjectID) (int, error) {
	filter := bson.M{"_id": id}
	result, err := deleteOne(userCollname, ctx, filter, nil)
	if err != nil {
		return int(result.DeletedCount), err
	}
	return int(result.DeletedCount), nil
}
