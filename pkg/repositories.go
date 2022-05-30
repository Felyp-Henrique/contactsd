package pkg

type IContactRepository[I any, D any, M IContact[I]] interface {
	GetAll() []IContact[I]
	GetById(id I) IContact[I]
	Insert(contact IContact[I])
	Update(contact IContact[I])
}
