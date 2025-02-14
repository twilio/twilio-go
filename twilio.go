/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
// Package twilio provides bindings for Twilio's REST APIs.
package twilio

import (
	"fmt"
	"os"
	"time"

	"github.com/twilio/twilio-go/client"
	AccountsV1 "github.com/twilio/twilio-go/rest/accounts/v1"
	Api "github.com/twilio/twilio-go/rest/api/v2010"
	AssistantsV1 "github.com/twilio/twilio-go/rest/assistants/v1"
	BulkexportsV1 "github.com/twilio/twilio-go/rest/bulkexports/v1"
	ChatV1 "github.com/twilio/twilio-go/rest/chat/v1"
	ChatV2 "github.com/twilio/twilio-go/rest/chat/v2"
	ChatV3 "github.com/twilio/twilio-go/rest/chat/v3"
	ContentV1 "github.com/twilio/twilio-go/rest/content/v1"
	ContentV2 "github.com/twilio/twilio-go/rest/content/v2"
	ConversationsV1 "github.com/twilio/twilio-go/rest/conversations/v1"
	EventsV1 "github.com/twilio/twilio-go/rest/events/v1"
	FlexV1 "github.com/twilio/twilio-go/rest/flex/v1"
	FlexV2 "github.com/twilio/twilio-go/rest/flex/v2"
	FrontlineV1 "github.com/twilio/twilio-go/rest/frontline/v1"
	IamV1 "github.com/twilio/twilio-go/rest/iam/v1"
	InsightsV1 "github.com/twilio/twilio-go/rest/insights/v1"
	InsightsV2 "github.com/twilio/twilio-go/rest/insights/v2"
	IntelligenceV2 "github.com/twilio/twilio-go/rest/intelligence/v2"
	IpMessagingV1 "github.com/twilio/twilio-go/rest/ip_messaging/v1"
	IpMessagingV2 "github.com/twilio/twilio-go/rest/ip_messaging/v2"
	LookupsV1 "github.com/twilio/twilio-go/rest/lookups/v1"
	LookupsV2 "github.com/twilio/twilio-go/rest/lookups/v2"
	MarketplaceV1 "github.com/twilio/twilio-go/rest/marketplace/v1"
	MessagingV1 "github.com/twilio/twilio-go/rest/messaging/v1"
	MicrovisorV1 "github.com/twilio/twilio-go/rest/microvisor/v1"
	MonitorV1 "github.com/twilio/twilio-go/rest/monitor/v1"
	NotifyV1 "github.com/twilio/twilio-go/rest/notify/v1"
	NumbersV1 "github.com/twilio/twilio-go/rest/numbers/v1"
	NumbersV2 "github.com/twilio/twilio-go/rest/numbers/v2"
	OauthV1 "github.com/twilio/twilio-go/rest/oauth/v1"
	PreviewIamTemp "github.com/twilio/twilio-go/rest/preview_iam/temp"
	PricingV1 "github.com/twilio/twilio-go/rest/pricing/v1"
	PricingV2 "github.com/twilio/twilio-go/rest/pricing/v2"
	ProxyV1 "github.com/twilio/twilio-go/rest/proxy/v1"
	RoutesV2 "github.com/twilio/twilio-go/rest/routes/v2"
	ServerlessV1 "github.com/twilio/twilio-go/rest/serverless/v1"
	StudioV1 "github.com/twilio/twilio-go/rest/studio/v1"
	StudioV2 "github.com/twilio/twilio-go/rest/studio/v2"
	SupersimV1 "github.com/twilio/twilio-go/rest/supersim/v1"
	SyncV1 "github.com/twilio/twilio-go/rest/sync/v1"
	TaskrouterV1 "github.com/twilio/twilio-go/rest/taskrouter/v1"
	TrunkingV1 "github.com/twilio/twilio-go/rest/trunking/v1"
	TrusthubV1 "github.com/twilio/twilio-go/rest/trusthub/v1"
	VerifyV2 "github.com/twilio/twilio-go/rest/verify/v2"
	VideoV1 "github.com/twilio/twilio-go/rest/video/v1"
	VoiceV1 "github.com/twilio/twilio-go/rest/voice/v1"
	WirelessV1 "github.com/twilio/twilio-go/rest/wireless/v1"
)

