package internal

import (
	"contactsd/pkg"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ContactDataSource struct {
	Database pkg.IDatabase[*mongo.Client]
	Context  *context.Context
}

func NewMongoDataSource(database pkg.IDatabase[*mongo.Client]) pkg.IDataSource[string, Contact] {
	return ContactDataSource{
		Database: database,
	}
}

func (m ContactDataSource) GetAll(collection string) []Contact {
	return []Contact{}
}

func (m ContactDataSource) GetById(collection string, id string) Contact {
	return Contact{}
}

func (m ContactDataSource) Insert(collection string, contact Contact) {

}

func (m ContactDataSource) Update(collection string, contact Contact) {

}
