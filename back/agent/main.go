package agent

import (
	"context"
	"fmt"
	"os"
	"taeho/mud/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

const (
	dbname   string        = "mud"
	minpool  uint64        = 3
	maxpool  uint64        = 7
	connidle time.Duration = 10
	timeout  time.Duration = 3
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

func CreateClient(ctx context.Context) error {
	options := options.Client().ApplyURI(makeDatabaseURI())
	options.SetMinPoolSize(minpool)
	options.SetMaxPoolSize(maxpool)
	options.SetMaxConnIdleTime(connidle)
	// mongodb의 커넥션이 증가 혹은 감소할 때 실행중인 세션의 수를 출력한다.
	// options.SetPoolMonitor(&event.PoolMonitor{
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
	client, err = mongo.NewClient(options)
	if err != nil {
		return err
	}
	if err := client.Connect(ctx); err != nil {
		return err
	}
	return nil
}

func DeleteClient(ctx context.Context) {
	client.Disconnect(ctx)
}

func getColl(collname string) *mongo.Collection {
	return client.Database(dbname).Collection(collname)
}

func createIndexes(collname string, indexModels []mongo.IndexModel) ([]string, error) {
	name, err := getColl(collname).Indexes().CreateMany(context.TODO(), indexModels, nil)
	if err != nil {
		return nil, err
	}
	return name, nil
}

func insertOne(collname string, entity model.Model, ctx context.Context, option *options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	dbctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := getColl(collname).InsertOne(dbctx, entity, option)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func insertMany(collname string, entity []model.Model, ctx context.Context, option *options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	dbctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := getColl(collname).InsertMany(dbctx, model.ConvertModelToInterface(entity), option)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func findOne(collname string, entity interface{}, ctx context.Context, filter interface{}, option *options.FindOneOptions) error {
	dbctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	var err error
	if option == nil {
		err = getColl(collname).FindOne(dbctx, filter).Decode(entity)
	} else {
		err = getColl(collname).FindOne(dbctx, filter, option).Decode(entity)
	}
	if err != nil {
		return err
	}
	return nil
}

func find(collname string, entity interface{}, ctx context.Context, filter interface{}, option *options.FindOptions) error {
	dbctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

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

func updateByID(collname string, id primitive.ObjectID, update model.Model, ctx context.Context, option *options.UpdateOptions) (*mongo.UpdateResult, error) {
	dbctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	result, err := getColl(collname).UpdateByID(dbctx, id, bson.M{"$set": update}, option)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func updateOne(collname string, update model.Model, ctx context.Context, filter interface{}, option *options.UpdateOptions) (*mongo.UpdateResult, error) {
	dbctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	result, err := getColl(collname).UpdateOne(dbctx, filter, bson.M{"$set": update}, option)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func deleteOne(collname string, ctx context.Context, filter interface{}, option *options.DeleteOptions) (*mongo.DeleteResult, error) {
	dbctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	result, err := getColl(collname).DeleteOne(dbctx, filter, option)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func deleteMany(collname string, ctx context.Context, filter interface{}, option *options.DeleteOptions) (*mongo.DeleteResult, error) {
	dbctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	result, err := getColl(collname).DeleteMany(dbctx, filter, option)
	if err != nil {
		return nil, err
	}
	return result, nil
}
