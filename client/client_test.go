package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient("https://api.example.com")

	if client.BaseURL != "https://api.example.com" {
		t.Errorf("Expected BaseURL to be 'https://api.example.com', got '%s'", client.BaseURL)
	}

	if client.HTTPClient.Timeout != 30*time.Second {
		t.Errorf("Expected timeout to be 30s, got %v", client.HTTPClient.Timeout)
	}

	if client.Headers == nil {
		t.Error("Expected Headers to be initialized")
	}
}

func TestSetHeader(t *testing.T) {
	client := NewClient("https://api.example.com")
	client.SetHeader("Authorization", "Bearer token123")

	if client.Headers["Authorization"] != "Bearer token123" {
		t.Errorf("Expected Authorization header to be 'Bearer token123', got '%s'", client.Headers["Authorization"])
	}
}

func TestGet(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/test" {
			t.Errorf("Expected path '/test', got '%s'", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	}))
	defer server.Close()

	client := NewClient(server.URL)
	resp, err := client.Get("/test")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestPost(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type 'application/json', got '%s'", r.Header.Get("Content-Type"))
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	testData := map[string]string{"key": "value"}
	resp, err := client.Post("/test", testData)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}
}