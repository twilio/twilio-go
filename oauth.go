package twilio

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/twilio/twilio-go/client"
	iam "github.com/twilio/twilio-go/rest/iam/v1"
)

type TokenAuth struct {
	token string
	OAuth client.OAuth
}

func (t *TokenAuth) NewTokenAuth(token string, o client.OAuth) *TokenAuth {
	return &TokenAuth{token: token, OAuth: o}
}

func (t *TokenAuth) FetchToken() (string, error) {
	if t.token != "" && !t.Expired() {
		return t.token, nil
	}

	token, err := t.OAuth.GetAccessToken(context.Background())
	if err != nil {
		return "", err
	}

	t.token = token
	return t.token, nil
}

func (t *TokenAuth) Expired() bool {
	token, _, err := new(jwt.Parser).ParseUnverified(t.token, jwt.MapClaims{})
	if err != nil {
		return true
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			expirationTime := time.Unix(int64(exp), 0).UTC()
			return time.Now().UTC().After(expirationTime)
		}
	}
	return true
}

type OAuthCredentials struct {
	// TODO: documentation.
	GrantType              string
	ClientId, ClientSecret string
}

type APIOAuth struct {
	iamService *iam.ApiService
	creds      *OAuthCredentials
}

func NewAPIOAuth(c *client.RequestHandler, creds *OAuthCredentials) *APIOAuth {
	a := &APIOAuth{iamService: iam.NewApiService(c), creds: creds}
	a.iamService.RequestHandler().Client.SetOauth(nil)
	return a
}

func (a *APIOAuth) GetAccessToken(ctx context.Context) (string, error) {
	if a == nil {
		panic("twilio: API OAuth object is nil")
	}
	if a.creds == nil {
		panic("twilio: API OAuth credentials are nil")
	}
	params := &iam.CreateTokenParams{}
	params.SetGrantType(a.creds.GrantType).
		SetClientId(a.creds.ClientId).
		SetClientSecret(a.creds.ClientSecret)
	token, err := a.iamService.CreateToken(params)
	if err != nil {
		return "", err
	}
	if token.AccessToken == nil {
		return "", fmt.Errorf("twilio: API response to create a token did not return a valid token")
	}

	return *token.AccessToken, nil
}
