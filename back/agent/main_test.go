package agent

import (
	"context"
	"taeho/mud/model"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	testcollname string = "test"
)

func TestMakeDatabaseURI(t *testing.T) {
	uri := makeDatabaseURI()
	t.Log("uri =", uri)
}

func TestCreateClient(t *testing.T) {
	CreateClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
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
		Owner:    "testuser1",
		Index:    0,
		Title:    "testtitle",
		Username: "testuser",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := insertOne(testcollname, entity, ctx)
	if err != nil {
		t.Fatalf("insertOne(testcollname, entity, ctx) failed. err = %v", err)
	}
	if result.InsertedID == nil {
		t.Fatalf("insertOne(testcollname, entity, ctx) failed. InsertedID = %v", result.InsertedID)
	}
}

func TestFindAll(t *testing.T) {
	CreateClient()
	var entity []model.Acc
	filter := bson.M{"deleted": false}
	option := options.Find().SetSort(bson.D{primitive.E{Key: "index", Value: 1}})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := findAll(testcollname, &entity, ctx, filter, option); err != nil {
		t.Fatalf("findAll(testcollname, &entity, ctx, filter, option) failed. err = %v", err)
	}
	if sizeOfEntity := len(entity); sizeOfEntity == 0 {
		t.Fatalf("Entity size is %v", sizeOfEntity)
	}
}

// func TestGetColl(t *testing.T) {
// 	connUri := makeDatabaseURI()
// 	CreateClient(connUri)
// 	coll, dbctx, dbcancel := GetColl(context.Background(), "test")
// 	defer client.Disconnect(dbctx)
// 	if coll == nil {
// 		t.Fatalf("coll = %v", coll)
// 	}
// 	if dbctx == nil {
// 		t.Fatalf("dbctx = %v", dbctx)
// 	}
// 	if dbcancel == nil {
// 		t.Fatalf("dbcancel = %v", dbcancel)
// 	}
// 	fmt.Printf("Type: coll = %T, dbctx = %T, dbcancel = %T\n", coll, dbctx, dbcancel)
// 	fmt.Printf("Value: coll = %v, dbctx = %v, dbcancel = %v\n", coll, dbctx, dbcancel)
// 	var collwant *mongo.Collection = nil
// 	fmt.Println(reflect.TypeOf(collwant))
// 	if reflect.TypeOf(coll) != reflect.TypeOf(*mongo.Collection) {
// 		t.Fatalf("coll type = %T. want = %T", coll, *mongo.Collection)
// 	}
// }
