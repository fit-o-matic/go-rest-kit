// Package auth provides authentication utilities for REST clients
package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

// TokenAuth represents a bearer token authenticator
type TokenAuth struct {
	Token string
}

// NewTokenAuth creates a new token-based authenticator
func NewTokenAuth(token string) *TokenAuth {
	return &TokenAuth{Token: token}
}

// SetAuth adds the Authorization header with bearer token to the request
func (t *TokenAuth) SetAuth(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+t.Token)
}

// BasicAuth represents a basic authentication authenticator
type BasicAuth struct {
	Username string
	Password string
}

// NewBasicAuth creates a new basic authentication authenticator
func NewBasicAuth(username, password string) *BasicAuth {
	return &BasicAuth{
		Username: username,
		Password: password,
	}
}

// SetAuth adds the Authorization header with basic auth to the request
func (b *BasicAuth) SetAuth(req *http.Request) {
	auth := base64.StdEncoding.EncodeToString([]byte(b.Username + ":" + b.Password))
	req.Header.Set("Authorization", "Basic "+auth)
}

// APIKeyAuth represents an API key authenticator
type APIKeyAuth struct {
	Key       string
	HeaderKey string
}

// NewAPIKeyAuth creates a new API key authenticator
func NewAPIKeyAuth(key, headerKey string) *APIKeyAuth {
	return &APIKeyAuth{
		Key:       key,
		HeaderKey: headerKey,
	}
}

// SetAuth adds the API key header to the request
func (a *APIKeyAuth) SetAuth(req *http.Request) {
	req.Header.Set(a.HeaderKey, a.Key)
}

// Authenticator is an interface for different authentication methods
type Authenticator interface {
	SetAuth(req *http.Request)
}

// AuthenticatedClient wraps an http.Client with authentication
type AuthenticatedClient struct {
	Client        *http.Client
	Authenticator Authenticator
}

// NewAuthenticatedClient creates a new authenticated client
func NewAuthenticatedClient(auth Authenticator) *AuthenticatedClient {
	return &AuthenticatedClient{
		Client:        &http.Client{},
		Authenticator: auth,
	}
}

// Do performs an HTTP request with authentication
func (ac *AuthenticatedClient) Do(req *http.Request) (*http.Response, error) {
	if ac.Authenticator != nil {
		ac.Authenticator.SetAuth(req)
	}
	return ac.Client.Do(req)
}

// BuildAuthHeader is a utility function to build authorization headers manually
func BuildAuthHeader(authType, credentials string) string {
	return fmt.Sprintf("%s %s", authType, credentials)
}