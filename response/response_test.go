package response

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestNewResponse(t *testing.T) {
	// Test with nil response
	_, err := NewResponse(nil)
	if err == nil {
		t.Error("Expected error for nil response")
	}

	// Test with valid response
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("test body")),
	}

	r, err := NewResponse(resp)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if string(r.Body) != "test body" {
		t.Errorf("Expected body 'test body', got '%s'", string(r.Body))
	}
}

func TestResponseJSON(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"name": "test", "value": 123}`)),
	}

	r, err := NewResponse(resp)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	var data map[string]interface{}
	err = r.JSON(&data)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if data["name"] != "test" {
		t.Errorf("Expected name 'test', got '%v'", data["name"])
	}
}

func TestResponseString(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("test response")),
	}

	r, err := NewResponse(resp)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if r.String() != "test response" {
		t.Errorf("Expected 'test response', got '%s'", r.String())
	}
}

func TestIsSuccess(t *testing.T) {
	tests := []struct {
		statusCode int
		expected   bool
	}{
		{200, true},
		{201, true},
		{299, true},
		{300, false},
		{400, false},
		{500, false},
	}

	for _, test := range tests {
		resp := &http.Response{
			StatusCode: test.statusCode,
			Body:       io.NopCloser(strings.NewReader("")),
		}

		r, _ := NewResponse(resp)
		if r.IsSuccess() != test.expected {
			t.Errorf("For status %d, expected IsSuccess() to be %v, got %v",
				test.statusCode, test.expected, r.IsSuccess())
		}
	}
}

func TestIsClientError(t *testing.T) {
	tests := []struct {
		statusCode int
		expected   bool
	}{
		{400, true},
		{404, true},
		{499, true},
		{200, false},
		{500, false},
	}

	for _, test := range tests {
		resp := &http.Response{
			StatusCode: test.statusCode,
			Body:       io.NopCloser(strings.NewReader("")),
		}

		r, _ := NewResponse(resp)
		if r.IsClientError() != test.expected {
			t.Errorf("For status %d, expected IsClientError() to be %v, got %v",
				test.statusCode, test.expected, r.IsClientError())
		}
	}
}

func TestCheckError(t *testing.T) {
	// Test success response
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("success")),
	}

	r, _ := NewResponse(resp)
	if err := r.CheckError(); err != nil {
		t.Errorf("Expected no error for success response, got %v", err)
	}

	// Test error response
	resp = &http.Response{
		StatusCode: 404,
		Body:       io.NopCloser(strings.NewReader("not found")),
	}

	r, _ = NewResponse(resp)
	if err := r.CheckError(); err == nil {
		t.Error("Expected error for 404 response")
	}
}

func TestCheckStatusCode(t *testing.T) {
	resp := &http.Response{StatusCode: 200}

	// Test expected status
	if err := CheckStatusCode(resp, 200, 201); err != nil {
		t.Errorf("Expected no error for status 200, got %v", err)
	}

	// Test unexpected status
	if err := CheckStatusCode(resp, 404, 500); err == nil {
		t.Error("Expected error for unexpected status code")
	}
}