package restkit

import "net/http"

type Header map[string][]string

func (h Header) SetContentType(contentType string) {
	h[ContentTypeKey] = []string{contentType}
}

func (h Header) GetContentType() string {
	if v, ok := h[ContentTypeKey]; ok {
		if len(v) > 0 {
			return v[0]
		}
	}
	return ""
}

func (h Header) CopyToHttpHeader(header http.Header) http.Header {
	for k, v := range h {
		if v != nil && len(v) > 0 {
			header.Set(k, v[0])
		}
	}
	return header
}

func NewHeaderFromHttpHeader(h http.Header) Header {
	result := make(Header)
	for k, v := range h {
		if len(v) > 0 {
			result[k] = v
		}
	}
	return result
}
