package agent

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

const (
	dbname        string        = "mud"
	minpool       uint64        = 5
	maxpool       uint64        = 20
	connidle      time.Duration = 10
	timeoutSecond time.Duration = 5
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

func CreateClient(connUri string) {
	opts := options.Client().ApplyURI(connUri)
	opts.SetMinPoolSize(minpool)
	opts.SetMaxPoolSize(maxpool)
	opts.SetMaxConnIdleTime(connidle)
	opts.SetPoolMonitor(&event.PoolMonitor{
		Event: func(evt *event.PoolEvent) {
			switch evt.Type {
			case event.GetSucceeded:
				log.Println("DB Conn++ :", client.NumberSessionsInProgress())
			case event.ConnectionReturned:
				log.Println("DB Conn-- :", client.NumberSessionsInProgress())
			}
		},
	})
	clientTmp, err := mongo.NewClient(opts)
	if err != nil {
		log.Panicln("Failed to create client. URI:", connUri)
	}
	client = clientTmp
}

func GetColl(ctx context.Context, collname string) (*mongo.Collection, context.Context, context.CancelFunc) {
	dbctx, dbcancel := context.WithTimeout(ctx, timeoutSecond*time.Second)
	err := client.Connect(dbctx)
	if err != nil {
		log.Panicln("Failed to connect to client.")
	}
	return client.Database(dbname).Collection(collname), dbctx, dbcancel
}
