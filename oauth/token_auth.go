package oauth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/twilio/twilio-go/client"
)

type TokenAuth struct {
	token        string
	tokenManager *TokenManager
}

func NewTokenAuthInitializer(token string, tokenManager *TokenManager) *TokenAuth {
	return &TokenAuth{
		token:        token,
		tokenManager: tokenManager,
	}
}
func NewTokenAuth(grantType, clientId, clientSecret, code, audience, refreshToken, scope string) *TokenAuth {
	return &TokenAuth{
		tokenManager: &TokenManager{
			GrantType:    grantType,
			ClientId:     clientId,
			ClientSecret: clientSecret,
			Code:         code,
			Audience:     audience,
			RefreshToken: refreshToken,
			Scope:        scope,
		},
	}
}

func (t *TokenAuth) FetchToken(c client.RequestHandler) (string, error) {
	if t.token != "" && !t.TokenExpired() {
		return t.token, nil
	}

	token, err := t.tokenManager.fetchAccessToken(c)
	if err != nil {
		return "", err
	}

	t.token = token
	return t.token, nil
}

func (t *TokenAuth) TokenExpired() bool {
	token, _, err := new(jwt.Parser).ParseUnverified(t.token, jwt.MapClaims{})
	if err != nil {
		return true
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			expirationTime := time.Unix(int64(exp), 0)
			return time.Now().After(expirationTime)
		}
	}
	return true
}
