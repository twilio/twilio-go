package twilio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:paralleltest
func TestClient_WithAccountSid(t *testing.T) {
	client := NewClientWithAccountSid("parentSid", "authToken", "subAccountSid")
	assert.Equal(t, client.Username, "parentSid")
	assert.Equal(t, client.Password, "authToken")
	assert.Equal(t, client.AccountSid, "subAccountSid")
}

//nolint:paralleltest
func TestClient_WithNoAccountSid(t *testing.T) {
	client := NewClient("parentSid", "authToken")
	assert.Equal(t, client.Username, "parentSid")
	assert.Equal(t, client.AccountSid, "parentSid")
}
