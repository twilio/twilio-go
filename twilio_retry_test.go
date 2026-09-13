package twilio_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/client"
)

func TestRestClient_SetAutoRetryConfiguration(t *testing.T) {
	// Create a server that returns 429 twice, then succeeds
	requestCount := int32(0)
	server := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			count := atomic.AddInt32(&requestCount, 1)
			if count <= 2 {
				writer.WriteHeader(http.StatusTooManyRequests)
				_, _ = writer.Write([]byte(`{"status": 429, "code": 20429, "message": "Too Many Requests", "more_info": "https://www.twilio.com/docs/errors/20429"}`))
			} else {
				writer.WriteHeader(http.StatusOK)
				d := map[string]interface{}{"response": "ok"}
				_ = json.NewEncoder(writer).Encode(&d)
			}
		}))
	defer server.Close()

	// Create a custom client with retry configuration
	baseClient := &client.Client{
		Credentials: client.NewCredentials("test", "test"),
	}
	baseClient.SetAccountSid("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

	restClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Client: baseClient,
	})

	// Configure retry behavior
	restClient.SetAutoRetry(true)
	restClient.SetMaxRetries(3)
	restClient.SetMaxRetryDelay(100) // Short delay for testing

	// Make a request
	resp, err := restClient.RequestHandler.Get(server.URL, nil, nil, "")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Should have made 3 requests (2 failures + 1 success)
	assert.Equal(t, int32(3), atomic.LoadInt32(&requestCount))
}

func TestRestClient_AutoRetryDisabledByDefault(t *testing.T) {
	// Create a server that always returns 429
	requestCount := int32(0)
	server := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			atomic.AddInt32(&requestCount, 1)
			writer.WriteHeader(http.StatusTooManyRequests)
			_, _ = writer.Write([]byte(`{"status": 429, "code": 20429, "message": "Too Many Requests", "more_info": "https://www.twilio.com/docs/errors/20429"}`))
		}))
	defer server.Close()

	restClient := twilio.NewRestClient()
	// Don't enable auto-retry

	_, err := restClient.RequestHandler.Get(server.URL, nil, nil, "")
	assert.Error(t, err)

	// Should only make one request
	assert.Equal(t, int32(1), atomic.LoadInt32(&requestCount))
}

func TestRestClient_SetAutoRetryCustomRetries(t *testing.T) {
	// Create a server that always returns 429
	requestCount := int32(0)
	server := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			atomic.AddInt32(&requestCount, 1)
			writer.WriteHeader(http.StatusTooManyRequests)
			_, _ = writer.Write([]byte(`{"status": 429, "code": 20429, "message": "Too Many Requests", "more_info": "https://www.twilio.com/docs/errors/20429"}`))
		}))
	defer server.Close()

	restClient := twilio.NewRestClient()
	restClient.SetAutoRetry(true)
	restClient.SetMaxRetries(5) // Custom retry count
	restClient.SetMaxRetryDelay(50)

	_, err := restClient.RequestHandler.Get(server.URL, nil, nil, "")
	assert.Error(t, err)

	// Should make 6 requests total (initial + 5 retries)
	assert.Equal(t, int32(6), atomic.LoadInt32(&requestCount))
}
