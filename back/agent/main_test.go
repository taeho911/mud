package agent

import (
	"context"
	"fmt"
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
	testID primitive.ObjectID
)

func TestMakeDatabaseURI(t *testing.T) {
	uri := makeDatabaseURI()
	t.Log("uri =", uri)
}

func TestCreateClient(t *testing.T) {
	CreateClient()
	ctx := context.TODO()
	t.Log(makeDatabaseURI())
	if err := client.Connect(ctx); err != nil {
		t.Fatalf("client.Connect(ctx) failed with %s. err = %v", makeDatabaseURI(), err)
	}
	defer client.Disconnect(ctx)
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		t.Fatalf("client.Ping(ctx, readpref.Primary()) failed with %s. err = %v", makeDatabaseURI(), err)
	}
}

func TestInsertOne(t *testing.T) {
	CreateClient()
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
}

func TestFind(t *testing.T) {
	CreateClient()

	var entity []model.Acc
	filter := bson.M{"deleted": false}
	option := options.Find().SetSort(bson.D{primitive.E{Key: "index", Value: 1}})

	if err := find(testcollname, &entity, context.TODO(), filter, option); err != nil {
		t.Fatalf("findAll failed. err = %v", err)
	}
	if sizeOfEntity := len(entity); sizeOfEntity == 0 {
		t.Fatalf("Entity size is %v", sizeOfEntity)
	}
}

func TestUpdateByID(t *testing.T) {
	CreateClient()
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

	var entity model.Acc
	filter := bson.M{"_id": testID}

	if err := findOne(testcollname, &entity, ctx, filter, nil); err != nil {
		t.Fatalf("find failed. err = %v", err)
	}

	printAcc(entity)
}

func TestUpdateOne(t *testing.T) {
	CreateClient()
}

func TestDeleteOne(t *testing.T) {
	CreateClient()

	filter := bson.M{"_id": testID}

	if result, err := deleteOne(testcollname, testID, context.TODO(), filter, nil); err != nil {
		t.Fatalf("deleteOne failed. err = %v", err)
	} else if result.DeletedCount == 0 {
		t.Fatalf("deleteOne failed. DeletedCount = %v", result.DeletedCount)
	}
}

func printAcc(entity model.Acc) {
	fmt.Println("ID:", entity.ID)
	fmt.Println("Owner:", entity.Owner)
	fmt.Println("Title:", entity.Title)
	fmt.Println("Username:", entity.Username)
	fmt.Println("Password:", entity.Password)
	fmt.Println("Location:", entity.Location)
	fmt.Println("Email:", entity.Email)
	fmt.Println("Memo:", entity.Memo)
	fmt.Println("Alias:", entity.Alias)
}
