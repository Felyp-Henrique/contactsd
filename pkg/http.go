package pkg

import "net/http"

type Response[D any] struct {
	Status int   `json:"status"`
	Data   D     `json:"data"`
	Error  error `json:"erros"`
}

func NewResponseOk[D any](data D) *Response[D] {
	return &Response[D]{
		Data:   data,
		Status: http.StatusOK,
		Error:  nil,
	}
}

func NewResponseCreated[D any](data D) *Response[D] {
	return &Response[D]{
		Status: http.StatusCreated,
		Data:   data,
		Error:  nil,
	}
}

func NewReponseNotFound(err error) *Response[any] {
	return &Response[any]{
		Status: http.StatusNotFound,
		Data:   nil,
		Error:  err,
	}
}

func NewResponseInternalError(err error) *Response[any] {
	return &Response[any]{
		Status: http.StatusInternalServerError,
		Data:   nil,
		Error:  err,
	}
}
