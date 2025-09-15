package restkit

import "net/http"

func ExecuteRequest(client *http.Client, req *Request) (*Response, error) {
	return req.Do(client)
}

func ExecuteRequestAndUnmarshalJSON(client *http.Client, req *Request, target any) error {
	resp, err := req.Do(client)
	if err != nil {
		return err
	}

	return resp.UnmarshalJSONBody(target)
}

func ExecuteRequestAndPrintResponse(client *http.Client, req *Request) error {
	resp, err := req.Do(client)
	if err != nil {
		return err
	}

	println(resp.PrettyString())
	return nil
}
