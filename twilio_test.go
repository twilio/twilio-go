package twilio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:paralleltest
func TestClient_WithParams(t *testing.T) {
	client := NewRestClientWithParams("parentSid", "authToken", RestClientParams{
		AccountSid: "subAccountSid",
	})

	assert.Equal(t, client.Username, "parentSid")
	assert.Equal(t, client.Password, "authToken")
	assert.Equal(t, client.AccountSid, "subAccountSid")
}

//nolint:paralleltest
func TestClient_WithNoAccountSid(t *testing.T) {
	client := NewRestClient("parentSid", "authToken")
	assert.Equal(t, client.Username, "parentSid")
	assert.Equal(t, client.AccountSid, "parentSid")
}
