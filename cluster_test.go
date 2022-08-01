//go:build cluster
// +build cluster

package twilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/localtunnel/go-localtunnel"
	"github.com/stretchr/testify/assert"
	twilio "github.com/twilio/twilio-go/client"
	Api "github.com/twilio/twilio-go/rest/api/v2010"
	ChatV2 "github.com/twilio/twilio-go/rest/chat/v2"
	EventsV1 "github.com/twilio/twilio-go/rest/events/v1"
	StudioV2 "github.com/twilio/twilio-go/rest/studio/v2"
)

var from string
var to string
var testClient *RestClient

func TestMain(m *testing.M) {
	from = os.Getenv("TWILIO_FROM_NUMBER")
	to = os.Getenv("TWILIO_TO_NUMBER")
	var apiKey = os.Getenv("TWILIO_API_KEY")
	var secret = os.Getenv("TWILIO_API_SECRET")
	var accountSid = os.Getenv("TWILIO_ACCOUNT_SID")

	testClient = NewRestClientWithParams(ClientParams{apiKey, secret, accountSid, nil})
	ret := m.Run()
	os.Exit(ret)
}

func TestSendingAText(t *testing.T) {
	params := &Api.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody("Hello there")

	resp, err := testClient.Api.CreateMessage(params)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Hello there", *resp.Body)
	assert.Equal(t, from, *resp.From)
	assert.Equal(t, to, *resp.To)
}

func TestListingNumbers(t *testing.T) {
	resp, err := testClient.Api.ListIncomingPhoneNumber(nil)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	// from, to numbers plus any other numbers that's configured for the account.
	assert.GreaterOrEqual(t, len(resp), 2)
}

func TestListingANumber(t *testing.T) {
	params := &Api.ListIncomingPhoneNumberParams{}
	params.SetPhoneNumber(from)

	resp, err := testClient.Api.ListIncomingPhoneNumber(params)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, len(resp))
	assert.Equal(t, from, *resp[0].PhoneNumber)
}

func TestSpecialCharacters(t *testing.T) {
	serviceParams := &ChatV2.CreateServiceParams{}
	serviceParams.SetFriendlyName("service|friendly&name")

	service, err := testClient.ChatV2.CreateService(serviceParams)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	userParams := &ChatV2.CreateUserParams{}
	userParams.SetIdentity("user|identity&string")

	user, err := testClient.ChatV2.CreateUser(*service.Sid, userParams)
	assert.Nil(t, err)
	assert.NotNil(t, user)

	err = testClient.ChatV2.DeleteUser(*service.Sid, *user.Identity)
	assert.Nil(t, err)
	err = testClient.ChatV2.DeleteService(*service.Sid)
	assert.Nil(t, err)
}

func TestListParams(t *testing.T) {
	sinkConfig := map[string]interface{}{
		"destination":  "http://example.org/webhook",
		"method":       "post",
		"batch_events": false,
	}
	sinkParams := &EventsV1.CreateSinkParams{}
	sinkParams.SetSinkConfiguration(sinkConfig)
	sinkParams.SetDescription("test sink")
	sinkParams.SetSinkType("webhook")
	sink, err := testClient.EventsV1.CreateSink(sinkParams)
	assert.Nil(t, err)
	assert.NotNil(t, sink)

	types := []interface{}{
		map[string]interface{}{
			"type": "com.twilio.messaging.message.delivered",
		},
		map[string]interface{}{
			"type": "com.twilio.messaging.message.sent",
		},
	}
	subscriptionsParams := &EventsV1.CreateSubscriptionParams{}
	subscriptionsParams.SetSinkSid(*sink.Sid)
	subscriptionsParams.SetTypes(types)
	subscriptionsParams.SetDescription("test subscription")
	subscription, err := testClient.EventsV1.CreateSubscription(subscriptionsParams)
	assert.Nil(t, err)
	assert.NotNil(t, subscription)

	// Clean up the resources we've created
	err = testClient.EventsV1.DeleteSubscription(*subscription.Sid)
	assert.Nil(t, err)
	err = testClient.EventsV1.DeleteSink(*sink.Sid)
	assert.Nil(t, err)
}

func createValidationServer(channel chan bool) *http.Server {
	server := &http.Server{}
	server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.Header["X-Forwarded-Proto"][0] + "://" + r.Header["X-Forwarded-Host"][0] + r.URL.RequestURI()

		//Converting url.Values to map for POST Request Params
		r.ParseForm()
		params := make(map[string]string)
		for k, v := range r.PostForm {
			params[k] = v[0]
		}
		requestValidator := twilio.NewRequestValidator(os.Getenv("TWILIO_AUTH_TOKEN"))
		channel <- requestValidator.Validate(url, params, r.Header["X-Twilio-Signature"][0])
		defer server.Close()
	})
	return server
}

func createStudioFlowParams(url string) *StudioV2.CreateFlowParams {
	jsonStr := fmt.Sprintf(`{
		"description": "Studio Flow",
		"states": [
		  {
			"name": "Trigger",
			"type": "trigger",
			"transitions": [
			  {
				"next": "httpRequest",
				"event": "incomingRequest"
			  }
			],
			"properties": {
			}
		  },
		  {
			"name": "httpRequest",
			"type": "make-http-request",
			"transitions": [],
			"properties": {
			  "method": "GET",
			  "content_type": "application/x-www-form-urlencoded;charset=utf-8",
			  "url": "%s"
			}
		  }
		],
		"initial_state": "Trigger",
		"flags": {
		  "allow_concurrent_calls": true
		}
	  }`, url)

	var definition interface{}
	_ = json.Unmarshal([]byte(jsonStr), &definition)

	params := &StudioV2.CreateFlowParams{
		Definition: &definition,
	}
	params.SetFriendlyName("Go Cluster Test Flow")
	params.SetStatus("published")
	return params
}

func createStudioExecutionParams() *StudioV2.CreateExecutionParams {
	executionParams := &StudioV2.CreateExecutionParams{}
	executionParams.SetTo(to)
	executionParams.SetFrom(from)
	return executionParams
}

func TestRequestValidation(t *testing.T) {
	//Invoke Localtunnel
	listener, ltErr := localtunnel.Listen(localtunnel.Options{})
	if ltErr != nil {
		t.Fatal("Error with Localtunnel: ", ltErr)
	}
	//Create Studio Flow
	params := createStudioFlowParams(listener.URL())
	resp, flowErr := testClient.StudioV2.CreateFlow(params)
	if flowErr != nil {
		t.Fatal("Error with Studio Flow Creation: ", flowErr)
	}
	flowSid := *resp.Sid
	channel := make(chan bool)

	//Create Validation Server & Listen
	server := createValidationServer(channel)
	go server.Serve(listener)

	//Extra time for Server to set up
	time.Sleep(1 * time.Second)

	//Trigger Studio Flow
	_, exeErr := testClient.StudioV2.CreateExecution(flowSid, createStudioExecutionParams())
	if exeErr != nil {
		t.Fatal("Error with Studio Execution Creation: ", exeErr)
	}

	//Await for Request Validation
	assert.True(t, <-channel)
	defer testClient.StudioV2.DeleteFlow(flowSid)
}
