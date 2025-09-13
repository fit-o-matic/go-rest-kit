package header

import "net/http"

type Header map[string]string

const (
	ContentType = "Content-Type"
)

func (h Header) SetContentType(contentType string) {
	h[ContentType] = contentType
}

func (h Header) GetContentType() string {
	if v, ok := h[ContentType]; ok {
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

func FromHttpHeader(h http.Header) Header {
	result := make(Header)
	for k, v := range h {
		if len(v) > 0 {
			result[k] = v[0]
		}
	}
	return result
}
