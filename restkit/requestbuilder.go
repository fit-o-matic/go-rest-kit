package restkit

type RequestBuilder struct {
	method      string
	header      Header
	baseURL     string
	path        string
	queryParams Query
	body        Body
}

// Request implements RequestBuilder
func (r *RequestBuilder) WithMethod(method string) *RequestBuilder {
	r.method = method
	return r
}

// WithHeader adds or overrides headers for the request. Accepts a map of header key-value pairs.
func (r *RequestBuilder) WithHeader(header Header) *RequestBuilder {
	if r.header == nil {
		r.header = make(map[string]string)
	}
	for k, v := range header {
		r.header[k] = v
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
func (r *RequestBuilder) WithBody(body Body) *RequestBuilder {
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

// Build finalizes and returns the constructed Request.
func (r *RequestBuilder) Build() *Request {
	var res *Request

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
	for k, v := range r.header {
		if v != "" {
			headers[k] = v
		}
	}

	res = &Request{
		method: r.method,
		Header: headers,
		URL:    url,
		body:   &r.body,
	}

	return res
}

// NewRequestBuilder initializes a new RequestBuilder.
func NewRequestBuilder() *RequestBuilder {
	r := &RequestBuilder{}
	return r
}
