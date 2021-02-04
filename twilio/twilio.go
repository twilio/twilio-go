// Package twilio provides bindings for Twilio's REST APIs.
package twilio

import (
	"github.com/twilio/twilio-go/client"
	"time"
	accountsV1 "github.com/twilio/twilio-go/rest/accounts/v1"
	apiV2010 "github.com/twilio/twilio-go/rest/api/v2010"
	autopilotV1 "github.com/twilio/twilio-go/rest/autopilot/v1"
	bulkexportsV1 "github.com/twilio/twilio-go/rest/bulkexports/v1"
	chatV1 "github.com/twilio/twilio-go/rest/chat/v1"
	chatV2 "github.com/twilio/twilio-go/rest/chat/v2"
	conversationsV1 "github.com/twilio/twilio-go/rest/conversations/v1"
	eventsV1 "github.com/twilio/twilio-go/rest/events/v1"
	faxV1 "github.com/twilio/twilio-go/rest/fax/v1"
	flexV1 "github.com/twilio/twilio-go/rest/flex/v1"
	insightsV1 "github.com/twilio/twilio-go/rest/insights/v1"
	ipMessagingV1 "github.com/twilio/twilio-go/rest/ip_messaging/v1"
	ipMessagingV2 "github.com/twilio/twilio-go/rest/ip_messaging/v2"
	lookupsV1 "github.com/twilio/twilio-go/rest/lookups/v1"
	messagingV1 "github.com/twilio/twilio-go/rest/messaging/v1"
	monitorV1 "github.com/twilio/twilio-go/rest/monitor/v1"
	notifyV1 "github.com/twilio/twilio-go/rest/notify/v1"
	numbersV2 "github.com/twilio/twilio-go/rest/numbers/v2"
	pricingV1 "github.com/twilio/twilio-go/rest/pricing/v1"
	pricingV2 "github.com/twilio/twilio-go/rest/pricing/v2"
	proxyV1 "github.com/twilio/twilio-go/rest/proxy/v1"
	serverlessV1 "github.com/twilio/twilio-go/rest/serverless/v1"
	studioV1 "github.com/twilio/twilio-go/rest/studio/v1"
	studioV2 "github.com/twilio/twilio-go/rest/studio/v2"
	supersimV1 "github.com/twilio/twilio-go/rest/supersim/v1"
	syncV1 "github.com/twilio/twilio-go/rest/sync/v1"
	taskrouterV1 "github.com/twilio/twilio-go/rest/taskrouter/v1"
	trunkingV1 "github.com/twilio/twilio-go/rest/trunking/v1"
	verifyV2 "github.com/twilio/twilio-go/rest/verify/v2"
	videoV1 "github.com/twilio/twilio-go/rest/video/v1"
	voiceV1 "github.com/twilio/twilio-go/rest/voice/v1"
	wirelessV1 "github.com/twilio/twilio-go/rest/wireless/v1"
)

// Twilio provides access to Twilio services.
type Twilio struct {
	*client.Credentials
	*client.Client
	defaultbaseURL *string
	common service
	AccountsV1 *accountsV1.DefaultApiService
	ApiV2010    *apiV2010.DefaultApiService
	AutopilotV1 *autopilotV1.DefaultApiService
	BulkexportsV1 *bulkexportsV1.DefaultApiService
	ChatV1 *chatV1.DefaultApiService
	ChatV2 *chatV2.DefaultApiService
	ConversationsV1 *conversationsV1.DefaultApiService
	EventsV1 *eventsV1.DefaultApiService
	FaxV1 *faxV1.DefaultApiService
	FlexV1 *flexV1.DefaultApiService
	InsightsV1 *insightsV1.DefaultApiService
	IpMessagingV1 *ipMessagingV1.DefaultApiService
	IpMessagingV2 *ipMessagingV2.DefaultApiService
	LookupsV1 *lookupsV1.DefaultApiService
	MessagingV1 *messagingV1.DefaultApiService
	MonitorV1 *monitorV1.DefaultApiService
	NotifyV1 *notifyV1.DefaultApiService
	NumbersV2 *numbersV2.DefaultApiService
	PricingV1 *pricingV1.DefaultApiService
	PricingV2 *pricingV2.DefaultApiService
	ProxyV1 *proxyV1.DefaultApiService
	ServerlessV1 *serverlessV1.DefaultApiService
	StudioV1 *studioV1.DefaultApiService
	StudioV2 *studioV2.DefaultApiService
	SupersimV1 *supersimV1.DefaultApiService
	SyncV1 *syncV1.DefaultApiService
	TaskrouterV1 *taskrouterV1.DefaultApiService
	TrunkingV1 *trunkingV1.DefaultApiService
	VerifyV2 *verifyV2.DefaultApiService
	VideoV1 *videoV1.DefaultApiService
	VoiceV1 *voiceV1.DefaultApiService
	WirelessV1 *wirelessV1.DefaultApiService
}

