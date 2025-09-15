package restkit

import "net/http"

type Header map[string]string

func (h Header) SetContentType(contentType string) {
	h[ContentTypeKey] = contentType
}

func (h Header) GetContentType() string {
	if v, ok := h[ContentTypeKey]; ok {
		return v
	}
	return ""
}

func (h Header) CopyToHttpHeader(header http.Header) http.Header {
	for k, v := range h {
		if v != "" {
			header.Set(k, v)
		}
	}
	return header
}

func NewHeaderFromHttpHeader(h http.Header) Header {
	result := make(Header)
	for k, v := range h {
		if len(v) > 0 {
			result[k] = v[0]
		}
	}
	return result
}
