package internal

import (
	"context"

	"github.com/facebookgo/inject"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDataSource struct {
	database   string
	collection string
	client     *mongo.Client
	context    *context.Context
}

func (m *MongoDataSource) GetAll() []ContactModel {
	m.client.Connect(*m.context)
	defer m.client.Disconnect(*m.context)
	collection := m.client.Database(m.database).Collection(m.collection)
	cursor, err := collection.Find(*m.context, bson.D{})
	if err != nil {
		println(err)
		return []ContactModel{}
	}
	results := []ContactModel{}
	result := ContactModel{}
	for cursor.Next(*m.context) {
		cursor.Decode(&result)
		results = append(results, result)
	}
	return results
}

func (m *MongoDataSource) GetById(id string) ContactModel {
	m.client.Connect(*m.context)
	defer m.client.Disconnect(*m.context)
	result := ContactModel{}
	collection := m.client.Database(m.database).Collection(m.collection)
	if err := collection.FindOne(*m.context, bson.M{"id": id}).Decode(&result); err != nil {
		println(err)
		return ContactModel{}
	}
	return result
}

func (m *MongoDataSource) Insert(contact ContactModel) {
	m.client.Connect(*m.context)
	defer m.client.Disconnect(*m.context)
	collection := m.client.Database(m.database).Collection(m.collection)
	if _, err := collection.InsertOne(*m.context, &contact); err != nil {
		println(err)
	}
}

func (m *MongoDataSource) Update(contact ContactModel) {
	m.client.Connect(*m.context)
	defer m.client.Disconnect(*m.context)
	collection := m.client.Database(m.database).Collection(m.collection)
	if _, err := collection.UpdateOne(*m.context, bson.M{"id": contact.GetId()}, *&contact); err != nil {
		println(err)
	}
}

func Init(dependencies *inject.Graph) error {
	return dependencies.Provide()
}
