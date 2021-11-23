package agent

import (
	"context"
	"taeho/mud/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collname string = "acc"
)

func AccFindAll(ctx context.Context) ([]model.Acc, error) {
	filter := bson.M{"deleted": false}
	// mongodb의 소트에는 1(asc), -1(desc)이 존재한다.
	option := options.Find().SetSort(bson.D{primitive.E{Key: "index", Value: 1}})
	var entity []model.Acc
	if err := findAll(collname, &entity, ctx, filter, option); err != nil {
		return nil, err
	}
	return entity, nil
}
