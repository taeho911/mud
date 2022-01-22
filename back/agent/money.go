package agent

import (
	"context"
	"taeho/mud/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONEY_COLL string = "money"
)

func moneyCreateIndexes() ([]string, error) {
	var money model.Money
	return createIndexes(MONEY_COLL, money.IndexFields())
}

func MoneyInsertOne(ctx context.Context, money *model.Money) error {
	result, err := insertOne(MONEY_COLL, money, ctx, nil)
	if err != nil {
		return err
	}
	money.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func MoneyFindByUsername(ctx context.Context, username string) ([]model.Money, error) {
	filter := bson.M{"username": username}
	option := options.Find().SetSort(bson.M{"date": -1})
	// option := options.Find().SetSort(bson.M{"date": -1}).SetSkip(0).SetLimit(10)
	var money []model.Money
	if err := find(MONEY_COLL, &money, ctx, filter, option); err != nil {
		return money, err
	}
	if len(money) == 0 {
		money = make([]model.Money, 0)
	}
	return money, nil
}

func MoneyFindByMonth(ctx context.Context, username string, year, month, count int) ([]model.Money, error) {
	filter := bson.M{
		"username": username,
		"date": bson.M{
			"$gte": time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC),
			"$lt":  time.Date(year, time.Month(month+count), 1, 0, 0, 0, 0, time.UTC),
		},
	}
	option := options.Find().SetSort(bson.M{"date": -1})
	var money []model.Money
	if err := find(MONEY_COLL, &money, ctx, filter, option); err != nil {
		return money, err
	}
	if len(money) == 0 {
		money = make([]model.Money, 0)
	}
	return money, nil
}

func MoneyFindByTagsIn(ctx context.Context, username string, tags []string) ([]model.Money, error) {
	filter := bson.M{
		"username": username,
		"tags":     bson.M{"$in": tags},
	}
	option := options.Find().SetSort(bson.M{"date": -1})
	var money []model.Money
	if err := find(MONEY_COLL, &money, ctx, filter, option); err != nil {
		return money, err
	}
	if len(money) == 0 {
		money = make([]model.Money, 0)
	}
	return money, nil
}

func MoneyFindByTagsAll(ctx context.Context, username string, tags []string) ([]model.Money, error) {
	filter := bson.M{
		"username": username,
		"tags":     bson.M{"$all": tags},
	}
	option := options.Find().SetSort(bson.M{"date": -1})
	var money []model.Money
	if err := find(MONEY_COLL, &money, ctx, filter, option); err != nil {
		return money, err
	}
	if len(money) == 0 {
		money = make([]model.Money, 0)
	}
	return money, nil
}

func MoneyUpdateOne(ctx context.Context, update *model.Money) (int, error) {
	filter := bson.M{"_id": update.ID}
	result, err := updateOne(MONEY_COLL, update, ctx, filter, nil)
	if err != nil {
		return 0, err
	}
	return int(result.ModifiedCount), nil
}

func MoneyDeleteByID(ctx context.Context, id primitive.ObjectID) (int, error) {
	filter := bson.M{"_id": id}
	result, err := deleteOne(MONEY_COLL, ctx, filter, nil)
	if err != nil {
		return 0, err
	}
	return int(result.DeletedCount), nil
}

func MoneyDeleteByUsername(ctx context.Context, username string) (int, error) {
	filter := bson.M{"username": username}
	result, err := deleteMany(MONEY_COLL, ctx, filter, nil)
	if err != nil {
		return 0, err
	}
	return int(result.DeletedCount), nil
}
