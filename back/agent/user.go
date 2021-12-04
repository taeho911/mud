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
	return createIndexes(USER_COLL, user.IndexFields())
}

func UserFindByUsername(ctx context.Context, username string) (model.User, error) {
	filter := bson.M{"username": username}
	var user model.User
	if err := findOne(USER_COLL, &user, ctx, filter, nil); err != nil {
		return user, err
	}
	return user, nil
}

func UserInsertOne(ctx context.Context, user *model.User) error {
	result, err := insertOne(USER_COLL, user, ctx, nil)
	if err != nil {
		return err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func UserDeleteByID(ctx context.Context, id primitive.ObjectID) (int, error) {
	filter := bson.M{"_id": id}
	result, err := deleteOne(USER_COLL, ctx, filter, nil)
	if err != nil {
		return int(result.DeletedCount), err
	}
	return int(result.DeletedCount), nil
}
