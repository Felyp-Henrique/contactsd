package internal

import (
	"contactsd/pkg"

	"github.com/google/uuid"
)

type ContactsRepository struct {
	Datasource pkg.MongoDataSource
}

func NewContactsRepositoryMongo(datasource pkg.MongoDataSource) ContactsRepository {
	return ContactsRepository{
		Datasource: datasource,
	}
}

func (c *ContactsRepository) GetAll() []Contact {
	return []Contact{}
}

func (c *ContactsRepository) GetById(id string) Contact {
	return Contact{}
}

func (c *ContactsRepository) Insert(contact Contact) {
	contact.Id = uuid.NewString()
}

func (c *ContactsRepository) Update(contact Contact) {

}
