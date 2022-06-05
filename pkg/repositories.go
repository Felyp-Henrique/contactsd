package pkg

type IContactRepository[I any, M any] interface {
	GetAll() ([]M, error)
	GetById(id I) (M, error)
	Insert(contact M) error
	Update(contact M) error
}
