package request

import (
	"encoding/json"
	"fmt"

	"github.com/fit-o-matic/go-rest-utils/httpx/body"
	"github.com/fit-o-matic/go-rest-utils/httpx/header"
)

type Response struct {
	Request    *Request
	StatusCode int
	Header     header.Header
	Body       body.Body
}

func (resp *Response) IsSuccess() bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

func (resp *Response) UnmarshalJSONBody(target interface{}) error {
	if !resp.IsSuccess() {
		return fmt.Errorf("response is not successful")
	}
	return json.Unmarshal(resp.Body.GetData(), target)
}
