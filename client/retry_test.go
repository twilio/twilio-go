package client_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/twilio/twilio-go/client"
)

func TestClient_AutoRetryDisabled(t *testing.T) {
	// Create a server that always returns 429
	requestCount := int32(0)
	server := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			atomic.AddInt32(&requestCount, 1)
			writer.WriteHeader(http.StatusTooManyRequests)
			_, _ = writer.Write([]byte(`{"status": 429, "code": 20429, "message": "Too Many Requests", "more_info": "https://www.twilio.com/docs/errors/20429"}`))
		}))
	defer server.Close()

	c := NewClient("user", "pass")
	c.AutoRetry = false // Explicitly disable auto-retry

	_, err := c.SendRequest("GET", server.URL, nil, nil) //nolint:bodyclose
	assert.Error(t, err)
	twilioError := err.(*client.TwilioRestError)
	assert.Equal(t, 429, twilioError.Status)
	assert.Equal(t, 20429, twilioError.Code)

	// Should only make one request
	assert.Equal(t, int32(1), atomic.LoadInt32(&requestCount))
}

func TestClient_AutoRetryEnabledEventualSuccess(t *testing.T) {
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

	c := NewClient("user", "pass")
	c.SetAutoRetry(true)
	c.SetMaxRetries(3)
	c.SetMaxRetryDelay(100) // Short delay for testing

	start := time.Now()
	resp, err := c.SendRequest("GET", server.URL, nil, nil) //nolint:bodyclose
	elapsed := time.Since(start)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Should make 3 requests total (2 failures + 1 success)
	assert.Equal(t, int32(3), atomic.LoadInt32(&requestCount))

	// Should have waited for retries (at least some time but not too long)
	assert.Greater(t, elapsed, time.Duration(0))
	assert.Less(t, elapsed, 2*time.Second) // Should be much faster with our short delay
}

func TestClient_AutoRetryMaxRetriesExceeded(t *testing.T) {
	// Create a server that always returns 429
	requestCount := int32(0)
	server := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			atomic.AddInt32(&requestCount, 1)
			writer.WriteHeader(http.StatusTooManyRequests)
			_, _ = writer.Write([]byte(`{"status": 429, "code": 20429, "message": "Too Many Requests", "more_info": "https://www.twilio.com/docs/errors/20429"}`))
		}))
	defer server.Close()

	c := NewClient("user", "pass")
	c.SetAutoRetry(true)
	c.SetMaxRetries(3)
	c.SetMaxRetryDelay(50) // Very short delay for testing

	_, err := c.SendRequest("GET", server.URL, nil, nil) //nolint:bodyclose
	assert.Error(t, err)
	twilioError := err.(*client.TwilioRestError)
	assert.Equal(t, 429, twilioError.Status)
	assert.Equal(t, 20429, twilioError.Code)

	// Should make 4 requests total (initial + 3 retries)
	assert.Equal(t, int32(4), atomic.LoadInt32(&requestCount))
}

func TestClient_AutoRetryCustomMaxRetries(t *testing.T) {
	// Create a server that always returns 429
	requestCount := int32(0)
	server := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			atomic.AddInt32(&requestCount, 1)
			writer.WriteHeader(http.StatusTooManyRequests)
			_, _ = writer.Write([]byte(`{"status": 429, "code": 20429, "message": "Too Many Requests", "more_info": "https://www.twilio.com/docs/errors/20429"}`))
		}))
	defer server.Close()

	c := NewClient("user", "pass")
	c.SetAutoRetry(true)
	c.SetMaxRetries(5) // Custom max retries
	c.SetMaxRetryDelay(50)

	_, err := c.SendRequest("GET", server.URL, nil, nil) //nolint:bodyclose
	assert.Error(t, err)
	twilioError := err.(*client.TwilioRestError)
	assert.Equal(t, 429, twilioError.Status)

	// Should make 6 requests total (initial + 5 retries)
	assert.Equal(t, int32(6), atomic.LoadInt32(&requestCount))
}