// RestClient provides access to Twilio services.
type RestClient struct {
	*client.RequestHandler
	AccountsV1      *AccountsV1.ApiService
	Api             *Api.ApiService
	AssistantsV1    *AssistantsV1.ApiService
	BulkexportsV1   *BulkexportsV1.ApiService
	ChatV1          *ChatV1.ApiService
	ChatV2          *ChatV2.ApiService
	ChatV3          *ChatV3.ApiService
	ContentV1       *ContentV1.ApiService
	ContentV2       *ContentV2.ApiService
	ConversationsV1 *ConversationsV1.ApiService
	EventsV1        *EventsV1.ApiService
	FlexV1          *FlexV1.ApiService
	FlexV2          *FlexV2.ApiService
	FrontlineV1     *FrontlineV1.ApiService
	IamV1           *IamV1.ApiService
	InsightsV1      *InsightsV1.ApiService
	InsightsV2      *InsightsV2.ApiService
	IntelligenceV2  *IntelligenceV2.ApiService
	IpMessagingV1   *IpMessagingV1.ApiService
	IpMessagingV2   *IpMessagingV2.ApiService
	LookupsV1       *LookupsV1.ApiService
	LookupsV2       *LookupsV2.ApiService
	MarketplaceV1   *MarketplaceV1.ApiService
	MessagingV1     *MessagingV1.ApiService
	MicrovisorV1    *MicrovisorV1.ApiService
	MonitorV1       *MonitorV1.ApiService
	NotifyV1        *NotifyV1.ApiService
	NumbersV1       *NumbersV1.ApiService
	NumbersV2       *NumbersV2.ApiService
	OauthV1         *OauthV1.ApiService
	PreviewIamToken *PreviewIamTemp.ApiService
	PricingV1       *PricingV1.ApiService
	PricingV2       *PricingV2.ApiService
	ProxyV1         *ProxyV1.ApiService
	RoutesV2        *RoutesV2.ApiService
	ServerlessV1    *ServerlessV1.ApiService
	StudioV1        *StudioV1.ApiService
	StudioV2        *StudioV2.ApiService
	SupersimV1      *SupersimV1.ApiService
	SyncV1          *SyncV1.ApiService
	TaskrouterV1    *TaskrouterV1.ApiService
	TrunkingV1      *TrunkingV1.ApiService
	TrusthubV1      *TrusthubV1.ApiService
	VerifyV2        *VerifyV2.ApiService
	VideoV1         *VideoV1.ApiService
	VoiceV1         *VoiceV1.ApiService
	WirelessV1      *WirelessV1.ApiService
}

// Meta holds relevant pagination resources.
type Meta struct {
	FirstPageURL    *string `json:"first_page_url"`
	Key             *string `json:"key"`
	LastPageURL     *string `json:"last_page_url,omitempty"`
	NextPageURL     *string `json:"next_page_url"`
	Page            *int    `json:"page"`
	PageSize        *int    `json:"page_size"`
	PreviousPageURL *string `json:"previous_page_url"`
	URL             *string `json:"url"`
}

type ClientParams struct {
	Username                    string
	Password                    string
	AccountSid                  string
	Client                      client.BaseClient
	OrgClientCredentialProvider *OrgClientCredentialProvider
	ClientCredentialProvider    *ClientCredentialProvider
}

// CredentialProvider for Oauth Applications
type CredentialProvider struct {
	AuthType string
}

type OrgClientCredentialProvider struct {
	GrantType    string
	ClientId     string
	ClientSecret string
	OrgsToken    string
	AuthStrategy CredentialProvider
}

type ClientCredentialProvider struct {
	GrantType    string
	ClientId     string
	ClientSecret string
	AuthStrategy CredentialProvider
}

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

func GetOrgAccessToken(manager OrgTokenManager, apiService *PreviewIamTemp.ApiService) (*PreviewIamTemp.OauthV1Token, error) {
	params := &PreviewIamTemp.CreateTokenParams{}
	params.SetGrantType(manager.GrantType)
	params.SetClientId(manager.ClientId)
	params.SetClientSecret(manager.ClientSecret)
	params.SetCode(manager.Code)
	params.SetRedirectUri(manager.RedirectUri)
	params.SetAudience(manager.Audience)
	params.SetRefreshToken(manager.RefreshToken)
	params.SetScope(manager.Scope)

	token, err := apiService.CreateToken(params)
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	return token, nil
}

