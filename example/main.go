package main

import (
	"fmt"
	"net/http"

	"github.com/fit-o-matic/go-rest-kit/restkit"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *Person) GetContentType() string {
	return "application/json"
}

func (p *Person) GetData() []byte {
	return []byte(fmt.Sprintf(`{"name": "%s", "age": %d}`, p.Name, p.Age))
}

func main() {
	resquest := restkit.NewRequestBuilder().
		WithMethod("GET").
		WithBaseURL("https://emojihub.yurace.pro/api").
		WithPath("/groups").
		Build()

	restkit.ExecuteRequestAndPrintResponse(http.DefaultClient, resquest)
}
