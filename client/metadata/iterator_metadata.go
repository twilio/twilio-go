// Package metadata provides wrappers for Twilio API responses with HTTP metadata
package metadata

import (
	"fmt"
	"net/http"
	"strings"
)

// IteratorMetadata is an interface for objects that have HTTP response metadata (headers, status code)
// and implement iterator functionality
type IteratorMetadata interface {
	// GetStatusCode returns the HTTP status code of the response
	GetStatusCode() int

	// GetHeaders returns the HTTP headers of the response
	GetHeaders() http.Header

	// GetHeader returns a specific HTTP header value
	GetHeader(name string) string

	// String returns a string representation of the metadata
	String() string
}

// BaseMetadata is a struct containing common HTTP metadata fields
type BaseMetadata struct {
	StatusCode int         // HTTP status code
	Headers    http.Header // HTTP response headers
}

// NewBaseMetadata creates a new BaseMetadata instance
func NewBaseMetadata(statusCode int, headers http.Header) *BaseMetadata {
	return &BaseMetadata{
		StatusCode: statusCode,
		Headers:    headers,
	}
}

// GetStatusCode returns the HTTP status code
func (m *BaseMetadata) GetStatusCode() int {
	return m.StatusCode
}

// GetHeaders returns all HTTP headers
func (m *BaseMetadata) GetHeaders() http.Header {
	return m.Headers
}

// GetHeader returns a specific HTTP header by name (case-insensitive)
func (m *BaseMetadata) GetHeader(name string) string {
	return m.Headers.Get(name)
}

// String returns a string representation of the metadata
func (m *BaseMetadata) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Status Code: %d\n", m.StatusCode))
	builder.WriteString("Headers:\n")

	for name, values := range m.Headers {
		builder.WriteString(fmt.Sprintf("  %s: %s\n", name, strings.Join(values, ", ")))
	}

	return builder.String()
}
