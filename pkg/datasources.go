package pkg

type IDataSource[I any, M any] interface {
	GetAll() ([]M, error)
	GetById(I) (M, error)
	Insert(M) error
	Update(M) error
}
