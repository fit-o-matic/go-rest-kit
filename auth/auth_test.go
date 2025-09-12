package auth

import (
	"net/http"
	"testing"
)

func TestNewTokenAuth(t *testing.T) {
	auth := NewTokenAuth("test-token")
	if auth.Token != "test-token" {
		t.Errorf("Expected token 'test-token', got '%s'", auth.Token)
	}
}

func TestTokenAuthSetAuth(t *testing.T) {
	auth := NewTokenAuth("test-token")
	req, _ := http.NewRequest("GET", "http://example.com", nil)

	auth.SetAuth(req)

	expected := "Bearer test-token"
	if req.Header.Get("Authorization") != expected {
		t.Errorf("Expected Authorization header '%s', got '%s'", expected, req.Header.Get("Authorization"))
	}
}

func TestNewBasicAuth(t *testing.T) {
	auth := NewBasicAuth("user", "pass")
	if auth.Username != "user" {
		t.Errorf("Expected username 'user', got '%s'", auth.Username)
	}
	if auth.Password != "pass" {
		t.Errorf("Expected password 'pass', got '%s'", auth.Password)
	}
}

func TestBasicAuthSetAuth(t *testing.T) {
	auth := NewBasicAuth("user", "pass")
	req, _ := http.NewRequest("GET", "http://example.com", nil)

	auth.SetAuth(req)

	// user:pass in base64 is dXNlcjpwYXNz
	expected := "Basic dXNlcjpwYXNz"
	if req.Header.Get("Authorization") != expected {
		t.Errorf("Expected Authorization header '%s', got '%s'", expected, req.Header.Get("Authorization"))
	}
}

func TestNewAPIKeyAuth(t *testing.T) {
	auth := NewAPIKeyAuth("api-key-123", "X-API-Key")
	if auth.Key != "api-key-123" {
		t.Errorf("Expected key 'api-key-123', got '%s'", auth.Key)
	}
	if auth.HeaderKey != "X-API-Key" {
		t.Errorf("Expected header key 'X-API-Key', got '%s'", auth.HeaderKey)
	}
}

func TestAPIKeyAuthSetAuth(t *testing.T) {
	auth := NewAPIKeyAuth("api-key-123", "X-API-Key")
	req, _ := http.NewRequest("GET", "http://example.com", nil)

	auth.SetAuth(req)

	if req.Header.Get("X-API-Key") != "api-key-123" {
		t.Errorf("Expected X-API-Key header 'api-key-123', got '%s'", req.Header.Get("X-API-Key"))
	}
}

func TestNewAuthenticatedClient(t *testing.T) {
	auth := NewTokenAuth("test-token")
	client := NewAuthenticatedClient(auth)

	if client.Client == nil {
		t.Error("Expected Client to be initialized")
	}
	if client.Authenticator != auth {
		t.Error("Expected Authenticator to be set")
	}
}

func TestBuildAuthHeader(t *testing.T) {
	result := BuildAuthHeader("Bearer", "token123")
	expected := "Bearer token123"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}