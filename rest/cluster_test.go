package rest

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

var from string
var to string
var client *twilio.RestClient

func TestMain(m *testing.M) {
	from = os.Getenv("TWILIO_FROM_NUMBER")
	to = os.Getenv("TWILIO_TO_NUMBER")

	client = twilio.NewRestClient()
	ret := m.Run()
	os.Exit(ret)
}

func TestSendingAText(t *testing.T) {
	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody("Hello there")

	resp, err := client.ApiV2010.CreateMessage(params)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Hello there", *resp.Body)
	assert.Equal(t, from, *resp.From)
	assert.Equal(t, to, *resp.To)
}

func TestListingNumbers(t *testing.T) {
	resp, err := client.ApiV2010.ListIncomingPhoneNumber(nil, 0)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.GreaterOrEqual(t, len(resp), 2)
}

func TestListingANumber(t *testing.T) {
	params := &openapi.ListIncomingPhoneNumberParams{}
	params.SetPhoneNumber(from)

	resp, err := client.ApiV2010.ListIncomingPhoneNumber(params, 0)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, len(resp))
	assert.Equal(t, from, *resp[0].PhoneNumber)
}

func TestCreateAndDeleteNumber(t *testing.T) {
	createParams := &openapi.CreateIncomingPhoneNumberParams{}
	createParams.SetAreaCode("415")
	resp, err := client.ApiV2010.CreateIncomingPhoneNumber(createParams)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	sid := *resp.Sid
	err = client.ApiV2010.DeleteIncomingPhoneNumber(sid, nil)
	assert.Nil(t, err)
}
