package oauth

import (
	"fmt"

	"github.com/twilio/twilio-go/client"
	preview_iam "github.com/twilio/twilio-go/rest/preview_iam/v1"
)

type TokenManager struct {
	GrantType    string
	ClientId     string
	ClientSecret string
	Code         string
	Audience     string
	RefreshToken string
	Scope        string
}

func NewTokenManager(grantType, clientId, clientSecret, code, audience, refreshToken, scope string) *TokenManager {
	return &TokenManager{
		GrantType:    grantType,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Code:         code,
		Audience:     audience,
		RefreshToken: refreshToken,
		Scope:        scope,
	}
}

func (tm *TokenManager) fetchAccessToken(c client.RequestHandler) (string, error) {
	params := &preview_iam.CreateTokenParams{}
	params.SetGrantType(tm.GrantType).
		SetClientId(tm.ClientId).
		SetClientSecret(tm.ClientSecret).
		SetCode(tm.Code).
		SetAudience(tm.Audience).
		SetRefreshToken(tm.RefreshToken).
		SetScope(tm.Scope)

	token, err := preview_iam.NewApiService(&c).CreateToken(params)
	if err != nil {
		return "", err
	}

	if token.AccessToken == nil {
		return "", fmt.Errorf("access token is nil")
	}

	return *token.AccessToken, nil
}
