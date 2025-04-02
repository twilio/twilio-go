//go:build cluster
// +build cluster

package twilio

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	Api "github.com/twilio/twilio-go/rest/api/v2010"
)

var from string
var to string
var testClient *RestClient

func TestMain(m *testing.M) {
	from = os.Getenv("TWILIO_FROM_NUMBER_OAUTH")
	to = os.Getenv("TWILIO_TO_NUMBER_OAUTH")
	var clientId = os.Getenv("TWILIO_CLIENT_ID")
	var clientSecret = os.Getenv("TWILIO_CLIENT_SECRET")
	var accountSid = os.Getenv("TWILIO_ACCOUNT_SID_OAUTH")

	clientCredentialProvider := &ClientCredentialProvider{
		GrantType:    "client_credentials",
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
	testClient = NewRestClientWithParams(
		ClientParams{ClientCredentialProvider: clientCredentialProvider, AccountSid: accountSid},
	)
	ret := m.Run()
	os.Exit(ret)
}

func TestClientWithParams(t *testing.T) {
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
