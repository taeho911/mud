package taeho/mud/agent

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDBClient(connUri string) (*mongo.Client, context.Context, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connUri))
	if err != nil {
		
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

}