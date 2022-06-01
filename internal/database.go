package internal

import (
	"contactsd/pkg"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	Database string
	Username string
	Password string
	Context  context.Context
}

func NewMongoDatabase(database string, context context.Context) pkg.IDatabase[*mongo.Client] {
	return pkg.IDatabase[*mongo.Client](MongoDatabase{
		Database: database,
		Context:  context,
	})
}

func (m MongoDatabase) Connect() (*mongo.Client, error) {
	auth := options.Credential{
		Username:   m.Username,
		Password:   m.Password,
		AuthSource: "admin",
	}
	options := options.
		Client().
		ApplyURI("mongodb://localhost:27017").
		SetAuth(auth)
	return mongo.Connect(m.Context, options)
}

func (m MongoDatabase) Disconnect(driver *mongo.Client) error {
	return driver.Disconnect(m.Context)
}
