package client_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twilio/twilio-go/client"
)

// TestCustomHTTPClientUsed tests that a custom HTTPClient is actually used for requests
func TestCustomHTTPClientUsed(t *testing.T) {
	// Track if custom transport was used
	customTransportUsed := false

	// Create a custom transport that tracks usage
	customTransport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			customTransportUsed = true
			// Return no proxy, just track that this was called
			return nil, nil
		},
	}

	// Create http.Client with custom transport
	httpClient := &http.Client{
		Transport: customTransport,
	}

	// Create mock server
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(200)
			writer.Write([]byte(`{"status":"ok"}`))
		}))
	defer mockServer.Close()

	// Create base client with custom HTTPClient
	baseClient := &client.Client{
		Credentials: client.NewCredentials("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		HTTPClient:  httpClient,
	}
	baseClient.SetAccountSid("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	// Test sending a request through the client
	resp, err := baseClient.SendRequest("GET", mockServer.URL+"/test", nil, map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	
	// Verify custom transport was used
	assert.True(t, customTransportUsed, "Custom HTTPClient transport should have been used")
}

// TestCustomHTTPClientViaRequestHandler tests that custom HTTPClient works through RequestHandler
func TestCustomHTTPClientViaRequestHandler(t *testing.T) {
	// Track if custom transport was used
	customTransportUsed := false

	// Create a custom transport that tracks usage
	customTransport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			customTransportUsed = true
			// Return no proxy, just track that this was called
			return nil, nil
		},
	}

	// Create http.Client with custom transport
	httpClient := &http.Client{
		Transport: customTransport,
	}

	// Create mock server
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(200)
			writer.Write([]byte(`{"status":"ok"}`))
		}))
	defer mockServer.Close()

	// Create base client with custom HTTPClient
	baseClient := &client.Client{
		Credentials: client.NewCredentials("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		HTTPClient:  httpClient,
	}
	baseClient.SetAccountSid("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	// Create RequestHandler with the base client
	requestHandler := client.NewRequestHandler(baseClient)

	// Test sending a request through the RequestHandler
	resp, err := requestHandler.Get(mockServer.URL+"/test", nil, map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	
	// Verify custom transport was used
	assert.True(t, customTransportUsed, "Custom HTTPClient transport should have been used through RequestHandler")
}

// TestDefaultHTTPClientCreatedWhenNil tests that a default HTTPClient is created when nil
func TestDefaultHTTPClientCreatedWhenNil(t *testing.T) {
	// Create mock server
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(200)
			writer.Write([]byte(`{"status":"ok"}`))
		}))
	defer mockServer.Close()

	// Create base client WITHOUT custom HTTPClient (nil)
	baseClient := &client.Client{
		Credentials: client.NewCredentials("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		HTTPClient:  nil, // Explicitly set to nil
	}
	baseClient.SetAccountSid("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	// Test sending a request - should create default HTTP client
	resp, err := baseClient.SendRequest("GET", mockServer.URL+"/test", nil, map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}
