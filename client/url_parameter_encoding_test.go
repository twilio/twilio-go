package client_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	twilio "github.com/twilio/twilio-go/client"
)

// createTestClient creates a new Client instance for URL encoding tests
func createTestClient(accountSid string, authToken string) *twilio.Client {
	c := &twilio.Client{
		Credentials: twilio.NewCredentials(accountSid, authToken),
		HTTPClient:  http.DefaultClient,
	}

	return c
}

// TestURLParameterEncoding tests the proper encoding of URL parameters according to RFC 3986
// URI Parameter Safety Guidelines.
func TestURLParameterEncoding(t *testing.T) {
	testClient := createTestClient("user", "pass")

	t.Run("URL-safe characters should not be percent-encoded", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, req *http.Request) {
				// Parse the URL to inspect query parameters
				assert.Equal(t, "AlphaNumeric123", req.URL.Query().Get("alphanumeric"))
				assert.Equal(t, "dash-underscore_dot.tilde~", req.URL.Query().Get("unreserved"))
				w.WriteHeader(http.StatusOK)
			}))
		defer server.Close()

		data := url.Values{}
		data.Set("alphanumeric", "AlphaNumeric123")
		data.Set("unreserved", "dash-underscore_dot.tilde~")

		resp, err := testClient.SendRequest("GET", server.URL, data, nil)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("Reserved characters should be percent-encoded", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, req *http.Request) {
				// The raw query string should contain encoded values
				rawQuery := req.URL.RawQuery

				// Check that specific characters were encoded
				assert.Contains(t, rawQuery, "%21") // ! encoded
				assert.Contains(t, rawQuery, "%2A") // * encoded
				assert.Contains(t, rawQuery, "%27") // ' encoded
				assert.Contains(t, rawQuery, "%28") // ( encoded
				assert.Contains(t, rawQuery, "%29") // ) encoded
				assert.Contains(t, rawQuery, "%3B") // ; encoded
				assert.Contains(t, rawQuery, "%3A") // : encoded
				assert.Contains(t, rawQuery, "%40") // @ encoded
				assert.Contains(t, rawQuery, "%26") // & encoded
				assert.Contains(t, rawQuery, "%3D") // = encoded
				assert.Contains(t, rawQuery, "%2B") // + encoded
				assert.Contains(t, rawQuery, "%24") // $ encoded
				assert.Contains(t, rawQuery, "%2C") // , encoded
				assert.Contains(t, rawQuery, "%2F") // / encoded
				assert.Contains(t, rawQuery, "%3F") // ? encoded
				assert.Contains(t, rawQuery, "%25") // % encoded
				assert.Contains(t, rawQuery, "%23") // # encoded
				assert.Contains(t, rawQuery, "%5B") // [ encoded
				assert.Contains(t, rawQuery, "%5D") // ] encoded

				// When decoded, the value should match the original
				assert.Equal(t, "!*'();:@&=+$,/?%#[]", req.URL.Query().Get("reserved"))
				w.WriteHeader(http.StatusOK)
			}))
		defer server.Close()

		data := url.Values{}
		data.Set("reserved", "!*'();:@&=+$,/?%#[]")

		resp, err := testClient.SendRequest("GET", server.URL, data, nil)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("UTF-8 characters should be properly percent-encoded", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, req *http.Request) {
				// When decoded, the value should match the original
				assert.Equal(t, "résumé", req.URL.Query().Get("utf8"))
				assert.Equal(t, "こんにちは", req.URL.Query().Get("japanese"))
				assert.Equal(t, "👍", req.URL.Query().Get("emoji"))
				w.WriteHeader(http.StatusOK)
			}))
		defer server.Close()

		data := url.Values{}
		data.Set("utf8", "résumé")
		data.Set("japanese", "こんにちは")
		data.Set("emoji", "👍")

		resp, err := testClient.SendRequest("GET", server.URL, data, nil)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("Parameter names should follow RFC 3986 guidelines", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, req *http.Request) {
				// Get raw query string and verify our parameter names are present
				rawQuery := req.URL.RawQuery
				assert.Contains(t, rawQuery, "valid-param_name")
				assert.Contains(t, rawQuery, "valid_param_name")

				// Parameter names with alphanumeric and unreserved characters should be present in the query
				q := req.URL.Query()
				assert.Equal(t, "value1", q.Get("valid-param_name"))
				assert.Equal(t, "value2", q.Get("valid_param_name"))
				w.WriteHeader(http.StatusOK)
			}))
		defer server.Close()

		data := url.Values{}
		data.Set("valid-param_name", "value1")
		data.Set("valid_param_name", "value2")

		resp, err := testClient.SendRequest("GET", server.URL, data, nil)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("Parameter value with spaces should be properly encoded", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, req *http.Request) {
				// The raw query string should contain encoded spaces (+ for space in query params is standard)
				assert.Contains(t, req.URL.RawQuery, "John+Doe")

				// When decoded, the value should match the original with spaces
				assert.Equal(t, "John Doe", req.URL.Query().Get("name"))
				w.WriteHeader(http.StatusOK)
			}))
		defer server.Close()

		data := url.Values{}
		data.Set("name", "John Doe")

		resp, err := testClient.SendRequest("GET", server.URL, data, nil)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("Special security-sensitive characters should be encoded", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, req *http.Request) {
				rawQuery := req.URL.RawQuery

				// Test that potentially dangerous script injection is properly encoded
				scriptValue := req.URL.Query().Get("script")
				assert.Equal(t, "<script>alert('XSS')</script>", scriptValue)
				// Verify script tags are encoded
				assert.Contains(t, rawQuery, "%3Cscript%3E")

				// Test that SQL injection attempt is properly encoded
				sqlValue := req.URL.Query().Get("sql")
				assert.Equal(t, "1' OR '1'='1", sqlValue)
				// Verify single quotes are encoded
				assert.Contains(t, rawQuery, "%27")

				w.WriteHeader(http.StatusOK)
			}))
		defer server.Close()

		data := url.Values{}
		data.Set("script", "<script>alert('XSS')</script>")
		data.Set("sql", "1' OR '1'='1")

		resp, err := testClient.SendRequest("GET", server.URL, data, nil)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
}

