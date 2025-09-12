package main

import (
	"fmt"

	"github.com/fit-o-matic/go-rest-utils/restutil"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	req, err := restutil.NewRequest().
		WithMethod("POST").
		WithBaseURL("https://api.example.com").
		WithPath("/v1/resource").
		WithHeader(map[string]string{
			"Content-Type":  "should be overridden by body",
			"Authorization": "Bearer token",
		}).
		WithQueryParam(map[string]string{
			"param1": "value1",
			"param2": "",
		}).
		WithBody(Person{Name: "John", Age: 30}).
		Build()
	if err != nil {
		panic(err)
	}
	fmt.Println(req)
}
