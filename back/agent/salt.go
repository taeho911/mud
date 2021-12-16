package agent

import (
	"context"
	"taeho/mud/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	SALT_COLL string = "salt"
)

func saltCreateIndexes() ([]string, error) {
	var salt model.Salt
	return createIndexes(SALT_COLL, salt.IndexFields())
}

func SaltFindByUsername(ctx context.Context, username string) (model.Salt, error) {
	filter := bson.M{"username": username}
	var salt model.Salt
	if err := findOne(SALT_COLL, &salt, ctx, filter, nil); err != nil {
		return salt, err
	}
	return salt, nil
}

func SaltInsertOne(ctx context.Context, salt *model.Salt) error {
	result, err := insertOne(SALT_COLL, salt, ctx, nil)
	if err != nil {
		return err
	}
	salt.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func SaltDeleteByID(ctx context.Context, id primitive.ObjectID) (int, error) {
	filter := bson.M{"_id": id}
	result, err := deleteOne(SALT_COLL, ctx, filter, nil)
	if err != nil {
		return int(result.DeletedCount), err
	}
	return int(result.DeletedCount), nil
}

func SaltDeleteByUsername(ctx context.Context, username string) (int, error) {
	filter := bson.M{"username": username}
	result, err := deleteOne(SALT_COLL, ctx, filter, nil)
	if err != nil {
		return 0, err
	}
	return int(result.DeletedCount), err
}
