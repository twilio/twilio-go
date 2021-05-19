package twilio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_WithParams(t *testing.T) {
	client := NewRestClientWithParams("parentSid", "authToken", RestClientParams{
		AccountSid: "subAccountSid",
	})

	assert.Equal(t, client.RequestHandler.Client.AccountSid(), "subAccountSid")
}

func TestClient_WithNoAccountSid(t *testing.T) {
	client := NewRestClient("parentSid", "authToken")
	assert.Equal(t, client.RequestHandler.Client.AccountSid(), "parentSid")
}
