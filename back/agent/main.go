package agent

import (
	"context"
	"log"
	"time"

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

func CreateClient(connUri string) {
	opts := options.Client().ApplyURI(connUri)
	opts.SetMinPoolSize(minpool)
	opts.SetMaxPoolSize(maxpool)
	opts.SetMaxConnIdleTime(connidle)
	clientTmp, err := mongo.NewClient()
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
