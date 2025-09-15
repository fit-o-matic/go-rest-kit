


# go-rest-kit

A lightweight, chainable HTTP request builder for Go.

## Documentation

All public functions are documented with Go doc comments. You can view detailed documentation using Go tools like `go doc` or by browsing the source code.

## Features

- Chainable builder pattern for constructing HTTP requests
- Supports setting method, base URL, path, headers, query parameters, and body
- Automatic JSON serialization for struct bodies
- Simple conversion to `http.Request`

## Installation

```sh
go get github.com/fit-o-matic/go-rest-kit@0.2.0
```

## Usage

```go
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
```

## API

### Builder Methods

- `WithMethod(method string)`
- `WithBaseURL(url string)`
- `WithPath(path string)`
- `WithHeader(headers map[string]string)`
- `WithQueryParam(params map[string]string)`
- `WithBody(body interface{})`
- `Build() (*Request, error)`

### Request Methods

- `ToHttpRequest() http.Request`
- `String() string`

### How to release a new version

1. Update the version number `README.md`.
2. Commit the changes with a message like `Bump version to vX.Y.Z`.
3. Create a new Git tag: `git tag vX.Y.Z`.
4. Push the tag to the repository: `git push origin vX.Y.Z`.

## License

Apache License 2.0
