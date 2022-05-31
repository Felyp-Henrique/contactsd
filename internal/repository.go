package internal

import (
	"contactsd/pkg"

	"github.com/google/uuid"
)

type ContactRepository struct {
	Datasource pkg.IDataSource[string, Contact]
}

func NewContactRepositoryMongo(datasource MongoDataSource) ContactRepository {
	return ContactRepository{
		Datasource: datasource,
	}
}

func (c *ContactRepository) GetAll() []Contact {
	return c.GetAll()
}

func (c *ContactRepository) GetById(id string) Contact {
	return c.GetById(id)
}

func (c *ContactRepository) Insert(contact Contact) {
	contact.Id = uuid.NewString()
	c.Insert(contact)
}

func (c *ContactRepository) Update(contact Contact) {
	c.Update(contact)
}
