# go-rest-utils

Utilities to speed up REST client development written in Go.

## Installation

```bash
go get github.com/fit-o-matic/go-rest-utils
```

## Features

- **Client**: Simple REST client with fluent API for GET, POST, PUT, DELETE operations
- **Auth**: Multiple authentication methods (Bearer token, Basic auth, API key)
- **Response**: Enhanced response handling with JSON parsing and error checking

## Usage

### Basic REST Client

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/fit-o-matic/go-rest-utils/client"
    "github.com/fit-o-matic/go-rest-utils/response"
)

func main() {
    // Create a new client
    c := client.NewClient("https://api.example.com")
    
    // Set default headers
    c.SetHeader("User-Agent", "MyApp/1.0")
    
    // Make a GET request
    resp, err := c.Get("/users")
    if err != nil {
        log.Fatal(err)
    }
    
    // Parse response
    r, err := response.NewResponse(resp)
    if err != nil {
        log.Fatal(err)
    }
    
    // Check for errors
    if err := r.CheckError(); err != nil {
        log.Fatal(err)
    }
    
    // Parse JSON response
    var users []map[string]interface{}
    if err := r.JSON(&users); err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Found %d users\n", len(users))
}
```

### Authentication

```go
package main

import (
    "log"
    "net/http"
    
    "github.com/fit-o-matic/go-rest-utils/auth"
)

func main() {
    // Bearer token authentication
    tokenAuth := auth.NewTokenAuth("your-bearer-token")
    
    // Basic authentication
    basicAuth := auth.NewBasicAuth("username", "password")
    
    // API key authentication
    apiKeyAuth := auth.NewAPIKeyAuth("your-api-key", "X-API-Key")
    
    // Use with authenticated client
    client := auth.NewAuthenticatedClient(tokenAuth)
    
    req, _ := http.NewRequest("GET", "https://api.example.com/protected", nil)
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
}
```

### Advanced Response Handling

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    
    "github.com/fit-o-matic/go-rest-utils/response"
)

func main() {
    // Make an HTTP request (example)
    resp, err := http.Get("https://api.example.com/data")
    if err != nil {
        log.Fatal(err)
    }
    
    // Wrap with enhanced response
    r, err := response.NewResponse(resp)
    if err != nil {
        log.Fatal(err)
    }
    
    // Check response status
    if r.IsSuccess() {
        fmt.Println("Request successful!")
    } else if r.IsClientError() {
        fmt.Println("Client error (4xx)")
    } else if r.IsServerError() {
        fmt.Println("Server error (5xx)")
    }
    
    // Get response as string
    fmt.Println("Response body:", r.String())
}
```

## API Reference

### Client Package

- `NewClient(baseURL string) *Client`: Create a new REST client
- `SetHeader(key, value string)`: Set default headers
- `Get(endpoint string) (*http.Response, error)`: Perform GET request
- `Post(endpoint string, body interface{}) (*http.Response, error)`: Perform POST request
- `Put(endpoint string, body interface{}) (*http.Response, error)`: Perform PUT request
- `Delete(endpoint string) (*http.Response, error)`: Perform DELETE request

### Auth Package

- `NewTokenAuth(token string) *TokenAuth`: Bearer token authentication
- `NewBasicAuth(username, password string) *BasicAuth`: Basic authentication
- `NewAPIKeyAuth(key, headerKey string) *APIKeyAuth`: API key authentication
- `NewAuthenticatedClient(auth Authenticator) *AuthenticatedClient`: Create authenticated client

### Response Package

- `NewResponse(resp *http.Response) (*Response, error)`: Wrap HTTP response
- `JSON(v interface{}) error`: Parse JSON response
- `String() string`: Get response as string
- `IsSuccess() bool`: Check if status is 2xx
- `IsClientError() bool`: Check if status is 4xx
- `IsServerError() bool`: Check if status is 5xx
- `CheckError() error`: Return error if response indicates failure

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the terms of the LICENSE file.
