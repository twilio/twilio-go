//go:build cluster
// +build cluster

package twilio

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	Api "github.com/twilio/twilio-go/rest/api/v2010"
	ChatV2 "github.com/twilio/twilio-go/rest/chat/v2"
	EventsV1 "github.com/twilio/twilio-go/rest/events/v1"
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

func TestListingAvailableNumber(t *testing.T) {
	params := &Api.ListAvailablePhoneNumberTollFreeParams{}
	params.SetLimit(2)

	resp, err := testClient.Api.ListAvailablePhoneNumberTollFree("US", params)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 2, len(resp))
}
