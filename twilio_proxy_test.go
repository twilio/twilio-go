package twilio

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twilio/twilio-go/client"
)

// TestCustomHTTPClientWithProxyURL tests the exact scenario from the issue:
// Using http.ProxyURL with a custom HTTPClient
func TestCustomHTTPClientWithProxyURL(t *testing.T) {
	// Create a mock proxy server that tracks if it received requests
	mockProxyServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			// A real proxy would forward the request, but we just track it
			writer.WriteHeader(200)
			writer.Write([]byte(`{"status":"proxied"}`))
		}))
	defer mockProxyServer.Close()

	// Parse proxy URL
	proxyURL, err := url.Parse(mockProxyServer.URL)
	assert.NoError(t, err)

	// Create http.Client with proxy (EXACT pattern from issue description)
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	// Create Twilio client with custom HTTPClient (EXACT pattern from issue description)
	accountSid := "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	authToken := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	baseClient := &client.Client{
		Credentials: client.NewCredentials(accountSid, authToken),
		HTTPClient:  httpClient,
	}
	baseClient.SetAccountSid(accountSid)

	twilioClient := NewRestClientWithParams(ClientParams{
		Client: baseClient,
	})

	// Verify client setup
	assert.NotNil(t, twilioClient)
	assert.NotNil(t, twilioClient.RequestHandler)
	assert.NotNil(t, twilioClient.RequestHandler.Client)

	// Create a mock Twilio API server (simulating api.twilio.com)
	mockTwilioServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			// This represents the actual Twilio API
			writer.WriteHeader(200)
			writer.Write([]byte(`{"account_sid":"ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","balance":"100.00","currency":"USD"}`))
		}))
	defer mockTwilioServer.Close()

	// Make a request - should go through proxy
	// Note: In a real scenario with a real proxy, this would work
	// For testing, we verify the HTTPClient with Proxy is set up correctly
	resp, err := baseClient.SendRequest("GET", mockTwilioServer.URL+"/Balance.json", nil, map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	})

	// The request should succeed (our mock server responds directly)
	// In a real proxy setup, the proxy would forward to the real server
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// The key verification: the Transport with Proxy is configured
	clientImpl, ok := twilioClient.RequestHandler.Client.(*client.Client)
	assert.True(t, ok)
	assert.NotNil(t, clientImpl.HTTPClient)
	assert.NotNil(t, clientImpl.HTTPClient.Transport)

	transportImpl, ok := clientImpl.HTTPClient.Transport.(*http.Transport)
	assert.True(t, ok, "Transport should be *http.Transport")
	assert.NotNil(t, transportImpl.Proxy, "Proxy function should be set")

	// Verify proxy function returns the correct proxy URL
	proxyURLFromTransport, err := transportImpl.Proxy(&http.Request{URL: &url.URL{Scheme: "https", Host: "api.twilio.com"}})
	assert.NoError(t, err)
	assert.Equal(t, proxyURL.String(), proxyURLFromTransport.String(), "Proxy URL should match")
}

// TestCustomHTTPClientProxyFunctionCalled tests that the proxy function is actually invoked
func TestCustomHTTPClientProxyFunctionCalled(t *testing.T) {
	proxyFunctionCalled := false
	callCount := 0

	// Create custom transport with tracking proxy function
	customTransport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			proxyFunctionCalled = true
			callCount++
			// Don't actually use a proxy, return nil
			return nil, nil
		},
	}

	httpClient := &http.Client{
		Transport: customTransport,
	}

	// Create Twilio client
	baseClient := &client.Client{
		Credentials: client.NewCredentials("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		HTTPClient:  httpClient,
	}
	baseClient.SetAccountSid("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	// Create mock server
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(200)
			writer.Write([]byte(`{}`))
		}))
	defer mockServer.Close()

	// Make request
	resp, err := baseClient.SendRequest("GET", mockServer.URL, nil, map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, proxyFunctionCalled, "Proxy function should have been called")
	assert.Greater(t, callCount, 0, "Proxy function should have been called at least once")
}
