package pkg

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(context context.Context, database string) (*mongo.Database, error) {
	credential := options.Credential{
		Username: "root",
		Password: "toor",
	}
	option := options.
		Client().
		ApplyURI("mongodb://localhost:27017").
		SetAuth(credential)
	if client, err := mongo.Connect(context, option); err != nil {
		return nil, ErrorDatabaseConnectionUnavailable
	} else {
		return client.Database(database), nil
	}
}