type service struct {
	client *Twilio
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

// NewClient provides an initialized Twilio client.
func NewClient(accountSID string, authToken string) *Twilio {
	credentials := &client.Credentials{AccountSID: accountSID, AuthToken: authToken}

	c := &Twilio{
		Credentials: credentials,
		Client: &client.Client{
			Credentials: credentials,
			BaseURL:     "twilio.com",
		},
	}

	c.common.client = c

	c.AccountsV1 = accountsV1.NewDefaultApiService(c.Client)
	c.ApiV2010 = apiV2010.NewDefaultApiService(c.Client)
	c.AutopilotV1 = autopilotV1.NewDefaultApiService(c.Client)
	c.BulkexportsV1 = bulkexportsV1.NewDefaultApiService(c.Client)
	c.ChatV1 = chatV1.NewDefaultApiService(c.Client)
	c.ChatV2 = chatV2.NewDefaultApiService(c.Client)
	c.ConversationsV1 = conversationsV1.NewDefaultApiService(c.Client)
	c.EventsV1 = eventsV1.NewDefaultApiService(c.Client)
	c.FaxV1 = faxV1.NewDefaultApiService(c.Client)
	c.FlexV1 = flexV1.NewDefaultApiService(c.Client)
	c.InsightsV1 = insightsV1.NewDefaultApiService(c.Client)
	c.IpMessagingV1 = ipMessagingV1.NewDefaultApiService(c.Client)
	c.IpMessagingV2 = ipMessagingV2.NewDefaultApiService(c.Client)
	c.LookupsV1 = lookupsV1.NewDefaultApiService(c.Client)
	c.MessagingV1 = messagingV1.NewDefaultApiService(c.Client)
	c.MonitorV1 = monitorV1.NewDefaultApiService(c.Client)
	c.NotifyV1 = notifyV1.NewDefaultApiService(c.Client)
	c.NumbersV2 = numbersV2.NewDefaultApiService(c.Client)
	c.PricingV1 = pricingV1.NewDefaultApiService(c.Client)
	c.ProxyV1 = proxyV1.NewDefaultApiService(c.Client)
	c.ServerlessV1 = serverlessV1.NewDefaultApiService(c.Client)
	c.StudioV1 = studioV1.NewDefaultApiService(c.Client)
	c.StudioV2 = studioV2.NewDefaultApiService(c.Client)
	c.SupersimV1 = supersimV1.NewDefaultApiService(c.Client)
	c.SyncV1 = syncV1.NewDefaultApiService(c.Client)
	c.TaskrouterV1 = taskrouterV1.NewDefaultApiService(c.Client)
	c.TrunkingV1 = trunkingV1.NewDefaultApiService(c.Client)
	c.VerifyV2 = verifyV2.NewDefaultApiService(c.Client)
	c.VideoV1 = videoV1.NewDefaultApiService(c.Client)
	c.VoiceV1 = voiceV1.NewDefaultApiService(c.Client)
	c.WirelessV1 = wirelessV1.NewDefaultApiService(c.Client)

	return c
}

// SetTimeout sets the Timeout for Twilio HTTP requests.
func (c *Twilio) SetTimeout(timeout time.Duration) {
	c.Client.SetTimeout(timeout)
}
