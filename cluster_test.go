// +build cluster

package twilio

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	ApiV2010 "github.com/twilio/twilio-go/rest/api/v2010"
	ChatV2 "github.com/twilio/twilio-go/rest/chat/v2"
	EventsV1 "github.com/twilio/twilio-go/rest/events/v1"
)

var from string
var to string
var testClient *RestClient

func TestMain(m *testing.M) {
	from = os.Getenv("TWILIO_FROM_NUMBER")
	to = os.Getenv("TWILIO_TO_NUMBER")

	testClient = NewRestClient()
	ret := m.Run()
	os.Exit(ret)
}

func TestSendingAText(t *testing.T) {
	params := &ApiV2010.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody("Hello there")

	resp, err := testClient.ApiV2010.CreateMessage(params)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Hello there", *resp.Body)
	assert.Equal(t, from, *resp.From)
	assert.Equal(t, to, *resp.To)
}

func TestListingNumbers(t *testing.T) {
	resp, err := testClient.ApiV2010.ListIncomingPhoneNumber(nil)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	// from, to numbers plus any other numbers that's configured for the account.
	assert.GreaterOrEqual(t, len(resp), 2)
}

func TestListingANumber(t *testing.T) {
	params := &ApiV2010.ListIncomingPhoneNumberParams{}
	params.SetPhoneNumber(from)

	resp, err := testClient.ApiV2010.ListIncomingPhoneNumber(params)
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

	types := []map[string]interface{}{
		{
			"type": "com.twilio.messaging.message.delivered",
		},
		{
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
