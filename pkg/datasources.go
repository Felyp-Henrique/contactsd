package pkg

type IDataSource[I any, M any] interface {
	GetAll(string) []M
	GetById(string, id I) M
	Insert(string, M)
	Update(string, M)
}
