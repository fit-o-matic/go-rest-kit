// Package client provides utilities for REST client development
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client represents a REST client with common configuration
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Headers    map[string]string
}

// NewClient creates a new REST client with default settings
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		Headers: make(map[string]string),
	}
}

// SetHeader sets a default header for all requests
func (c *Client) SetHeader(key, value string) {
	c.Headers[key] = value
}

// Get performs a GET request
func (c *Client) Get(endpoint string) (*http.Response, error) {
	return c.makeRequest("GET", endpoint, nil)
}

// Post performs a POST request with JSON body
func (c *Client) Post(endpoint string, body interface{}) (*http.Response, error) {
	return c.makeRequest("POST", endpoint, body)
}

// Put performs a PUT request with JSON body
func (c *Client) Put(endpoint string, body interface{}) (*http.Response, error) {
	return c.makeRequest("PUT", endpoint, body)
}

// Delete performs a DELETE request
func (c *Client) Delete(endpoint string) (*http.Response, error) {
	return c.makeRequest("DELETE", endpoint, nil)
}

// makeRequest is a helper method to create and execute HTTP requests
func (c *Client) makeRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	url := c.BaseURL + endpoint

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set default headers
	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	// Set Content-Type for requests with body
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return c.HTTPClient.Do(req)
}