package internal

import (
	"contactsd/pkg"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InjectConfigurations(injection *pkg.Injection) {
	injection.Put("database", "contacts")
	injection.Put("collection", "contacts")
	injection.Put("context", context.TODO())
}

func InjectionExternal(injection *pkg.Injection) {
	injection.Put("client", func() *mongo.Client {
		options := options.Client().ApplyURI("mongodb://localhost:27017")
		client, err := mongo.Connect(injection.Get("context").(context.Context), options)
		if err != nil {
			panic(err)
		}
		return client
	})
}

func InjectDataSources(injection *pkg.Injection) {
	injection.Put("datasource", &MongoDataSource{
		database:   injection.Get("database").(string),
		collection: injection.Get("collection").(string),
		context:    injection.Get("context").(*context.Context),
		client:     injection.Get("client").(func() *mongo.Client)(),
	})
}

func InjectRepositories(injection *pkg.Injection) {
	injection.Put("repository", &ContactRepository{
		datasource: injection.Get("datasource").(*MongoDataSource),
	})
}
