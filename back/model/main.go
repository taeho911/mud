package model

import "go.mongodb.org/mongo-driver/mongo"

type Model interface {
	NotNullFields() []interface{}
	IndexFields() []mongo.IndexModel
	SetMaketime()
}

func ConvertModelToInterface(entity []Model) []interface{} {
	interfaces := make([]interface{}, len(entity))
	for i, v := range entity {
		v.SetMaketime()
		interfaces[i] = v
	}
	return interfaces
}

// How to create indexes
// https://kb.objectrocket.com/mongo-db/how-to-create-an-index-using-the-golang-driver-for-mongodb-455
// https://stackoverflow.com/questions/56759074/how-do-i-create-a-text-index-in-mongodb-with-golang-and-the-mongo-go-driver
// https://docs.mongodb.com/manual/indexes/
