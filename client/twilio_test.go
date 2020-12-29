package client_test

import (
	"github.com/stretchr/testify/assert"
	twilio "github.com/twilio/twilio-go/client"
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

	client := twilio.NewClient("user", "pass")
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

	client := twilio.NewClient("user", "pass")
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

func TestClient_SendRequestWithRedirect(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(307)
			writer.Write([]byte(`{"redirect_to": "some_place"}`))
		}))
	defer mockServer.Close()

	client := NewClient("user", "pass")
	resp, _ := client.SendRequest("get", mockServer.URL, nil, nil)
	assert.Equal(t, 307, resp.StatusCode)
}
