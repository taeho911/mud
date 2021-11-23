package agent

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

const (
	dbname   string        = "mud"
	minpool  uint64        = 5
	maxpool  uint64        = 20
	connidle time.Duration = 10
)

func makeDatabaseURI() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "27017"
	}
	if username != "" && password != "" {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", username, password, host, port)
	} else {
		return fmt.Sprintf("mongodb://%s:%s/", host, port)
	}
}

func CreateClient() {
	connUri := makeDatabaseURI()
	opts := options.Client().ApplyURI(makeDatabaseURI())
	opts.SetMinPoolSize(minpool)
	opts.SetMaxPoolSize(maxpool)
	opts.SetMaxConnIdleTime(connidle)
	// opts.SetPoolMonitor(&event.PoolMonitor{
	// 	Event: func(evt *event.PoolEvent) {
	// 		switch evt.Type {
	// 		case event.GetSucceeded:
	// 			log.Println("DB Conn++ :", client.NumberSessionsInProgress())
	// 		case event.ConnectionReturned:
	// 			log.Println("DB Conn-- :", client.NumberSessionsInProgress())
	// 		}
	// 	},
	// })
	var err error
	client, err = mongo.NewClient(opts)
	if err != nil {
		log.Panicln("Failed to create client. URI =", connUri)
	}
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := client.Connect(ctx); err != nil {
	// 	log.Panicln("Failed to connect to database. URI =", connUri)
	// }
	// defer client.Disconnect(ctx)
	// if err := client.Ping(ctx, readpref.Primary()); err != nil {
	// 	log.Panicf("Failed to ping database. URI = %s, err = %v", connUri, err)
	// }
}

// func disconnClient(dbctx context.Context, cancel context.CancelFunc) {
// 	client.Disconnect(dbctx)
// 	cancel()
// }

func getColl(collname string) *mongo.Collection {
	return client.Database(dbname).Collection(collname)
}

func insertOne(collname string, entity interface{}, ctx context.Context) (*mongo.InsertOneResult, error) {
	dbctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := client.Connect(dbctx); err != nil {
		return nil, err
	}
	defer client.Disconnect(dbctx)

	result, err := getColl(collname).InsertOne(dbctx, entity)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func findAll(collname string, entity interface{}, ctx context.Context, filter interface{}, option *options.FindOptions) error {
	dbctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := client.Connect(dbctx); err != nil {
		return err
	}
	defer client.Disconnect(dbctx)

	curosr, err := getColl(collname).Find(dbctx, filter, option)
	if err != nil {
		return err
	}
	defer curosr.Close(dbctx)
	if err := curosr.All(dbctx, entity); err != nil {
		return err
	}
	return nil
}
