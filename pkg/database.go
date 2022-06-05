package pkg

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(context context.Context, database string) (*mongo.Database, error) {
	option := options.Client().ApplyURI("mongodb://localhost:27017")
	if client, err := mongo.Connect(context, option); err != nil {
		return nil, err
	} else {
		return client.Database(database), ErrorDatabaseConnectionUnavailable
	}
}
