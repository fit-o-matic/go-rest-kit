package restkit

import (
	"bytes"
	"io"
	"net/http"
)

type Request struct {
	method string
	Header Header
	URL    string
	body   *Body
}

func (r *Request) ToHttpRequest() (*http.Request, error) {
	httpReq, err := http.NewRequest(r.method, r.URL, nil)
	if err != nil {
		return nil, err
	}
	r.Header.CopyToHttpHeader(httpReq.Header)

	if r.body != nil {
		httpReq.Body = io.NopCloser(bytes.NewReader(r.body.Data))
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

	var body *Body

	body, err = NewBodyFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &Response{
		Request:    r,
		StatusCode: resp.StatusCode,
		Header:     NewHeaderFromHttpHeader(resp.Header),
		Body:       body,
	}, nil
}
