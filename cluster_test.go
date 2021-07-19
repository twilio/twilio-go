// +build cluster

package twilio

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
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
	params := &openapi.CreateMessageParams{}
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
	resp, err := testClient.ApiV2010.ListIncomingPhoneNumber(nil, 0)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	// from, to numbers plus any other numbers that's configured for the account.
	assert.GreaterOrEqual(t, len(resp), 2)
}

func TestListingANumber(t *testing.T) {
	params := &openapi.ListIncomingPhoneNumberParams{}
	params.SetPhoneNumber(from)

	resp, err := testClient.ApiV2010.ListIncomingPhoneNumber(params, 0)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, len(resp))
	assert.Equal(t, from, *resp[0].PhoneNumber)
}
