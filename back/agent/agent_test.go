package agent

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestCreateClient(t *testing.T) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	connUri := fmt.Sprintf("mongodb://%s:%s@localhost:27017/?authSource=admin", username, password)
	CreateClient(connUri)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// if err := client.Connect(ctx); err != nil {
	// 	t.Fatalf("client.Connect(ctx) failed with %s. err = %v", connUri, err)
	// }
	// defer client.Disconnect(ctx)
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		t.Fatalf("client.Ping(ctx, readpref.Primary()) failed with %s. err = %v", connUri, err)
	}
}

func TestGetColl(t *testing.T) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	connUri := fmt.Sprintf("mongodb://%s:%s@localhost:27017/?authSource=admin", username, password)
	CreateClient(connUri)
	coll, dbctx, dbcancel := GetColl(context.Background(), "test")
	if coll == nil {
		t.Fatalf("coll = %v", coll)
	}
	if dbctx == nil {
		t.Fatalf("dbctx = %v", dbctx)
	}
	if dbcancel == nil {
		t.Fatalf("dbcancel = %v", dbcancel)
	}
	// fmt.Printf("Type: coll = %T, dbctx = %T, dbcancel = %T\n", coll, dbctx, dbcancel)
	// fmt.Printf("Value: coll = %v, dbctx = %v, dbcancel = %v\n", coll, dbctx, dbcancel)
	// var collwant *mongo.Collection = nil
	// fmt.Println(reflect.TypeOf(collwant))
	// if reflect.TypeOf(coll) != reflect.TypeOf(*mongo.Collection) {
	// 	t.Fatalf("coll type = %T. want = %T", coll, *mongo.Collection)
	// }
}
