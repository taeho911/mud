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
	UserCreateIndexes()
	AccountCreateIndexes()
	returnCode := m.Run()
	os.Exit(returnCode)
}

func TestCreateIndexes(t *testing.T) {
	var account model.Account
	_, err := createIndexes("account", account.IndexFields())
	if err != nil {
		t.Fatalf("createIndexes failed. err = %v", err)
	}
}

func TestCheckNotNullFields(t *testing.T) {
	entity := model.Test{
		Dummy: "TestCheckNotNullFields",
	}
	err := checkNotNullFields(&entity)
	if err == nil {
		t.Fatal("checkNotNullFields failed.")
	}
}

func TestInsertOne(t *testing.T) {
	collname := "account"
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

	entity_2 := model.Account{
		Owner:    "testuser",
		Password: "testpass",
	}

	if _, err := insertOne(collname, &entity_2, ctx, nil); err == nil {
		t.Fatalf("insertOne failed.")
	}
}

func TestInsertMany(t *testing.T) {
	collname := "account"
	ctx := context.TODO()
	insertEntity := []model.Model{
		&model.Account{
			Owner:    "TestInsertMany",
			Title:    "TestInsertMany_1",
			Username: "TestInsertMany_1",
			Password: "TestInsertMany_1",
		},
		&model.Account{
			Owner:    "TestInsertMany",
			Title:    "TestInsertMany_2",
			Username: "TestInsertMany_2",
			Password: "TestInsertMany_2",
		},
	}

	result, err := insertMany(collname, insertEntity, ctx, nil)
	if err != nil {
		t.Fatalf("InsertMany failed. err = %v", err)
	}
	if sizeOfResult := len(result.InsertedIDs); sizeOfResult != 2 {
		t.Fatalf("InsertMany failed. sizeOfResult = %v", sizeOfResult)
	}

	for _, v := range result.InsertedIDs {
		deleteOne(collname, ctx, bson.M{"_id": v}, nil)
	}
}

func TestFindOne(t *testing.T) {
	collname := "account"
	ctx := context.TODO()
	insertEntity := model.Account{
		Owner:    "TestFindOne",
		Title:    "TestFindOne",
		Username: "TestFindOne",
		Password: "TestFindOne",
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
	collname := "account"
	ctx := context.TODO()
	insertEntity := []model.Model{
		&model.Account{
			Owner:    "TestFind",
			Title:    "TestFind_1",
			Username: "TestFind_1",
			Password: "TestFind_1",
		},
		&model.Account{
			Owner:    "TestFind",
			Title:    "TestFind_2",
			Username: "TestFind_2",
			Password: "TestFind_2",
		},
	}
	result, _ := insertMany(collname, insertEntity, ctx, nil)

	entity := []model.Account{}
	filter := bson.M{"owner": "TestFind"}
	option := options.Find().SetSort(bson.D{primitive.E{Key: "index", Value: 1}})

	if err := find(collname, &entity, ctx, filter, option); err != nil {
		t.Fatalf("find failed. err = %v", err)
	}
	if sizeOfEntity := len(entity); sizeOfEntity < 2 {
		t.Fatalf("Entity size is %v", sizeOfEntity)
	}

	for _, v := range result.InsertedIDs {
		deleteOne(collname, ctx, bson.M{"_id": v}, nil)
	}
}

func TestUpdateByID(t *testing.T) {
	collname := "account"
	ctx := context.TODO()
	insertEntity := model.Account{
		Owner:    "TestUpdateByID",
		Title:    "TestUpdateByID",
		Username: "TestUpdateByID",
		Password: "TestUpdateByID",
		Email:    "TestUpdateByID@haha.com",
		Alias:    []string{"1", "2"},
	}
	result, _ := insertOne(collname, &insertEntity, ctx, nil)

	update := model.Account{
		Owner:    "updatebyid",
		Title:    "updatebyid",
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
	collname := "account"
	ctx := context.TODO()
	insertEntity := model.Account{
		Owner:    "TestUpdateOne",
		Title:    "TestUpdateOne",
		Username: "TestUpdateOne",
		Password: "TestUpdateOne",
		Email:    "TestUpdateOne@hoho.com",
		Memo:     "ipsum rorem",
	}
	result, _ := insertOne(collname, &insertEntity, ctx, nil)

	update := model.Account{
		Owner:    "foo",
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

func TestDeleteMany(t *testing.T) {
	collname := "account"
	ctx := context.TODO()
	insertEntity := []model.Model{
		&model.Account{
			Owner:    "TestDeleteMany",
			Title:    "TestDeleteMany_1",
			Username: "TestDeleteMany_1",
			Password: "TestDeleteMany_1",
		},
		&model.Account{
			Owner:    "TestDeleteMany",
			Title:    "TestDeleteMany_2",
			Username: "TestDeleteMany_2",
			Password: "TestDeleteMany_2",
		},
	}
	insertMany(collname, insertEntity, ctx, nil)

	if result, err := deleteMany(collname, ctx, bson.M{"owner": "TestDeleteMany"}, nil); err != nil {
		t.Fatalf("deleteMany failed. err = %v", err)
	} else if result.DeletedCount != 2 {
		t.Fatalf("deleteMany failed. result.DeletedCount = %v", result.DeletedCount)
	}
}
