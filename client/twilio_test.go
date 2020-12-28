package client

import (
	"github.com/stretchr/testify/assert"
	"github.com/twilio/twilio-go/framework/error"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_SendRequestError(t *testing.T) {
	errorResponse := `{
	"status": 400,
	"code":20001,
	"message":"Bad request",
	"more_info":"https://www.twilio.com/docs/errors/20001"
}`
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(resp http.ResponseWriter, req *http.Request) {
			resp.WriteHeader(400)
			_, _ = resp.Write([]byte(errorResponse))
		}))
	defer mockServer.Close()

	client := NewClient("user", "pass")
	resp, err := client.SendRequest("get", mockServer.URL, nil, nil)
	twilioError := err.(*error.TwilioRestError)
	assert.Nil(t, resp)
	assert.Equal(t, 400, twilioError.Status)
	assert.Equal(t, 20001, twilioError.Code)
	assert.Equal(t, "https://www.twilio.com/docs/errors/20001", twilioError.MoreInfo)
	assert.Equal(t, "Bad request", twilioError.Message)
	assert.Nil(t, twilioError.Details)
}

func TestClient_SendRequestErrorWithDetails(t *testing.T) {
	errorResponse := []byte(`{
	"status": 400,
	"message": "Bad request",
	"code": 20001,
	"more_info": "https://www.twilio.com/docs/errors/20001",
	"details": {
		"foo": "bar"
	}
}`)
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(resp http.ResponseWriter, req *http.Request) {
			resp.WriteHeader(400)
			_, _ = resp.Write(errorResponse)
		}))
	defer mockServer.Close()

	client := NewClient("user", "pass")
	resp, err := client.SendRequest("get", mockServer.URL, nil, nil)
	twilioError := err.(*error.TwilioRestError)
	details := make(map[string]interface{})
	details["foo"] = "bar"
	assert.Nil(t, resp)
	assert.Equal(t, 400, twilioError.Status)
	assert.Equal(t, 20001, twilioError.Code)
	assert.Equal(t, "https://www.twilio.com/docs/errors/20001", twilioError.MoreInfo)
	assert.Equal(t, "Bad request", twilioError.Message)
	assert.Equal(t, details, twilioError.Details)
}
