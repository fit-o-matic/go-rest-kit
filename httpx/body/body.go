package body

import (
	"io"
	"net/http"
)

type Body interface {
	GetContentType() string
	GetData() []byte
}

type SimpleBody struct {
	contentType string
	data        []byte
}

func (b *SimpleBody) GetContentType() string {
	return b.contentType
}

func (b *SimpleBody) GetData() []byte {
	return b.data
}

func FromHttpResponse(response *http.Response) *SimpleBody {
	contentType := response.Header.Get("Content-Type")
	data, _ := io.ReadAll(response.Body)
	return &SimpleBody{
		contentType: contentType,
		data:        data,
	}
}
