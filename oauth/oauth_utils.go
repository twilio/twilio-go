package oauth

import (
	"fmt"

	PreviewIamTemp "github.com/twilio/twilio-go/rest/preview_iam/temp"
)

type OrgTokenManager struct {
	GrantType    string
	ClientId     string
	ClientSecret string
	Code         string
	RedirectUri  string
	Audience     string
	RefreshToken string
	Scope        string
}

type ClientTokenManager struct {
	GrantType    string
	ClientId     string
	ClientSecret string
	Code         string
	RedirectUri  string
	Audience     string
	RefreshToken string
	Scope        string
}

func GetOrgAccessToken(manager OrgTokenManager) (*PreviewIamTemp.OauthV1Token, error) {
	params := &PreviewIamTemp.CreateTokenParams{}
	params.SetGrantType(manager.GrantType)
	params.SetClientId(manager.ClientId)
	params.SetClientSecret(manager.ClientSecret)
	params.SetCode(manager.Code)
	params.SetRedirectUri(manager.RedirectUri)
	params.SetAudience(manager.Audience)
	params.SetRefreshToken(manager.RefreshToken)
	params.SetScope(manager.Scope)

	token, err := PreviewIamTemp.NewApiService(NewRestClient().RequestHandler).CreateToken(params)
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	return token, nil
}

func GetClientAccessToken(manager ClientTokenManager) (*PreviewIamTemp.OauthV1Token, error) {
	params := &PreviewIamTemp.CreateTokenParams{}
	params.SetGrantType(manager.GrantType)
	params.SetClientId(manager.ClientId)
	params.SetClientSecret(manager.ClientSecret)
	params.SetCode(manager.Code)
	params.SetRedirectUri(manager.RedirectUri)
	params.SetAudience(manager.Audience)
	params.SetRefreshToken(manager.RefreshToken)
	params.SetScope(manager.Scope)

	token, err := PreviewIamTemp.NewApiService(NewRestClient().RequestHandler).CreateToken(params)
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	return token, nil
}
