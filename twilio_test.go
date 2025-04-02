package twilio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_WithParams(t *testing.T) {
	client := NewRestClientWithParams(ClientParams{
		Username:   "parentSid",
		Password:   "authToken",
		AccountSid: "subAccountSid",
	})

	assert.Equal(t, client.RequestHandler.Client.AccountSid(), "subAccountSid")
}

func TestClient_WithNoAccountSid(t *testing.T) {
	client := NewRestClientWithParams(ClientParams{
		Username: "parentSid",
		Password: "authToken",
	})
	assert.Equal(t, client.RequestHandler.Client.AccountSid(), "parentSid")
}

func TestClientCredentialProvider(t *testing.T) {
	creds := ClientCredentialProvider{
		GrantType:    "client_credentials",
		ClientId:     "mock_client_id",
		ClientSecret: "mock_client_secret",
	}
	client := NewRestClientWithParams(ClientParams{
		Username:                 "parentSid",
		Password:                 "authToken",
		ClientCredentialProvider: &creds,
	},
	)

	assert.Equal(t, "client_credentials", client.Client.OAuth().(*APIOAuth).creds.GrantType)
	assert.Equal(t, "mock_client_id", client.Client.OAuth().(*APIOAuth).creds.ClientId)
	assert.Equal(t, "mock_client_secret", client.Client.OAuth().(*APIOAuth).creds.ClientSecret)
}
