package internal

type ContactModel struct {
	id        string `bson:"id"`
	name      string `bson:"name"`
	email     string `bson:"email"`
	cellphone string `bson:"cellphone"`
}

func (c *ContactModel) GetId() string {
	return c.id
}

func (c *ContactModel) SetId(id string) {
	c.id = id
}

func (c *ContactModel) GetName() string {
	return c.name
}

func (c *ContactModel) SetName(name string) {
	c.name = name
}

func (c *ContactModel) GetEmail() string {
	return c.email
}

func (c *ContactModel) SetEmail(email string) {
	c.email = email
}

func (c *ContactModel) GetCellphone() string {
	return c.cellphone
}

func (c *ContactModel) SetCellphone(cellphone string) {
	c.cellphone = cellphone
}
