// Package response provides utilities for handling REST API responses
package response

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Response wraps an HTTP response with additional utilities
type Response struct {
	*http.Response
	Body []byte
}

// NewResponse creates a new Response from an http.Response
func NewResponse(resp *http.Response) (*Response, error) {
	if resp == nil {
		return nil, fmt.Errorf("response cannot be nil")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	resp.Body.Close()

	return &Response{
		Response: resp,
		Body:     body,
	}, nil
}

// JSON unmarshals the response body into the provided interface
func (r *Response) JSON(v interface{}) error {
	if len(r.Body) == 0 {
		return fmt.Errorf("response body is empty")
	}

	if err := json.Unmarshal(r.Body, v); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}

// String returns the response body as a string
func (r *Response) String() string {
	return string(r.Body)
}

// IsSuccess returns true if the status code indicates success (2xx)
func (r *Response) IsSuccess() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

// IsClientError returns true if the status code indicates client error (4xx)
func (r *Response) IsClientError() bool {
	return r.StatusCode >= 400 && r.StatusCode < 500
}

// IsServerError returns true if the status code indicates server error (5xx)
func (r *Response) IsServerError() bool {
	return r.StatusCode >= 500 && r.StatusCode < 600
}

// ErrorInfo represents error information from API responses
type ErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Error returns the error message
func (e *ErrorInfo) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("API error %d: %s (%s)", e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("API error %d: %s", e.Code, e.Message)
}

// CheckError examines the response and returns an error if the status indicates failure
func (r *Response) CheckError() error {
	if r.IsSuccess() {
		return nil
	}

	// Try to parse error response as JSON
	var errInfo ErrorInfo
	if err := r.JSON(&errInfo); err == nil {
		errInfo.Code = r.StatusCode
		return &errInfo
	}

	// Fallback to generic error with status and body
	return fmt.Errorf("HTTP %d: %s", r.StatusCode, r.String())
}

// ParseJSON is a utility function to parse JSON from http.Response directly
func ParseJSON(resp *http.Response, v interface{}) error {
	r, err := NewResponse(resp)
	if err != nil {
		return err
	}
	return r.JSON(v)
}

// CheckStatusCode is a utility function to check if response status is in expected range
func CheckStatusCode(resp *http.Response, expectedCodes ...int) error {
	for _, code := range expectedCodes {
		if resp.StatusCode == code {
			return nil
		}
	}
	return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
}