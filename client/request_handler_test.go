package client_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twilio/twilio-go/client"
)

func NewRequestHandler(accountSid string, authToken string) *client.RequestHandler {
	c := NewClient(accountSid, authToken)
	return client.NewRequestHandler(c)
}

func TestRequestHandler_BuildUrlSetRegion(t *testing.T) {
	// Region set via url
	requestHandler := NewRequestHandler("user", "pass")
	assert.Equal(t, "https://api.region.twilio.com", requestHandler.BuildUrl("https://api.region.twilio.com"))

	// Region set via requestHandler
	requestHandler.Region = "region"
	assert.Equal(t, "https://api.region.twilio.com", requestHandler.BuildUrl("https://api.twilio.com"))
	assert.Equal(t, "https://api.region.twilio.com", requestHandler.BuildUrl("https://api.urlRegion.twilio.com"))
}

func TestRequestHandler_BuildUrlSetEdgeDefaultRegion(t *testing.T) {
	// Edge set via client
	requestHandler := NewRequestHandler("user", "pass")
	requestHandler.Edge = "edge"
	assert.Equal(t, "https://api.edge.us1.twilio.com", requestHandler.BuildUrl("https://api.twilio.com"))
}

func TestRequestHandler_BuildUrlSetEdgeRegion(t *testing.T) {
	//Edge and Region set via url
	requestHandler := NewRequestHandler("user", "pass")
	assert.Equal(t, "https://api.edge.region.twilio.com", requestHandler.BuildUrl("https://api.edge.region.twilio.com"))

	// Edge and Region set via client
	requestHandler.Edge = "edge"
	assert.Equal(t, "https://api.edge.region.twilio.com", requestHandler.BuildUrl("https://api.region.twilio.com"))
	requestHandler.Region = "region"
	assert.Equal(t, "https://api.edge.region.twilio.com", requestHandler.BuildUrl("https://api.twilio.com"))
	assert.Equal(t, "https://api.edge.region.twilio.com", requestHandler.BuildUrl("https://api.urlEdge.urlRegion.twilio.com"))
}

func TestRequestHandler_BuildHostRawHostWithoutPeriods(t *testing.T) {
	requestHandler := NewRequestHandler("user", "pass")
	assert.Equal(t, "https://prism_twilio:4010", requestHandler.BuildUrl("https://prism_twilio:4010"))
}

func TestRequestHandler_Post(t *testing.T) {
	mockServer := newMockServer(t)
	defer mockServer.Close()

	body := url.Values{
		"priority": []string{"1"},
	}
	headers := map[string]interface{}{}

	requestHandler := NewRequestHandler("user", "pass")
	resp, err := requestHandler.Post(mockServer.URL, body, headers)

	assert.Nil(t, err)
	assert.Equal(t, "POST", resp.Request.Method)
	assert.Equal(t, "application/x-www-form-urlencoded", resp.Request.Header.Get("Content-Type"))
}

func TestRequestHandler_PostJson(t *testing.T) {
	mockServer := newMockServer(t)
	defer mockServer.Close()

	body := map[string]interface{}{
		"priority": 1,
	}
	headers := map[string]interface{}{}

	requestHandler := NewRequestHandler("user", "pass")
	resp, err := requestHandler.PostJson(mockServer.URL, body, headers)

	assert.NoError(t, err)
	assert.Equal(t, "POST", resp.Request.Method)
	assert.Equal(t, "application/json", resp.Request.Header.Get("Content-Type"))
}

func TestRequestHandler_Get(t *testing.T) {
	mockServer := newMockServer(t)
	defer mockServer.Close()

	headers := map[string]interface{}{}

	requestHandler := NewRequestHandler("user", "pass")
	resp, err := requestHandler.Get(mockServer.URL, nil, headers)

	assert.NoError(t, err)
	assert.Equal(t, "GET", resp.Request.Method)
	assert.Equal(t, "", resp.Request.URL.RawQuery)
	assert.Equal(t, "", resp.Request.Header.Get("Content-Type"))
}

func TestRequestHandler_GetWithQueryData(t *testing.T) {
	mockServer := newMockServer(t)
	defer mockServer.Close()

	headers := map[string]interface{}{}
	queryData := url.Values{
		"status": []string{"closed"},
	}

	requestHandler := NewRequestHandler("user", "pass")
	resp, err := requestHandler.Get(mockServer.URL, queryData, headers)

	assert.NoError(t, err)
	assert.Equal(t, "GET", resp.Request.Method)
	assert.Equal(t, "status=closed", resp.Request.URL.RawQuery)
	assert.Equal(t, "", resp.Request.Header.Get("Content-Type"))
}

func TestRequestHandler_Delete(t *testing.T) {
	mockServer := newMockServer(t)
	defer mockServer.Close()

	headers := map[string]interface{}{}

	requestHandler := NewRequestHandler("user", "pass")
	resp, err := requestHandler.Delete(mockServer.URL, nil, headers)

	assert.NoError(t, err)
	assert.Equal(t, resp.Request.Method, "DELETE")
	assert.Equal(t, "", resp.Request.Header.Get("Content-Type"))
}

func newMockServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			data := map[string]interface{}{
				"response": "ok",
			}
			if err := json.NewEncoder(writer).Encode(&data); err != nil {
				t.Error(err)
			}
		}))
}
