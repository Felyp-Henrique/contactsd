package pkg

type Response[D any] struct {
	Status int      `json:"status"`
	Data   D        `json:"data"`
	Erros  []string `json:"erros"`
}

func NewResponseOk[D any](data D) *Response[D] {
	return &Response[D]{
		Data:   data,
		Status: 200,
		Erros:  []string{},
	}
}
