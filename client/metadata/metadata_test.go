package metadata

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseMetadata(t *testing.T) {
	headers := make(http.Header)
	headers.Add("Content-Type", "application/json")
	headers.Add("Twilio-Request-Id", "RQ123456789")

	metadata := NewBaseMetadata(200, headers)

	// Test GetStatusCode
	assert.Equal(t, 200, metadata.GetStatusCode())

	// Test GetHeaders
	assert.Equal(t, headers, metadata.GetHeaders())

	// Test GetHeader
	assert.Equal(t, "application/json", metadata.GetHeader("Content-Type"))
	assert.Equal(t, "RQ123456789", metadata.GetHeader("Twilio-Request-Id"))

	// Test case insensitivity
	assert.Equal(t, "application/json", metadata.GetHeader("content-type"))

	// Test String method
	str := metadata.String()
	assert.Contains(t, str, "Status Code: 200")
	assert.Contains(t, str, "Content-Type: application/json")
	assert.Contains(t, str, "Twilio-Request-Id: RQ123456789")
}

// Mock Resource for testing
type MockResource struct {
	ID   string
	Name string
}

func TestResourceMetadata(t *testing.T) {
	headers := make(http.Header)
	headers.Add("Content-Type", "application/json")

	resource := MockResource{
		ID:   "RS123",
		Name: "Test Resource",
	}

	metadata := NewResourceMetadata[MockResource](resource, 201, headers)

	// Test GetStatusCode
	assert.Equal(t, 201, metadata.GetStatusCode())

	// Test GetResource
	assert.Equal(t, resource, metadata.GetResource())

	// Test String
	str := metadata.String()
	assert.Contains(t, str, "ResourceMetadata for metadata.MockResource")
	assert.Contains(t, str, "Status Code: 201")
}
