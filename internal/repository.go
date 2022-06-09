package internal

import (
	"contactsd/pkg"
)

type ContactsRepository struct {
	Datasource ContactsMongoDataSource
}

func NewContactsRepositoryMongo(datasource ContactsMongoDataSource) ContactsRepository {
	return ContactsRepository{
		Datasource: datasource,
	}
}

func (c *ContactsRepository) GetAll() ([]pkg.Contact, error) {
	return c.Datasource.GetAll()
}

func (c *ContactsRepository) GetById(id string) (*pkg.Contact, error) {
	return c.Datasource.GetById(id)
}

func (c *ContactsRepository) Insert(contact pkg.Contact) (*pkg.Contact, error) {
	return c.Datasource.Insert(contact)
}

func (c *ContactsRepository) Update(contact pkg.Contact) error {
	return c.Datasource.Update(contact)
}
