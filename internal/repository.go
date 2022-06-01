package internal

import (
	"contactsd/pkg"

	"github.com/google/uuid"
)

type ContactsRepository struct {
	Datasource pkg.IDataSource[string, Contact]
}

func NewContactsRepositoryMongo(datasource pkg.IDataSource[string, Contact]) ContactsRepository {
	return ContactsRepository{
		Datasource: datasource,
	}
}

func (c *ContactsRepository) GetAll() []Contact {
	return c.Datasource.GetAll("contacts")
}

func (c *ContactsRepository) GetById(id string) Contact {
	return c.Datasource.GetById("contacts", id)
}

func (c *ContactsRepository) Insert(contact Contact) {
	contact.Id = uuid.NewString()
	c.Datasource.Insert("contacts", contact)
}

func (c *ContactsRepository) Update(contact Contact) {
	c.Datasource.Update("contacts", contact)
}
