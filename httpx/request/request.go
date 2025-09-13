package request

import (
	"bytes"
	"io"
	"net/http"

	"github.com/fit-o-matic/go-rest-utils/httpx/body"
	"github.com/fit-o-matic/go-rest-utils/httpx/header"
)

type Request struct {
	method string
	Header header.Header
	URL    string
	body   []byte
}

func (r *Request) ToHttpRequest() (*http.Request, error) {
	httpReq, err := http.NewRequest(r.method, r.URL, nil)
	if err != nil {
		return nil, err
	}
	r.Header.CopyToHttpHeader(httpReq.Header)

	if r.body != nil {
		httpReq.Body = io.NopCloser(bytes.NewReader(r.body))
	}
	return httpReq, nil
}

func (r *Request) Do(client *http.Client) (*Response, error) {
	httpReq, err := r.ToHttpRequest()
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(httpReq)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &Response{
		Request:    r,
		StatusCode: resp.StatusCode,
		Header:     header.FromHttpHeader(resp.Header),
		Body:       body.FromHttpResponse(resp),
	}, nil
}
