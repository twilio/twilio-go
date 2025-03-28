package twilio

import (
	"context"
	"github.com/twilio/twilio-go/client"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

type MockOAuth struct{}

func (m *MockOAuth) GetAccessToken(ctx context.Context) (string, error) {
	return "mock_token", nil
}

type MockTokenAuth struct {
	token string
	OAuth client.OAuth
}

func (m *MockTokenAuth) FetchToken() (string, error) {
	t := &TokenAuth{token: m.token, OAuth: m.OAuth}
	return t.FetchToken()
}

func (m *MockTokenAuth) TokenExpired() bool {
	return false
}

func TestTokenAuth_FetchToken(t *testing.T) {
	oauth := &MockOAuth{}
	tokenAuth := &TokenAuth{OAuth: oauth}

	// Test fetching a new token
	token, err := tokenAuth.FetchToken()
	assert.NoError(t, err)
	assert.Equal(t, "mock_token", token)

	// Test fetching the cached token
	mockTokenAuth := MockTokenAuth{OAuth: oauth, token: "cached_token"}
	token, err = mockTokenAuth.FetchToken()
	assert.NoError(t, err)
	assert.Equal(t, "cached_token", token)
}

func TestTokenAuth_Expired(t *testing.T) {
	// Test with an expired token
	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(-time.Hour).UTC().Unix(),
	})
	expiredTokenString, _ := expiredToken.SignedString([]byte("secret"))

	tokenAuth := &TokenAuth{token: expiredTokenString}
	hasExpired, err := tokenAuth.Expired()
	assert.True(t, hasExpired)
	assert.NoError(t, err)

	// Test with a valid token
	validToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).UTC().Unix(),
	})
	validTokenString, _ := validToken.SignedString([]byte("secret"))

	tokenAuth.token = validTokenString
	hasExpired, err = tokenAuth.Expired()
	assert.False(t, hasExpired)
	assert.NoError(t, err)
}
