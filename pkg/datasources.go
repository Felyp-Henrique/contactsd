package pkg

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type IDataSource[I any, M any] interface {
	GetAll(string) []M
	GetById(string, id I) M
	Insert(string, M)
	Update(string, M)
}

type MongoDataSource struct {
	Database IDatabase[*mongo.Client]
	Context  *context.Context
}

func NewMongoDataSource(database IDatabase[*mongo.Client]) MongoDataSource {
	return MongoDataSource{
		Database: database,
	}
}

func (m MongoDataSource) GetAll(collection string) []map[string]interface{} {
	return []map[string]interface{}{}
}

func (m MongoDataSource) GetById(collection string, id string) map[string]interface{} {
	return map[string]interface{}{}
}

func (m MongoDataSource) Insert(collection string, data map[string]interface{}) {

}

func (m MongoDataSource) Update(collection string, data map[string]interface{}) {

}
