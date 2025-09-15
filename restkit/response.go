package restkit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Response struct {
	Request    *Request
	StatusCode int
	Header     Header
	Body       *Body
}

func (resp *Response) IsSuccess() bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

func (resp *Response) UnmarshalJSONBody(target interface{}) error {
	if !resp.IsSuccess() {
		return fmt.Errorf("response is not successful")
	}
	return json.Unmarshal(resp.Body.Data, target)
}

func (resp *Response) PrettyString() string {
	var res strings.Builder
	res.WriteString(fmt.Sprintf("Status: %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode)))
	res.WriteString("Headers:")
	for k, v := range resp.Header {
		fmt.Printf("  %s: %s\n", k, v)
	}
	res.WriteString("Body:")
	res.WriteString(resp.Body.PrettyString())
	return res.String()
}
