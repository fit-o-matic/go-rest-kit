// Package main demonstrates basic usage of go-rest-utils
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/fit-o-matic/go-rest-utils/auth"
	"github.com/fit-o-matic/go-rest-utils/client"
	"github.com/fit-o-matic/go-rest-utils/response"
)

func main() {
	// Create a test server for demonstration
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for authentication
		if auth := r.Header.Get("Authorization"); auth != "" {
			fmt.Printf("Received auth header: %s\n", auth)
		}

		// Return mock data based on endpoint
		switch r.URL.Path {
		case "/users":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[{"id": 1, "name": "John Doe"}, {"id": 2, "name": "Jane Smith"}]`))
		case "/posts":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"id": 1, "title": "New Post", "status": "created"}`))
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "Not found"}`))
		}
	}))
	defer server.Close()

	fmt.Println("=== Basic Client Usage ===")
	basicClientExample(server.URL)

	fmt.Println("\n=== Authentication Example ===")
	authenticationExample(server.URL)

	fmt.Println("\n=== Response Handling Example ===")
	responseHandlingExample(server.URL)
}

func basicClientExample(baseURL string) {
	// Create a new client
	c := client.NewClient(baseURL)

	// Set default headers
	c.SetHeader("User-Agent", "go-rest-utils-example/1.0")

	// Make a GET request
	resp, err := c.Get("/users")
	if err != nil {
		log.Printf("Error making GET request: %v", err)
		return
	}

	// Parse response
	r, err := response.NewResponse(resp)
	if err != nil {
		log.Printf("Error parsing response: %v", err)
		return
	}

	// Parse JSON response
	var users []map[string]interface{}
	if err := r.JSON(&users); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return
	}

	fmt.Printf("Found %d users\n", len(users))
	for _, user := range users {
		fmt.Printf("- %s (ID: %.0f)\n", user["name"], user["id"])
	}

	// Make a POST request
	newPost := map[string]string{
		"title": "Example Post",
		"body":  "This is an example post",
	}

	postResp, err := c.Post("/posts", newPost)
	if err != nil {
		log.Printf("Error making POST request: %v", err)
		return
	}

	postR, err := response.NewResponse(postResp)
	if err != nil {
		log.Printf("Error parsing POST response: %v", err)
		return
	}

	var createdPost map[string]interface{}
	if err := postR.JSON(&createdPost); err != nil {
		log.Printf("Error parsing POST JSON: %v", err)
		return
	}

	fmt.Printf("Created post with ID: %.0f\n", createdPost["id"])
}

func authenticationExample(baseURL string) {
	// Bearer token authentication
	tokenAuth := auth.NewTokenAuth("example-bearer-token-123")
	authClient := auth.NewAuthenticatedClient(tokenAuth)

	req, _ := http.NewRequest("GET", baseURL+"/users", nil)
	resp, err := authClient.Do(req)
	if err != nil {
		log.Printf("Error with authenticated request: %v", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Authenticated request completed with status: %d\n", resp.StatusCode)

	// Basic authentication example
	basicAuth := auth.NewBasicAuth("user123", "secret456")
	req2, _ := http.NewRequest("GET", baseURL+"/users", nil)
	basicAuth.SetAuth(req2)

	fmt.Printf("Basic auth header set: %s\n", req2.Header.Get("Authorization"))

	// API key authentication example
	apiKeyAuth := auth.NewAPIKeyAuth("api-key-789", "X-API-Key")
	req3, _ := http.NewRequest("GET", baseURL+"/users", nil)
	apiKeyAuth.SetAuth(req3)

	fmt.Printf("API key header set: %s = %s\n", "X-API-Key", req3.Header.Get("X-API-Key"))
}

func responseHandlingExample(baseURL string) {
	// Test successful response
	resp, err := http.Get(baseURL + "/users")
	if err != nil {
		log.Printf("Error making request: %v", err)
		return
	}

	r, err := response.NewResponse(resp)
	if err != nil {
		log.Printf("Error creating response: %v", err)
		return
	}

	fmt.Printf("Response status: %d\n", r.StatusCode)
	fmt.Printf("Is success: %v\n", r.IsSuccess())
	fmt.Printf("Is client error: %v\n", r.IsClientError())
	fmt.Printf("Is server error: %v\n", r.IsServerError())

	// Test error response
	errorResp, err := http.Get(baseURL + "/nonexistent")
	if err != nil {
		log.Printf("Error making error request: %v", err)
		return
	}

	errorR, err := response.NewResponse(errorResp)
	if err != nil {
		log.Printf("Error creating error response: %v", err)
		return
	}

	fmt.Printf("\nError response status: %d\n", errorR.StatusCode)
	fmt.Printf("Check error result: %v\n", errorR.CheckError())
}