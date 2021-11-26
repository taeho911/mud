package agent

import (
	"context"
	"taeho/mud/model"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	testcollname string = "test"
)

var (
	testID     primitive.ObjectID
	testEntity model.Acc
)

func TestMakeDatabaseURI(t *testing.T) {
	uri := makeDatabaseURI()
	t.Log("uri =", uri)
}

func TestCreateClient(t *testing.T) {
	ctx := context.TODO()
	if err := CreateClient(ctx); err != nil {
		t.Fatalf("CreateClient(ctx) failed with %s. err = %v", makeDatabaseURI(), err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		t.Fatalf("client.Ping(ctx, readpref.Primary()) failed with %s. err = %v", makeDatabaseURI(), err)
	}
}

func TestInsertOne(t *testing.T) {
	entity := model.Acc{
		Owner:    "testuser",
		Index:    0,
		Title:    "testtitle",
		Username: "testuser",
		Password: "testpass",
	}

	result, err := insertOne(testcollname, entity, context.TODO())
	if err != nil {
		t.Fatalf("insertOne failed. err = %v", err)
	}
	if result.InsertedID == nil {
		t.Fatalf("insertOne failed. InsertedID = %v", result.InsertedID)
	}
	testID = result.InsertedID.(primitive.ObjectID)
	testEntity = entity
}

func TestFindOne(t *testing.T) {
	entity := model.Acc{}
	filter := bson.M{"_id": testID}

	t.Log("entity =", entity)

	if err := findOne(testcollname, &entity, context.TODO(), filter, nil); err != nil {
		t.Fatalf("findOne failed. err = %v", err)
	}
	if entity.Owner != testEntity.Owner {
		t.Fatalf("entity.Owner is wrong. Owner = %v", entity.Owner)
	}
	if entity.Title != testEntity.Title {
		t.Fatalf("entity.Title is wrong. Title = %v", entity.Title)
	}
}

func TestFind(t *testing.T) {
	entity := []model.Acc{}
	filter := bson.M{"deleted": false}
	option := options.Find().SetSort(bson.D{primitive.E{Key: "index", Value: 1}})

	if err := find(testcollname, &entity, context.TODO(), filter, option); err != nil {
		t.Fatalf("find failed. err = %v", err)
	}
	if sizeOfEntity := len(entity); sizeOfEntity == 0 {
		t.Fatalf("Entity size is %v", sizeOfEntity)
	}
}

func TestUpdateByID(t *testing.T) {
	ctx := context.TODO()
	update := model.Acc{
		Username: "updatebyid",
		Password: "haha",
		Email:    "UpdateByID@gmail.com",
	}

	if result, err := updateByID(testcollname, testID, update, ctx, nil); err != nil {
		t.Fatalf("updateByID failed. err = %v", err)
	} else if result.ModifiedCount == 0 {
		t.Fatalf("updateByID failed. ModifiedCount = %v", result.ModifiedCount)
	}

	entity := model.Acc{}
	filter := bson.M{"_id": testID}

	if err := findOne(testcollname, &entity, ctx, filter, nil); err != nil {
		t.Fatalf("find failed. err = %v", err)
	}
}

func TestUpdateOne(t *testing.T) {
	ctx := context.TODO()
	update := model.Acc{
		Title:    "Hello world",
		Username: "UpdateOne",
		Password: "---",
		Email:    "UpdateOne@naver.com",
	}
	filter := bson.M{"_id": testID}

	if result, err := updateOne(testcollname, update, ctx, filter, nil); err != nil {
		t.Fatalf("updateOne failed. err = %v", err)
	} else if result.ModifiedCount == 0 {
		t.Fatalf("updateOne failed. ModifiedCount = %v", result.ModifiedCount)
	}
}

func TestDeleteOne(t *testing.T) {
	filter := bson.M{"_id": testID}

	if result, err := deleteOne(testcollname, testID, context.TODO(), filter, nil); err != nil {
		t.Fatalf("deleteOne failed. err = %v", err)
	} else if result.DeletedCount == 0 {
		t.Fatalf("deleteOne failed. DeletedCount = %v", result.DeletedCount)
	}
}
