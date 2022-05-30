package pkg

type IContact[I any] interface {
	GetId() I
	SetId(id I)
	GetName() string
	SetName(name string)
	GetEmail() string
	SetEmail(email string)
	GetCellphone() string
	SetCellphone(cellphone string)
}
