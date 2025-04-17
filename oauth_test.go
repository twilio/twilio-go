package twilio

import (
	"context"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestNewTokenAuth(t *testing.T) {
	token := "test-token"
	auth := &TokenAuth{Token: token}
	assert.Equal(t, token, auth.Token)
}

func TestFetchToken_ValidToken(t *testing.T) {
	token := createTestToken(time.Now().Add(1 * time.Hour))
	auth := &TokenAuth{Token: token}
	fetchedToken, err := auth.FetchToken(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, token, fetchedToken)
}

func TestFetchToken_ExpiredToken(t *testing.T) {
	token := createTestToken(time.Now().Add(-1 * time.Hour))
	auth := &TokenAuth{Token: token}
	fetchedToken, err := auth.FetchToken(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, "", fetchedToken)
}

func TestExpired_ValidToken(t *testing.T) {
	token := createTestToken(time.Now().Add(1 * time.Hour))
	auth := &TokenAuth{Token: token}
	expired, err := auth.Expired(context.Background())
	assert.Nil(t, err)
	assert.False(t, expired)
}

func TestExpired_ExpiredToken(t *testing.T) {
	token := createTestToken(time.Now().Add(-1 * time.Hour))
	auth := &TokenAuth{Token: token}
	expired, err := auth.Expired(context.Background())
	assert.Nil(t, err)
	assert.True(t, expired)
}

func createTestToken(expirationTime time.Time) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": expirationTime.Unix(),
	})
	tokenString, _ := token.SignedString([]byte("test-secret"))
	return tokenString
}
