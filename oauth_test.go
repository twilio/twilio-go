package twilio

import (
	"context"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/twilio/twilio-go/client"
)

type MockOAuth struct{}

func (m *MockOAuth) GetAccessToken(ctx context.Context) (string, error) {
	return "mock_token", nil
}

type MockTokenAuth struct {
	token string
	OAuth client.OAuth
}

func (m *MockTokenAuth) FetchToken(ctx context.Context) (string, error) {
	t := &TokenAuth{token: m.token, OAuth: m.OAuth}
	return t.FetchToken(ctx)
}

func (m *MockTokenAuth) TokenExpired() bool {
	return false
}

func TestTokenAuth_FetchToken(t *testing.T) {
	oauth := &MockOAuth{}
	tokenAuth := &TokenAuth{OAuth: oauth}

	// Test fetching a new token
	token, err := tokenAuth.FetchToken(t.Context())
	assert.NoError(t, err)
	assert.Equal(t, "mock_token", token)

	// Test fetching the cached token
	mockTokenAuth := MockTokenAuth{OAuth: oauth, token: "cached_token"}
	token, err = mockTokenAuth.FetchToken(t.Context())
	assert.NoError(t, err)
	assert.Equal(t, "cached_token", token)
}

func TestTokenAuth_Expired(t *testing.T) {
	// Test with an expired token
	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(-time.Hour).UTC().Unix(),
	})
	expiredTokenString, err := expiredToken.SignedString([]byte("secret"))
	assert.NoError(t, err)

	tokenAuth := &TokenAuth{token: expiredTokenString}
	hasExpired, err := tokenAuth.Expired(t.Context())
	assert.True(t, hasExpired)
	assert.NoError(t, err)

	// Test with a valid token
	validToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).UTC().Unix(),
	})
	validTokenString, _ := validToken.SignedString([]byte("secret"))

	tokenAuth.token = validTokenString
	hasExpired, err = tokenAuth.Expired(t.Context())
	assert.False(t, hasExpired)
	assert.NoError(t, err)
}
