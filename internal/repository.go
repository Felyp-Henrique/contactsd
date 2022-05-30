package internal

import (
	"contactsd/pkg"

	"github.com/google/uuid"
)

type ContactRepository struct {
	datasource pkg.IDataSourceNoSQL[string, ContactModel]
}

func (c *ContactRepository) GetAll() []ContactModel {
	return c.GetAll()
}

func (c *ContactRepository) GetById(id string) ContactModel {
	return c.GetById(id)
}

func (c *ContactRepository) Insert(contact ContactModel) {
	contact.SetId(uuid.NewString())
	c.Insert(contact)
}

func (c *ContactRepository) Update(contact ContactModel) {
	c.Update(contact)
}
