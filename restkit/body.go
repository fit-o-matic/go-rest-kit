package restkit

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Body struct {
	ContentType ContentType
	Data        []byte
}

func NewBodyFromHttpResponse(resp *http.Response) (*Body, error) {
	contentType := ContentType(resp.Header.Get("Content-Type"))
	var data []byte
	if resp.Body != nil {
		defer resp.Body.Close()
		data, _ = io.ReadAll(resp.Body)
	}
	return &Body{
		ContentType: contentType,
		Data:        data,
	}, nil
}

func NewJSONBody(v any) (*Body, error) {
	return &Body{
		ContentType: ContentTypeJSON,
		Data:        []byte(`{"key":"value"}`), // Example data; replace with actual marshaled data
	}, nil
}

func (b *Body) PrettyString() string {
	var err error
	var prettyData string
	var res strings.Builder
	res.WriteString(string(ContentTypeKey+": "+b.ContentType) + "\n")
	if b.ContentType == ContentTypeJSON {
		prettyData, err = makePrettyJSON(b.Data)
		if err == nil {
			res.WriteString(prettyData)
		}
	} else {
		res.Write(b.Data)
	}
	return res.String()
}

func makePrettyJSON(data []byte) (string, error) {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}