func TestClient_AutoRetryOnlyRetries429(t *testing.T) {
	// Create a server that returns a different error code
	requestCount := int32(0)
	server := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			atomic.AddInt32(&requestCount, 1)
			writer.WriteHeader(http.StatusBadRequest) // 400, not 429
			_, _ = writer.Write([]byte(`{"status": 400, "code": 20001, "message": "Bad request", "more_info": "https://www.twilio.com/docs/errors/20001"}`))
		}))
	defer server.Close()

	c := NewClient("user", "pass")
	c.SetAutoRetry(true)
	c.SetMaxRetries(3)

	_, err := c.SendRequest("GET", server.URL, nil, nil) //nolint:bodyclose
	assert.Error(t, err)
	twilioError := err.(*client.TwilioRestError)
	assert.Equal(t, 400, twilioError.Status)
	assert.Equal(t, 20001, twilioError.Code)

	// Should only make one request (no retries for non-429 errors)
	assert.Equal(t, int32(1), atomic.LoadInt32(&requestCount))
}

func TestClient_AutoRetryDefaultValues(t *testing.T) {
	// Create a server that always returns 429
	requestCount := int32(0)
	server := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			atomic.AddInt32(&requestCount, 1)
			writer.WriteHeader(http.StatusTooManyRequests)
			_, _ = writer.Write([]byte(`{"status": 429, "code": 20429, "message": "Too Many Requests", "more_info": "https://www.twilio.com/docs/errors/20429"}`))
		}))
	defer server.Close()

	c := NewClient("user", "pass")
	c.SetAutoRetry(true)
	// Don't set MaxRetries or MaxRetryDelay - should use defaults

	_, err := c.SendRequest("GET", server.URL, nil, nil) //nolint:bodyclose
	assert.Error(t, err)

	// Should use default max retries (3), so 4 requests total
	assert.Equal(t, int32(4), atomic.LoadInt32(&requestCount))
}

func TestClient_AutoRetryExponentialBackoff(t *testing.T) {
	// Create a server that always returns 429
	requestCount := int32(0)
	requestTimes := make([]time.Time, 0)
	server := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			atomic.AddInt32(&requestCount, 1)
			requestTimes = append(requestTimes, time.Now())
			writer.WriteHeader(http.StatusTooManyRequests)
			_, _ = writer.Write([]byte(`{"status": 429, "code": 20429, "message": "Too Many Requests", "more_info": "https://www.twilio.com/docs/errors/20429"}`))
		}))
	defer server.Close()

	c := NewClient("user", "pass")
	c.SetAutoRetry(true)
	c.SetMaxRetries(2)
	c.SetMaxRetryDelay(1000) // 1 second max

	_, err := c.SendRequest("GET", server.URL, nil, nil) //nolint:bodyclose
	assert.Error(t, err)

	// Should make 3 requests total (initial + 2 retries)
	assert.Equal(t, int32(3), atomic.LoadInt32(&requestCount))

	// Verify that delays increase (allowing for jitter, just check they happened)
	assert.Greater(t, len(requestTimes), 1)
	// First retry should have some delay
	firstDelay := requestTimes[1].Sub(requestTimes[0])
	assert.Greater(t, firstDelay, time.Duration(0))
	// Delays should be present (not testing exact exponential due to jitter)
	secondDelay := requestTimes[2].Sub(requestTimes[1])
	assert.Greater(t, secondDelay, time.Duration(0))
}

func TestClient_AutoRetrySuccessDoesNotRetry(t *testing.T) {
	// Create a server that always succeeds
	requestCount := int32(0)
	server := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			atomic.AddInt32(&requestCount, 1)
			writer.WriteHeader(http.StatusOK)
			d := map[string]interface{}{"response": "ok"}
			_ = json.NewEncoder(writer).Encode(&d)
		}))
	defer server.Close()

	c := NewClient("user", "pass")
	c.SetAutoRetry(true)
	c.SetMaxRetries(3)

	resp, err := c.SendRequest("GET", server.URL, nil, nil) //nolint:bodyclose
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Should only make one request (success on first try)
	assert.Equal(t, int32(1), atomic.LoadInt32(&requestCount))
}
