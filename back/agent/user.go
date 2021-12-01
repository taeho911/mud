package agent

import (
	"context"
	"taeho/mud/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	USER_COLL string = "user"
)

func UserCreateIndexes() ([]string, error) {
	var user model.User
	name, err := createIndexes(USER_COLL, user.IndexFields())
	if err != nil {
		return nil, err
	}
	return name, err
}

func UserFindByUsername(ctx context.Context, username string) (*model.User, error) {
	filter := bson.M{"username": username}
	var entity model.User
	if err := findOne(USER_COLL, &entity, ctx, filter, nil); err != nil {
		return &entity, err
	}
	return &entity, nil
}

func UserInsertOne(ctx context.Context, entity *model.User) (*model.User, error) {
	result, err := insertOne(USER_COLL, entity, ctx, nil)
	if err != nil {
		return entity, err
	}
	entity.ID = result.InsertedID.(primitive.ObjectID)
	return entity, nil
}

func UserDeleteByID(ctx context.Context, id primitive.ObjectID) (int, error) {
	filter := bson.M{"_id": id}
	result, err := deleteOne(USER_COLL, ctx, filter, nil)
	if err != nil {
		return int(result.DeletedCount), err
	}
	return int(result.DeletedCount), nil
}
