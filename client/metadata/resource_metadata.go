package metadata

import (
	"fmt"
	"net/http"
	"strings"
)

// ResourceMetadata is a wrapper containing a resource along with HTTP response metadata
type ResourceMetadata[T any] struct {
	*BaseMetadata
	Resource T // The resource being wrapped
}

// NewResourceMetadata creates a new ResourceMetadata instance
func NewResourceMetadata[T any](resource T, statusCode int, headers http.Header) *ResourceMetadata[T] {
	return &ResourceMetadata[T]{
		BaseMetadata: NewBaseMetadata(statusCode, headers),
		Resource:     resource,
	}
}

// GetResource returns the underlying resource
func (m *ResourceMetadata[T]) GetResource() T {
	return m.Resource
}

// String returns a string representation of the resource metadata
func (m *ResourceMetadata[T]) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("ResourceMetadata for %T\n", m.Resource))
	builder.WriteString(m.BaseMetadata.String())
	return builder.String()
}
