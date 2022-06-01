package internal

import (
	"contactsd/pkg"

	"github.com/google/uuid"
)

type ContactRepository struct {
	Datasource pkg.IDataSource[string, Contact]
}

func NewContactRepositoryMongo(datasource pkg.IDataSource[string, Contact]) ContactRepository {
	return ContactRepository{
		Datasource: datasource,
	}
}

func (c *ContactRepository) GetAll() []Contact {
	return c.Datasource.GetAll("contacts")
}

func (c *ContactRepository) GetById(id string) Contact {
	return c.Datasource.GetById("contacts", id)
}

func (c *ContactRepository) Insert(contact Contact) {
	contact.Id = uuid.NewString()
	c.Datasource.Insert("contacts", contact)
}

func (c *ContactRepository) Update(contact Contact) {
	c.Datasource.Update("contacts", contact)
}
