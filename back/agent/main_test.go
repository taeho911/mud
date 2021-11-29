package agent

import (
	"context"
	"fmt"
	"os"
	"taeho/mud/model"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMain(m *testing.M) {
	ctx := context.TODO()
	if err := CreateClient(ctx); err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(-1)
	}
	defer DeleteClient(ctx)
	returnCode := m.Run()
	os.Exit(returnCode)
}

// func TestMakeDatabaseURI(t *testing.T) {
// 	uri := makeDatabaseURI()
// 	t.Log("uri =", uri)
// }

// func TestCreateClient(t *testing.T) {
// 	ctx := context.TODO()
// 	if err := CreateClient(ctx); err != nil {
// 		t.Fatalf("CreateClient(ctx) failed with %s. err = %v", makeDatabaseURI(), err)
// 	}
// 	if err := client.Ping(ctx, readpref.Primary()); err != nil {
// 		t.Fatalf("client.Ping(ctx, readpref.Primary()) failed with %s. err = %v", makeDatabaseURI(), err)
// 	}
// }

func TestCreateIndexes(t *testing.T) {
	var account model.Account
	name, err := createIndexes("test", account.IndexFields())
	if err != nil {
		t.Fatalf("createIndexes failed. err = %v", err)
	}
	fmt.Println("name =", name)
}

func TestInsertOne(t *testing.T) {
	collname := "test"
	ctx := context.TODO()
	entity := model.Account{
		Owner:    "testuser",
		Title:    "testtitle",
		Username: "testuser",
		Password: "testpass",
	}

	result, err := insertOne(collname, &entity, ctx, nil)
	if err != nil {
		t.Fatalf("insertOne failed. err = %v", err)
	}
	if result.InsertedID == nil {
		t.Fatalf("insertOne failed. InsertedID = %v", result.InsertedID)
	}

	deleteOne(collname, ctx, bson.M{"_id": result.InsertedID}, nil)
}

func TestInsertMany(t *testing.T) {
	collname := "test"
	ctx := context.TODO()
	insertEntity := []model.Model{
		&model.Account{
			Title:    "TestInsertMany",
			Username: "TestInsertMany_1",
		},
		&model.Account{
			Title:    "TestInsertMany",
			Username: "TestInsertMany_2",
		},
	}

	result, err := insertMany(collname, insertEntity, ctx, nil)
	if err != nil {
		t.Fatalf("InsertMany failed. err = %v", err)
	}
	if sizeOfResult := len(result.InsertedIDs); sizeOfResult != 2 {
		t.Fatalf("InsertMany failed. sizeOfResult = %v", sizeOfResult)
	}

	for item := range result.InsertedIDs {
		deleteOne(collname, ctx, bson.M{"_id": item}, nil)
	}
}

func TestFindOne(t *testing.T) {
	collname := "test"
	ctx := context.TODO()
	insertEntity := model.Account{
		Title: "TestFindOne",
	}
	result, _ := insertOne(collname, &insertEntity, ctx, nil)

	var entity model.Account
	filter := bson.M{"_id": result.InsertedID}

	if err := findOne(collname, &entity, ctx, filter, nil); err != nil {
		t.Fatalf("findOne failed. err = %v", err)
	}
	if entity.Title != insertEntity.Title {
		t.Fatalf("entity.Title is wrong. Title = %v", entity.Title)
	}

	deleteOne(collname, ctx, bson.M{"_id": result.InsertedID}, nil)
}

func TestFind(t *testing.T) {
	collname := "test"
	ctx := context.TODO()
	insertEntity := []model.Model{
		&model.Account{
			Title:    "TestFind",
			Username: "TestFind_1",
		},
		&model.Account{
			Title:    "TestFind",
			Username: "TestFind_2",
		},
	}
	result, _ := insertMany(collname, insertEntity, ctx, nil)

	entity := []model.Account{}
	filter := bson.M{"title": "TestFind"}
	option := options.Find().SetSort(bson.D{primitive.E{Key: "index", Value: 1}})

	if err := find(collname, &entity, ctx, filter, option); err != nil {
		t.Fatalf("find failed. err = %v", err)
	}
	if sizeOfEntity := len(entity); sizeOfEntity < 2 {
		t.Fatalf("Entity size is %v", sizeOfEntity)
	}

	for item := range result.InsertedIDs {
		deleteOne(collname, ctx, bson.M{"_id": item}, nil)
	}
}

func TestUpdateByID(t *testing.T) {
	collname := "test"
	ctx := context.TODO()
	insertEntity := model.Account{
		Username: "TestUpdateByID",
		Email:    "TestUpdateByID@haha.com",
		Alias:    []string{"1", "2"},
	}
	result, _ := insertOne(collname, &insertEntity, ctx, nil)

	update := model.Account{
		Username: "updatebyid",
		Password: "haha",
		Email:    "UpdateByID@gmail.com",
	}

	if result, err := updateByID(collname, result.InsertedID.(primitive.ObjectID), &update, ctx, nil); err != nil {
		t.Fatalf("updateByID failed. err = %v", err)
	} else if result.ModifiedCount == 0 {
		t.Fatalf("updateByID failed. ModifiedCount = %v", result.ModifiedCount)
	}

	deleteOne(collname, ctx, bson.M{"_id": result.InsertedID}, nil)
}

func TestUpdateOne(t *testing.T) {
	collname := "test"
	ctx := context.TODO()
	insertEntity := model.Account{
		Username: "TestUpdateOne",
		Email:    "TestUpdateOne@hoho.com",
		Memo:     "ipsum rorem",
	}
	result, _ := insertOne(collname, &insertEntity, ctx, nil)

	update := model.Account{
		Title:    "Hello world",
		Username: "UpdateOne",
		Password: "---",
		Email:    "UpdateOne@naver.com",
	}
	filter := bson.M{"_id": result.InsertedID}

	if result, err := updateOne(collname, &update, ctx, filter, nil); err != nil {
		t.Fatalf("updateOne failed. err = %v", err)
	} else if result.ModifiedCount == 0 {
		t.Fatalf("updateOne failed. ModifiedCount = %v", result.ModifiedCount)
	}

	deleteOne(collname, ctx, bson.M{"_id": result.InsertedID}, nil)
}

// func TestDeleteOne(t *testing.T) {
// 	filter := bson.M{"_id": testID}

// 	if result, err := deleteOne(testcollname, context.TODO(), filter, nil); err != nil {
// 		t.Fatalf("deleteOne failed. err = %v", err)
// 	} else if result.DeletedCount == 0 {
// 		t.Fatalf("deleteOne failed. DeletedCount = %v", result.DeletedCount)
// 	}
// }
