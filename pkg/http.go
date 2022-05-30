package pkg

type Response[D any] struct {
	Status int
	Data   D
	Erros  []string
}

func NewResponseOk[D any](data D) *Response[D] {
	return &Response[D]{
		Data:   data,
		Status: 200,
		Erros:  []string{},
	}
}
