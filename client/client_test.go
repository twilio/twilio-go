package client_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	twilio "github.com/twilio/twilio-go/client"
)

func NewClient(accountSid string, authToken string) *twilio.Client {
	c := &twilio.Client{
		Credentials: twilio.NewCredentials(accountSid, authToken),
		HTTPClient:  http.DefaultClient,
	}

	return c
}

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
	resp, err := client.SendRequest("GET", mockServer.URL, nil, nil) //nolint:bodyclose
	twilioError := err.(*twilio.TwilioRestError)
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
	resp, err := client.SendRequest("GET", mockServer.URL, nil, nil) //nolint:bodyclose
	twilioError := err.(*twilio.TwilioRestError)
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
			_, _ = writer.Write([]byte(`{"redirect_to": "some_place"}`))
		}))
	defer mockServer.Close()

	client := NewClient("user", "pass")
	resp, _ := client.SendRequest("GET", mockServer.URL, nil, nil) //nolint:bodyclose
	assert.Equal(t, 307, resp.StatusCode)
}

func TestClient_SetTimeoutTimesOut(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			d := map[string]interface{}{
				"response": "ok",
			}
			time.Sleep(100 * time.Microsecond)
			encoder := json.NewEncoder(writer)
			err := encoder.Encode(&d)
			if err != nil {
				t.Error(err)
			}
			writer.WriteHeader(http.StatusOK)
		}))
	defer mockServer.Close()

	client := NewClient("user", "pass")
	client.SetTimeout(10 * time.Microsecond)
	_, err := client.SendRequest("GET", mockServer.URL, nil, nil) //nolint:bodyclose
	assert.Error(t, err)
}

func TestClient_SetTimeoutSucceeds(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			d := map[string]interface{}{
				"response": "ok",
			}
			time.Sleep(100 * time.Microsecond)
			encoder := json.NewEncoder(writer)
			err := encoder.Encode(&d)
			if err != nil {
				t.Error(err)
			}
		}))
	defer mockServer.Close()

	client := NewClient("user", "pass")
	client.SetTimeout(10 * time.Second)
	resp, err := client.SendRequest("GET", mockServer.URL, nil, nil) //nolint:bodyclose
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestClient_SetTimeoutCreatesClient(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			d := map[string]interface{}{
				"response": "ok",
			}
			time.Sleep(100 * time.Microsecond)
			encoder := json.NewEncoder(writer)
			err := encoder.Encode(&d)
			if err != nil {
				t.Error(err)
			}
		}))
	defer mockServer.Close()

	client := &twilio.Client{
		Credentials: twilio.NewCredentials("user", "pass"),
	}
	client.SetTimeout(20 * time.Second)
	resp, err := client.SendRequest("GET", mockServer.URL, nil, nil) //nolint:bodyclose
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestClient_UnicodeResponse(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			d := map[string]interface{}{
				"testing-unicode": "Ω≈ç√, 💩",
			}
			encoder := json.NewEncoder(writer)
			err := encoder.Encode(&d)
			if err != nil {
				t.Error(err)
			}
		}))
	defer mockServer.Close()

	client := NewClient("user", "pass")
	resp, _ := client.SendRequest("GET", mockServer.URL, nil, nil) //nolint:bodyclose
	assert.Equal(t, 200, resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, "{\"testing-unicode\":\"Ω≈ç√, 💩\"}\n", string(body))
}

func TestClient_RequestBodyShouldContainJson(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			bytes, err := ioutil.ReadAll(request.Body)
			assert.NoError(t, err)
			assert.Equal(t, `{"account_sid":"ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","chat_service_instance_sid":"ISxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`, string(bytes))

			writer.WriteHeader(200)
		}),
	)
	defer mockServer.Close()

	client := &twilio.Client{
		Credentials: twilio.NewCredentials("user", "pass"),
	}

	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}
	body := map[string]interface{}{
		"account_sid":               "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"chat_service_instance_sid": "ISxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	}

	resp, err := client.SendRequest("POST", mockServer.URL, body, headers) //nolint:bodyclose
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	// Request body is asserted in the mock server. This can't be asserted here because the body has been consumed and will be nil
}

func TestClient_RequestBodyShouldContainFormData(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			bytes, err := ioutil.ReadAll(request.Body)
			assert.NoError(t, err)
			assert.Equal(t, "AccountSid=ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx&ChatServiceInstanceSid=ISxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", string(bytes))

			writer.WriteHeader(200)
		}),
	)
	defer mockServer.Close()

	client := &twilio.Client{
		Credentials: twilio.NewCredentials("user", "pass"),
	}

	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	body := url.Values{
		"AccountSid":             []string{"ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"},
		"ChatServiceInstanceSid": []string{"ISxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"},
	}

	resp, err := client.SendRequest("POST", mockServer.URL, body, headers) //nolint:bodyclose
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	// Request body is asserted in the mock server. This can't be asserted here because the body has been consumed and will be nil
}

func TestClient_ShouldThrowErrorWhenNoContentTypeHeaderIsPresentOnAPostRequest(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(400) // This shouldn't be hit as an error should have already been returned
		}),
	)
	defer mockServer.Close()

	client := &twilio.Client{
		Credentials: twilio.NewCredentials("user", "pass"),
	}

	headers := map[string]interface{}{}
	body := url.Values{
		"account_sid":               []string{"ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"},
		"chat_service_instance_sid": []string{"ISxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"},
	}

	_, err := client.SendRequest("POST", mockServer.URL, body, headers) //nolint:bodyclose
	assert.EqualError(t, err, "the 'Content-Type' header must be set on a POST request")
}

func TestClient_ShouldMakeGetRequestWithQueryStringParams(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			d := map[string]interface{}{
				"sid":          "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
				"status":       "closed",
				"date_updated": "2021-06-01T12:00:00Z",
			}
			encoder := json.NewEncoder(writer)
			err := encoder.Encode(&d)
			if err != nil {
				t.Error(err)
			}
		}))
	defer mockServer.Close()

	client := &twilio.Client{
		Credentials: twilio.NewCredentials("user", "pass"),
	}

	queryParams := url.Values{
		"status": []string{"closed"},
	}

	resp, err := client.SendRequest("GET", mockServer.URL, queryParams, nil) //nolint:bodyclose
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, resp.Request.URL.RawQuery, "status=closed")
}
