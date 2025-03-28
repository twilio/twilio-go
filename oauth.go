package twilio

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/twilio/twilio-go/client"
	iam "github.com/twilio/twilio-go/rest/iam/v1"
)

// TokenAuth handles token-based authentication using OAuth.
type TokenAuth struct {
	// token is the cached OAuth token.
	token string
	// OAuth is the OAuth client used to fetch new tokens.
	OAuth client.OAuth
}

// NewTokenAuth creates a new TokenAuth instance with the provided token and OAuth client.
func (t *TokenAuth) NewTokenAuth(token string, o client.OAuth) *TokenAuth {
	return &TokenAuth{token: token, OAuth: o}
}

// FetchToken retrieves the current token if it is valid, or fetches a new token using the OAuth client.
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

// Expired checks if the current token is expired.
func (t *TokenAuth) Expired() bool {
	token, _, err := new(jwt.Parser).ParseUnverified(t.token, jwt.MapClaims{})
	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err)
		return true
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Error asserting token claims")
		return true
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		fmt.Println("Expiration time (exp) not found in token claims")
		return true
	}

	expirationTime := time.Unix(int64(exp), 0).UTC()
	return time.Now().UTC().After(expirationTime)
}

// OAuthCredentials holds the necessary credentials for OAuth authentication.
type OAuthCredentials struct {
	// GrantType specifies the type of grant being used for OAuth.
	GrantType string
	// ClientId is the identifier for the client application. ClientSecret is the secret key for the client application.
	ClientId, ClientSecret string
}

// APIOAuth handles OAuth authentication for the Twilio API.
type APIOAuth struct {
	// iamService is the service used to interact with the IAM API.
	iamService *iam.ApiService
	// creds holds the necessary credentials for OAuth authentication.
	creds *OAuthCredentials
}

// NewAPIOAuth creates a new APIOAuth instance with the provided request handler and credentials.
func NewAPIOAuth(c *client.RequestHandler, creds *OAuthCredentials) *APIOAuth {
	a := &APIOAuth{iamService: iam.NewApiService(c), creds: creds}
	a.iamService.RequestHandler().Client.SetOauth(nil)
	return a
}

// GetAccessToken retrieves an access token using the OAuth credentials.
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
