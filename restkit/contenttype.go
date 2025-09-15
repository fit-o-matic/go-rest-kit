package restkit

type ContentType string

const (
	ContentTypeKey              = "Content-Type"
	ContentTypeJSON ContentType = "application/json"
	ContentTypeXML  ContentType = "application/xml"
	ContentTypeForm ContentType = "application/x-www-form-urlencoded"
)

func ParseContentType(s string) ContentType {
	switch s {
	case "application/json":
		return ContentTypeJSON
	case "application/xml":
		return ContentTypeXML
	case "application/x-www-form-urlencoded":
		return ContentTypeForm
	default:
		return ContentType(s)
	}
}
