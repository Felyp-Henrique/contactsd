package internal

import (
	"contactsd/pkg"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContactsMongoDataSource struct {
	Database *mongo.Database
	Context  context.Context
}

func NewContactsMongoDataSource(context context.Context, database *mongo.Database) ContactsMongoDataSource {
	return ContactsMongoDataSource{
		Database: database,
		Context:  context,
	}
}
func (c *ContactsMongoDataSource) GetAll() ([]pkg.Contact, error) {
	query := c.Database.Collection("contacts")
	contacts := []pkg.Contact{}
	contact := &pkg.Contact{}
	if cursor, err := query.Find(c.Context, bson.D{}); err == nil {
		for cursor.Next(c.Context) {
			cursor.Decode(contact)
			contacts = append(contacts, *contact)
		}
	} else {
		return contacts, pkg.ErrorContactsEmpty
	}
	return contacts, nil
}

func (c *ContactsMongoDataSource) GetById(id string) (*pkg.Contact, error) {
	query := c.Database.Collection("contacts")
	contact := pkg.Contact{}
	if cursor, err := query.Find(c.Context, bson.M{"id": id}); err == nil {
		if cursor.Next(c.Context) {
			cursor.Decode(&contact)
		} else {
			return nil, pkg.ErrorContactsNotFound
		}
	} else {
		return nil, pkg.ErrorContactsNotFound
	}
	return &contact, nil
}

func (c *ContactsMongoDataSource) Insert(contact pkg.Contact) (*pkg.Contact, error) {
	contact.Id = uuid.NewString()
	query := c.Database.Collection("contacts")
	if _, err := query.InsertOne(c.Context, &contact); err != nil {
		return nil, pkg.ErrorContactsInsert
	}
	return c.GetById(contact.Id)
}

func (c *ContactsMongoDataSource) Update(contact pkg.Contact) error {
	query := c.Database.Collection("contacts")
	if _, err := query.UpdateOne(c.Context, bson.M{"id": contact.Id}, contact); err != nil {
		return pkg.ErrorContactsUpdate
	}
	return nil
}