func GetClientAccessToken(manager ClientTokenManager, apiService *PreviewIamTemp.ApiService) (*PreviewIamTemp.OauthV1Token, error) {
	params := &PreviewIamTemp.CreateTokenParams{}
	params.SetGrantType(manager.GrantType)
	params.SetClientId(manager.ClientId)
	params.SetClientSecret(manager.ClientSecret)
	params.SetCode(manager.Code)
	params.SetRedirectUri(manager.RedirectUri)
	params.SetAudience(manager.Audience)
	params.SetRefreshToken(manager.RefreshToken)
	params.SetScope(manager.Scope)

	token, err := apiService.CreateToken(params)
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	return token, nil
}

// NewRestClientWithParams provides an initialized Twilio RestClient with params.
func NewRestClientWithParams(params ClientParams) *RestClient {
	requestHandler := client.NewRequestHandler(params.Client)

	if params.Client == nil && params.OrgClientCredentialProvider == nil && params.ClientCredentialProvider == nil {
		username := params.Username
		if username == "" {
			username = os.Getenv("TWILIO_ACCOUNT_SID")
		}

		password := params.Password
		if password == "" {
			password = os.Getenv("TWILIO_AUTH_TOKEN")
		}

		defaultClient := &client.Client{
			Credentials: client.NewCredentials(username, password),
		}

		if params.AccountSid != "" {
			defaultClient.SetAccountSid(params.AccountSid)
		} else {
			defaultClient.SetAccountSid(username)
		}
		requestHandler = client.NewRequestHandler(defaultClient)
	} else if params.OrgClientCredentialProvider != nil {
		// get AccessToken
		// set accessToken in client
		orgTokenManager := OrgTokenManager{
			GrantType:    params.OrgClientCredentialProvider.GrantType,
			ClientId:     params.OrgClientCredentialProvider.ClientId,
			ClientSecret: params.OrgClientCredentialProvider.ClientSecret,
			Code:         "",
			RedirectUri:  "",
			Audience:     "",
			RefreshToken: "",
			Scope:        ""}

		token, err := GetOrgAccessToken(orgTokenManager, NewRestClient().PreviewIamToken)
		//fmt.Println(*token.AccessToken)
		if err != nil {
		}
		fmt.Println(token)

		defaultClient := &client.Client{
			Credentials: client.NewCredentials("", ""),
		}
		if params.AccountSid != "" {
			defaultClient.SetAccountSid(params.AccountSid)
		}
		defaultClient.SetBearerToken("eyJhbGciOiJSUzI1NiIsImtpZCI6InVmQXhySkZ5VmNSNkdnQmZNQ29BTFoiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJodHRwczovL2FwaS50d2lsaW8uY29tL3YyIiwiYXpwIjoiT1EwYmFkOWI1Yjc5ODEzNjhkOGI2NDFkNWFiY2MwOThjZCIsImlhdCI6MTczOTQ0OTIwMCwiZXhwIjoxNzM5NDUyODAwLCJpc3MiOiJodHRwczovL3ByZXZpZXctaWFtLnR3aWxpby5jb20iLCJzdWIiOiJPUTBiYWQ5YjViNzk4MTM2OGQ4YjY0MWQ1YWJjYzA5OGNkIiwiZ3R5IjoiY2xpZW50X2NyZWRlbnRpYWxzIiwiaHR0cDovL3R3aWxpby9hY3QiOnsic3ViIjoiQUM1NmM5YTA0ZDJlNjM5OTE1NWM3YzRhNTJmMDQ0YzEwNSJ9LCJodHRwOi8vdHdpbGlvL3R5cCI6InZuZC50d2lsaW8ub2F1dGguYXQrand0OyIsImh0dHA6Ly90d2lsaW8vc3ViIjoiT1EwYmFkOWI1Yjc5ODEzNjhkOGI2NDFkNWFiY2MwOThjZCIsImh0dHA6Ly90d2lsaW8vdmFsaWRyZWdpb25zIjoidXMxIn0.ColuyOS-CDSLmS0E16GU6GXHEt8Y5qxj4383S1r2j9YDI3swSKNfjpWvacevLJQBqxtBMZoawV7fyxZk-X0HmmlVJN_DUtwIY04u93BdM114chiKdL4rO49kxdzxdC6tViulWlFQzpvAg3hC1W-pgBIb8v48Z6SwpotZT3HM8VGVgEYPeFb3I6aQOpvNM9o1fNGFb4GWs6V78d9XvvjhvRsB58TfoFWEyfEwZZt-yWp__fmE-VaC5acnm3ATQ6TGInUiGzH6GqaDMtluj_TW0kefcawL7mJYjWFig9Mj27THk9ESa9H8JKm1-c0tJj0yVivp-UHN0ZfvcuDiHdhcEw")
		requestHandler = client.NewRequestHandler(defaultClient)
	} else if params.ClientCredentialProvider != nil {
		// get AccessToken
		// set accessToken in client
		clientTokenManager := ClientTokenManager{
			GrantType:    params.ClientCredentialProvider.GrantType,
			ClientId:     params.ClientCredentialProvider.ClientId,
			ClientSecret: params.ClientCredentialProvider.ClientSecret,
			Code:         "",
			RedirectUri:  "",
			Audience:     "",
			RefreshToken: "",
			Scope:        ""}

		token, err := GetClientAccessToken(clientTokenManager, NewRestClient().PreviewIamToken)
		//fmt.Println(*token.AccessToken)
		if err != nil {
		}
		fmt.Println(token)

		defaultClient := &client.Client{
			Credentials: client.NewCredentials("", ""),
		}
		if params.AccountSid != "" {
			defaultClient.SetAccountSid(params.AccountSid)
		}
		defaultClient.SetBearerToken("eyJhbGciOiJSUzI1NiIsImtpZCI6InVmQXhySkZ5VmNSNkdnQmZNQ29BTFoiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJodHRwczovL2FwaS50d2lsaW8uY29tL3YyIiwiYXpwIjoiT1EwYmFkOWI1Yjc5ODEzNjhkOGI2NDFkNWFiY2MwOThjZCIsImlhdCI6MTczOTQ0OTIwMCwiZXhwIjoxNzM5NDUyODAwLCJpc3MiOiJodHRwczovL3ByZXZpZXctaWFtLnR3aWxpby5jb20iLCJzdWIiOiJPUTBiYWQ5YjViNzk4MTM2OGQ4YjY0MWQ1YWJjYzA5OGNkIiwiZ3R5IjoiY2xpZW50X2NyZWRlbnRpYWxzIiwiaHR0cDovL3R3aWxpby9hY3QiOnsic3ViIjoiQUM1NmM5YTA0ZDJlNjM5OTE1NWM3YzRhNTJmMDQ0YzEwNSJ9LCJodHRwOi8vdHdpbGlvL3R5cCI6InZuZC50d2lsaW8ub2F1dGguYXQrand0OyIsImh0dHA6Ly90d2lsaW8vc3ViIjoiT1EwYmFkOWI1Yjc5ODEzNjhkOGI2NDFkNWFiY2MwOThjZCIsImh0dHA6Ly90d2lsaW8vdmFsaWRyZWdpb25zIjoidXMxIn0.ColuyOS-CDSLmS0E16GU6GXHEt8Y5qxj4383S1r2j9YDI3swSKNfjpWvacevLJQBqxtBMZoawV7fyxZk-X0HmmlVJN_DUtwIY04u93BdM114chiKdL4rO49kxdzxdC6tViulWlFQzpvAg3hC1W-pgBIb8v48Z6SwpotZT3HM8VGVgEYPeFb3I6aQOpvNM9o1fNGFb4GWs6V78d9XvvjhvRsB58TfoFWEyfEwZZt-yWp__fmE-VaC5acnm3ATQ6TGInUiGzH6GqaDMtluj_TW0kefcawL7mJYjWFig9Mj27THk9ESa9H8JKm1-c0tJj0yVivp-UHN0ZfvcuDiHdhcEw")
		requestHandler = client.NewRequestHandler(defaultClient)
	}

	c := &RestClient{
		RequestHandler: requestHandler,
	}

	c.AccountsV1 = AccountsV1.NewApiService(c.RequestHandler)
	c.Api = Api.NewApiService(c.RequestHandler)
	c.AssistantsV1 = AssistantsV1.NewApiService(c.RequestHandler)
	c.BulkexportsV1 = BulkexportsV1.NewApiService(c.RequestHandler)
	c.ChatV1 = ChatV1.NewApiService(c.RequestHandler)
	c.ChatV2 = ChatV2.NewApiService(c.RequestHandler)
	c.ChatV3 = ChatV3.NewApiService(c.RequestHandler)
	c.ContentV1 = ContentV1.NewApiService(c.RequestHandler)
	c.ContentV2 = ContentV2.NewApiService(c.RequestHandler)
	c.ConversationsV1 = ConversationsV1.NewApiService(c.RequestHandler)
	c.EventsV1 = EventsV1.NewApiService(c.RequestHandler)
	c.FlexV1 = FlexV1.NewApiService(c.RequestHandler)
	c.FlexV2 = FlexV2.NewApiService(c.RequestHandler)
	c.FrontlineV1 = FrontlineV1.NewApiService(c.RequestHandler)
	c.IamV1 = IamV1.NewApiService(c.RequestHandler)
	c.InsightsV1 = InsightsV1.NewApiService(c.RequestHandler)
	c.InsightsV2 = InsightsV2.NewApiService(c.RequestHandler)
	c.IntelligenceV2 = IntelligenceV2.NewApiService(c.RequestHandler)
	c.IpMessagingV1 = IpMessagingV1.NewApiService(c.RequestHandler)
	c.IpMessagingV2 = IpMessagingV2.NewApiService(c.RequestHandler)
	c.LookupsV1 = LookupsV1.NewApiService(c.RequestHandler)
	c.LookupsV2 = LookupsV2.NewApiService(c.RequestHandler)
	c.MarketplaceV1 = MarketplaceV1.NewApiService(c.RequestHandler)
	c.MessagingV1 = MessagingV1.NewApiService(c.RequestHandler)
	c.MicrovisorV1 = MicrovisorV1.NewApiService(c.RequestHandler)
	c.MonitorV1 = MonitorV1.NewApiService(c.RequestHandler)
	c.NotifyV1 = NotifyV1.NewApiService(c.RequestHandler)
	c.NumbersV1 = NumbersV1.NewApiService(c.RequestHandler)
	c.NumbersV2 = NumbersV2.NewApiService(c.RequestHandler)
	c.OauthV1 = OauthV1.NewApiService(c.RequestHandler)
	c.PreviewIamToken = PreviewIamTemp.NewApiService(c.RequestHandler)
	c.PricingV1 = PricingV1.NewApiService(c.RequestHandler)
	c.PricingV2 = PricingV2.NewApiService(c.RequestHandler)
	c.ProxyV1 = ProxyV1.NewApiService(c.RequestHandler)
	c.RoutesV2 = RoutesV2.NewApiService(c.RequestHandler)
	c.ServerlessV1 = ServerlessV1.NewApiService(c.RequestHandler)
	c.StudioV1 = StudioV1.NewApiService(c.RequestHandler)
	c.StudioV2 = StudioV2.NewApiService(c.RequestHandler)
	c.SupersimV1 = SupersimV1.NewApiService(c.RequestHandler)
	c.SyncV1 = SyncV1.NewApiService(c.RequestHandler)
	c.TaskrouterV1 = TaskrouterV1.NewApiService(c.RequestHandler)
	c.TrunkingV1 = TrunkingV1.NewApiService(c.RequestHandler)
	c.TrusthubV1 = TrusthubV1.NewApiService(c.RequestHandler)
	c.VerifyV2 = VerifyV2.NewApiService(c.RequestHandler)
	c.VideoV1 = VideoV1.NewApiService(c.RequestHandler)
	c.VoiceV1 = VoiceV1.NewApiService(c.RequestHandler)
	c.WirelessV1 = WirelessV1.NewApiService(c.RequestHandler)

	return c
}

