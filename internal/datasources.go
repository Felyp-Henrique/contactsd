package internal

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CONTACTS_ENTITY_NAME = "contacts"

type MongoDataSource struct {
	Database string
	Client   *mongo.Client
	Context  *context.Context
}

func NewMongoDataSource(database string, context context.Context) MongoDataSource {
	client, err := mongo.Connect(context, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	return MongoDataSource{
		Context:  &context,
		Client:   client,
		Database: database,
	}
}

func (m MongoDataSource) GetAll() []Contact {
	m.Client.Connect(*m.Context)
	defer m.Client.Disconnect(*m.Context)
	collection := m.Client.Database(m.Database).Collection(CONTACTS_ENTITY_NAME)
	cursor, err := collection.Find(*m.Context, bson.D{})
	if err != nil {
		println(err)
		return []Contact{}
	}
	results := []Contact{}
	result := Contact{}
	for cursor.Next(*m.Context) {
		cursor.Decode(&result)
		results = append(results, result)
	}
	return results
}

func (m MongoDataSource) GetById(id string) Contact {
	m.Client.Connect(*m.Context)
	defer m.Client.Disconnect(*m.Context)
	result := Contact{}
	collection := m.Client.Database(m.Database).Collection(CONTACTS_ENTITY_NAME)
	if err := collection.FindOne(*m.Context, bson.M{"id": id}).Decode(&result); err != nil {
		println(err)
		return Contact{}
	}
	return result
}

func (m MongoDataSource) Insert(contact Contact) {
	m.Client.Connect(*m.Context)
	defer m.Client.Disconnect(*m.Context)
	collection := m.Client.Database(m.Database).Collection(CONTACTS_ENTITY_NAME)
	if _, err := collection.InsertOne(*m.Context, &contact); err != nil {
		println(err)
	}
}

func (m MongoDataSource) Update(contact Contact) {
	m.Client.Connect(*m.Context)
	defer m.Client.Disconnect(*m.Context)
	collection := m.Client.Database(m.Database).Collection(CONTACTS_ENTITY_NAME)
	if _, err := collection.UpdateOne(*m.Context, bson.M{"id": contact.Id}, *&contact); err != nil {
		println(err)
	}
}