// TestURLParameterEncodingWithJsonContent tests URL parameter encoding when Content-Type is application/json
func TestURLParameterEncodingWithJsonContent(t *testing.T) {
	testClient := createTestClient("user", "pass")

	t.Run("URL parameters should be properly encoded with JSON content type", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, req *http.Request) {
				// Parse the URL to inspect query parameters
				assert.Equal(t, "value with spaces", req.URL.Query().Get("param"))
				assert.Equal(t, "!*'();:@&=+$,/?%#[]", req.URL.Query().Get("reserved"))

				// Verify Content-Type header is set to application/json
				assert.Equal(t, "application/json", req.Header.Get("Content-Type"))

				w.WriteHeader(http.StatusOK)
			}))
		defer server.Close()

		data := url.Values{}
		data.Set("param", "value with spaces")
		data.Set("reserved", "!*'();:@&=+$,/?%#[]")

		headers := map[string]interface{}{
			"Content-Type": "application/json",
		}

		// Send an empty JSON body with the request
		jsonBody := []byte("{}")
		resp, err := testClient.SendRequest("POST", server.URL, data, headers, jsonBody...)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
}

// TestURLParameterEncodingWithFormUrlencoded tests URL parameter encoding when Content-Type is application/x-www-form-urlencoded
func TestURLParameterEncodingWithFormUrlencoded(t *testing.T) {
	testClient := createTestClient("user", "pass")

	t.Run("Parameters should be properly encoded in form body for POST requests", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, req *http.Request) {
				// Parse the form data
				err := req.ParseForm()
				assert.NoError(t, err)

				// Form values should be properly decoded
				assert.Equal(t, "value with spaces", req.Form.Get("param"))
				assert.Equal(t, "!*'();:@&=+$,/?%#[]", req.Form.Get("reserved"))

				// Verify Content-Type header
				assert.Equal(t, "application/x-www-form-urlencoded", req.Header.Get("Content-Type"))

				w.WriteHeader(http.StatusOK)
			}))
		defer server.Close()

		data := url.Values{}
		data.Set("param", "value with spaces")
		data.Set("reserved", "!*'();:@&=+$,/?%#[]")

		headers := map[string]interface{}{
			"Content-Type": "application/x-www-form-urlencoded",
		}

		resp, err := testClient.SendRequest("POST", server.URL, data, headers)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
}

// TestURLParameterEncodingEdgeCases tests edge cases for URL parameter encoding
func TestURLParameterEncodingEdgeCases(t *testing.T) {
	testClient := createTestClient("user", "pass")

	t.Run("Very long parameter values should be handled properly", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, req *http.Request) {
				// Create a very long string (1000 characters)
				expectedValue := ""
				for i := 0; i < 1000; i++ {
					expectedValue += "a"
				}

				// The value should be preserved
				assert.Equal(t, expectedValue, req.URL.Query().Get("longparam"))
				w.WriteHeader(http.StatusOK)
			}))
		defer server.Close()

		// Create a very long parameter value
		longValue := ""
		for i := 0; i < 1000; i++ {
			longValue += "a"
		}

		data := url.Values{}
		data.Set("longparam", longValue)

		resp, err := testClient.SendRequest("GET", server.URL, data, nil)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("Empty parameter values should be preserved", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, req *http.Request) {
				// Empty parameter should be present but with empty value
				assert.True(t, req.URL.Query().Has("emptyparam"))
				assert.Equal(t, "", req.URL.Query().Get("emptyparam"))
				w.WriteHeader(http.StatusOK)
			}))
		defer server.Close()

		data := url.Values{}
		data.Set("emptyparam", "")

		resp, err := testClient.SendRequest("GET", server.URL, data, nil)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
}
