package main

import (
	"fmt"
	"net/http"

	"github.com/fit-o-matic/go-rest-utils/httpx/request"
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
	resquest := request.Builder().
		WithMethod("GET").
		WithBaseURL("https://emojihub.yurace.pro/api").
		WithPath("/groups").
		Build()

	response, err := resquest.Do(&http.Client{})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var groups []string
	if err := response.UnmarshalJSONBody(&groups); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(response.Request.URL)
	fmt.Println(response.Header)

	fmt.Println("Available Groups:", groups)
}