func NewRestClientWithClientCredentials(params ClientCredentialProvider) {
	// get AccessToken
	// set accessToken in client
	clientTokenManager := ClientTokenManager{
		GrantType:    params.GrantType,
		ClientId:     params.ClientId,
		ClientSecret: params.ClientSecret,
		Code:         "",
		RedirectUri:  "",
		Audience:     "",
		RefreshToken: "",
		Scope:        ""}

	token, err := GetClientAccessToken(clientTokenManager, NewRestClient().PreviewIamToken)
	fmt.Println(*token.AccessToken)
	if err != nil {
	}

	defaultClient := &client.Client{
		Credentials: client.NewCredentials("", ""),
	}
	defaultClient.SetBearerToken(*token.AccessToken)
}

// NewRestClient provides an initialized Twilio RestClient.
func NewRestClient() *RestClient {
	return NewRestClientWithParams(ClientParams{})
}

// SetTimeout sets the Timeout for Twilio HTTP requests.
func (c *RestClient) SetTimeout(timeout time.Duration) {
	c.RequestHandler.Client.SetTimeout(timeout)
}

// SetEdge sets the Edge for the Twilio request.
func (c *RestClient) SetEdge(edge string) {
	c.RequestHandler.Edge = edge
}

// SetRegion sets the Region for the Twilio request. Defaults to "us1" if an edge is provided.
func (c *RestClient) SetRegion(region string) {
	c.RequestHandler.Region = region
}
