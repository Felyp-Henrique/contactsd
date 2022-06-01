package pkg

type IDatabase[D any] interface {
	Connect() (D, error)
	Disconnect(D) error
}
