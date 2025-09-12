package restutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Request struct {
	method  string
	headers map[string]string
	URL     string
	body    []byte
}

func (r *Request) ToHttpRequest() http.Request {
	req, _ := http.NewRequest(r.method, r.URL, nil)
	for k, v := range r.headers {
		req.Header.Set(k, v)
	}
	return *req
}

type Body struct {
	ContentType string
	Data        []byte
}

// CreateBody serializes the provided body into a Body struct with content type and data.
func CreateBody(body interface{}) (*Body, error) {
	switch v := body.(type) {
	case string:
		return &Body{
			ContentType: "text/plain",
			Data:        []byte(v),
		}, nil
	case []byte:
		return &Body{
			ContentType: "application/octet-stream",
			Data:        v,
		}, nil
	default:
		// serialize to json
		data, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		return &Body{
			ContentType: "application/json",
			Data:        data,
		}, nil
	}
}

func (r *Request) String() string {
	// pretty print in multi line and tabbed format
	return fmt.Sprintf("Request{\n\tmethod: %s,\n\theaders: %v,\n\tURL: %s,\n\tbody: %v\n}", r.method, r.headers, r.URL, r.body)
}

type RequestBuilder struct {
	method      string
	headers     map[string]string
	baseURL     string
	path        string
	queryParams map[string]string
	body        interface{}
}

// Request implements RequestBuilder
func (r *RequestBuilder) WithMethod(method string) *RequestBuilder {
	r.method = method
	return r
}

// WithHeader adds or overrides headers for the request. Accepts a map of header key-value pairs.
func (r *RequestBuilder) WithHeader(values map[string]string) *RequestBuilder {
	if r.headers == nil {
		r.headers = make(map[string]string)
	}
	for k, v := range values {
		r.headers[k] = v
	}
	return r
}

// WithQueryParam adds or overrides query parameters for the request. Accepts a map of query key-value pairs.
func (r *RequestBuilder) WithQueryParam(values map[string]string) *RequestBuilder {
	if r.queryParams == nil {
		r.queryParams = make(map[string]string)
	}
	for k, v := range values {
		r.queryParams[k] = v
	}
	return r
}

// WithBody sets the body of the request.
func (r *RequestBuilder) WithBody(body interface{}) *RequestBuilder {
	r.body = body
	return r
}

// WithBaseURL sets the base URL (e.g., "https://api.example.com") for the request.
func (r *RequestBuilder) WithBaseURL(baseURL string) *RequestBuilder {
	r.baseURL = baseURL
	return r
}

// WithPath sets the path (e.g., "/v1/resource") to be appended to the base URL.
func (r *RequestBuilder) WithPath(path string) *RequestBuilder {
	r.path = path
	return r
}

func (r *RequestBuilder) check() error {
	if r.method == "" {
		return errors.New("method is required")
	}
	if r.method != "GET" && r.method != "POST" && r.method != "PUT" && r.method != "DELETE" && r.method != "PATCH" && r.method != "HEAD" && r.method != "OPTIONS" {
		return errors.New("invalid method")
	}
	if r.baseURL == "" {
		return errors.New("baseURL is required")
	}
	if (r.method == "HEAD" || r.method == "GET") && r.body != nil {
		return errors.New("body is not allowed for HEAD or GET requests")
	}
	return nil
}

// Build finalizes and returns the constructed Request.
func (r *RequestBuilder) Build() (*Request, error) {
	var err error
	var res *Request

	if err = r.check(); err != nil {
		return nil, err
	}

	// Construct the full URL
	url := r.baseURL
	if r.path != "" {
		url += r.path
	}
	if len(r.queryParams) > 0 {
		first := true
		for k, v := range r.queryParams {
			if v != "" {
				if first {
					url += "?"
					first = false
				} else {
					url += "&"
				}
				url += k + "=" + v
			} else {
				continue
			}
		}
	}

	headers := make(map[string]string)
	for k, v := range r.headers {
		if v != "" {
			headers[k] = v
		}
	}

	body, err := CreateBody(r.body)
	if err != nil {
		return nil, err
	}
	if body != nil {
		headers["Content-Type"] = body.ContentType
	}

	res = &Request{
		method:  r.method,
		headers: headers,
		URL:     url,
		body:    body.Data,
	}

	return res, nil
}

// NewRequest initializes a new RequestBuilder.
func NewRequest() *RequestBuilder {
	r := &RequestBuilder{}
	return r
}
