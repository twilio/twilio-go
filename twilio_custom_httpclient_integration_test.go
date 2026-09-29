package twilio

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/twilio/twilio-go/client"
)

// TestRestClientWithCustomHTTPClient tests that custom HTTPClient works through full RestClient flow
func TestRestClientWithCustomHTTPClient(t *testing.T) {
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

	// Create mock Twilio API server
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			// Verify Authorization header is present (basic auth)
			authHeader := request.Header.Get("Authorization")
			assert.NotEmpty(t, authHeader, "Authorization header should be present")

			writer.WriteHeader(200)
			writer.Write([]byte(`{"account_sid":"ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","balance":"100.00","currency":"USD"}`))
		}))
	defer mockServer.Close()

	// Create Twilio base client with custom HTTPClient (following documentation pattern)
	baseClient := &client.Client{
		Credentials: client.NewCredentials("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		HTTPClient:  httpClient,
	}
	baseClient.SetAccountSid("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	// Create Twilio RestClient with custom client
	twilioClient := NewRestClientWithParams(ClientParams{
		Client: baseClient,
	})

	// Verify the custom client was passed through
	assert.NotNil(t, twilioClient.RequestHandler)
	assert.NotNil(t, twilioClient.RequestHandler.Client)
	assert.Equal(t, baseClient, twilioClient.RequestHandler.Client)

	// Verify HTTPClient is accessible through the chain
	clientImpl, ok := twilioClient.RequestHandler.Client.(*client.Client)
	assert.True(t, ok, "Client should be of type *client.Client")
	assert.Equal(t, httpClient, clientImpl.HTTPClient, "Custom HTTPClient should be preserved")

	// Make a request through the client directly to verify custom HTTPClient is used
	resp, err := baseClient.SendRequest("GET", mockServer.URL+"/Balance.json", nil, map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify custom transport was used
	assert.True(t, customTransportUsed, "Custom HTTPClient transport should have been used for API requests")
}

// TestRestClientWithCustomHTTPClientAndTimeout tests that SetTimeout works with custom HTTPClient
func TestRestClientWithCustomHTTPClientAndTimeout(t *testing.T) {
	// Create a custom http.Client
	customHTTPClient := &http.Client{
		Transport: http.DefaultTransport,
	}

	// Create Twilio base client with custom HTTPClient
	baseClient := &client.Client{
		Credentials: client.NewCredentials("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		HTTPClient:  customHTTPClient,
	}
	baseClient.SetAccountSid("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	// Create Twilio RestClient
	twilioClient := NewRestClientWithParams(ClientParams{
		Client: baseClient,
	})

	// Set timeout via RestClient
	twilioClient.SetTimeout(30 * time.Second)

	// Verify the custom HTTPClient is still the same instance (not replaced)
	clientImpl, ok := twilioClient.RequestHandler.Client.(*client.Client)
	assert.True(t, ok)
	assert.Equal(t, customHTTPClient, clientImpl.HTTPClient, "SetTimeout should not replace custom HTTPClient")
	assert.Equal(t, 30*time.Second, clientImpl.HTTPClient.Timeout, "Timeout should be updated on custom HTTPClient")
}

// TestRestClientWithoutCustomClient tests default behavior when no custom client provided
func TestRestClientWithoutCustomClient(t *testing.T) {
	// Create RestClient without custom client (should create default)
	twilioClient := NewRestClientWithParams(ClientParams{
		Username: "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		Password: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	})

	// Verify a client was created
	assert.NotNil(t, twilioClient.RequestHandler)
	assert.NotNil(t, twilioClient.RequestHandler.Client)

	// Verify it's a default client
	clientImpl, ok := twilioClient.RequestHandler.Client.(*client.Client)
	assert.True(t, ok)

	// HTTPClient will be nil until first request (lazy initialization via defaultHTTPClient())
	assert.Nil(t, clientImpl.HTTPClient, "HTTPClient should be nil before first use")
}
