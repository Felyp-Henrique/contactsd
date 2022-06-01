package pkg

type IContactRepository[I any, M any] interface {
	GetAll() []M
	GetById(id I) M
	Insert(contact M)
	Update(contact M)
}
