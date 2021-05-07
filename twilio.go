// Package twilio provides bindings for Twilio's REST APIs.
package twilio

import (
	"os"
	"time"

	"github.com/twilio/twilio-go/client"
	AccountsV1 "github.com/twilio/twilio-go/rest/accounts/v1"
	ApiV2010 "github.com/twilio/twilio-go/rest/api/v2010"
	AutopilotV1 "github.com/twilio/twilio-go/rest/autopilot/v1"
	BulkexportsV1 "github.com/twilio/twilio-go/rest/bulkexports/v1"
	ChatV1 "github.com/twilio/twilio-go/rest/chat/v1"
	ChatV2 "github.com/twilio/twilio-go/rest/chat/v2"
	ConversationsV1 "github.com/twilio/twilio-go/rest/conversations/v1"
	EventsV1 "github.com/twilio/twilio-go/rest/events/v1"
	FaxV1 "github.com/twilio/twilio-go/rest/fax/v1"
	FlexV1 "github.com/twilio/twilio-go/rest/flex/v1"
	InsightsV1 "github.com/twilio/twilio-go/rest/insights/v1"
	Ip_MessagingV1 "github.com/twilio/twilio-go/rest/ip_messaging/v1"
	Ip_MessagingV2 "github.com/twilio/twilio-go/rest/ip_messaging/v2"
	LookupsV1 "github.com/twilio/twilio-go/rest/lookups/v1"
	MessagingV1 "github.com/twilio/twilio-go/rest/messaging/v1"
	MonitorV1 "github.com/twilio/twilio-go/rest/monitor/v1"
	NotifyV1 "github.com/twilio/twilio-go/rest/notify/v1"
	NumbersV2 "github.com/twilio/twilio-go/rest/numbers/v2"
	PricingV1 "github.com/twilio/twilio-go/rest/pricing/v1"
	PricingV2 "github.com/twilio/twilio-go/rest/pricing/v2"
	ProxyV1 "github.com/twilio/twilio-go/rest/proxy/v1"
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
	*client.Credentials
	*client.Client
	common          service
	AccountsV1      *AccountsV1.DefaultApiService
	ApiV2010        *ApiV2010.DefaultApiService
	AutopilotV1     *AutopilotV1.DefaultApiService
	BulkexportsV1   *BulkexportsV1.DefaultApiService
	ChatV1          *ChatV1.DefaultApiService
	ChatV2          *ChatV2.DefaultApiService
	ConversationsV1 *ConversationsV1.DefaultApiService
	EventsV1        *EventsV1.DefaultApiService
	FaxV1           *FaxV1.DefaultApiService
	FlexV1          *FlexV1.DefaultApiService
	InsightsV1      *InsightsV1.DefaultApiService
	Ip_MessagingV1  *Ip_MessagingV1.DefaultApiService
	Ip_MessagingV2  *Ip_MessagingV2.DefaultApiService
	LookupsV1       *LookupsV1.DefaultApiService
	MessagingV1     *MessagingV1.DefaultApiService
	MonitorV1       *MonitorV1.DefaultApiService
	NotifyV1        *NotifyV1.DefaultApiService
	NumbersV2       *NumbersV2.DefaultApiService
	PricingV1       *PricingV1.DefaultApiService
	PricingV2       *PricingV2.DefaultApiService
	ProxyV1         *ProxyV1.DefaultApiService
	ServerlessV1    *ServerlessV1.DefaultApiService
	StudioV1        *StudioV1.DefaultApiService
	StudioV2        *StudioV2.DefaultApiService
	SupersimV1      *SupersimV1.DefaultApiService
	SyncV1          *SyncV1.DefaultApiService
	TaskrouterV1    *TaskrouterV1.DefaultApiService
	TrunkingV1      *TrunkingV1.DefaultApiService
	TrusthubV1      *TrusthubV1.DefaultApiService
	VerifyV2        *VerifyV2.DefaultApiService
	VideoV1         *VideoV1.DefaultApiService
	VoiceV1         *VoiceV1.DefaultApiService
	WirelessV1      *WirelessV1.DefaultApiService
}

type service struct {
	client *RestClient
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

type RestClientParams struct {
	AccountSid string
}

// NewRestClientWithParams provides an initialized Twilio RestClient with params.
func NewRestClientWithParams(username string, password string, params RestClientParams) *RestClient {
	credentials := &client.Credentials{Username: username, Password: password}

	c := &RestClient{
		Credentials: credentials,
		Client: &client.Client{
			Credentials: credentials,
			BaseURL:     "twilio.com",
			Edge:        os.Getenv("TWILIO_EDGE"),
			Region:      os.Getenv("TWILIO_REGION"),
			AccountSid:  params.AccountSid,
		},
	}

	c.common.client = c

	c.AccountsV1 = AccountsV1.NewDefaultApiService(c.Client)
	c.ApiV2010 = ApiV2010.NewDefaultApiService(c.Client)
	c.AutopilotV1 = AutopilotV1.NewDefaultApiService(c.Client)
	c.BulkexportsV1 = BulkexportsV1.NewDefaultApiService(c.Client)
	c.ChatV1 = ChatV1.NewDefaultApiService(c.Client)
	c.ChatV2 = ChatV2.NewDefaultApiService(c.Client)
	c.ConversationsV1 = ConversationsV1.NewDefaultApiService(c.Client)
	c.EventsV1 = EventsV1.NewDefaultApiService(c.Client)
	c.FaxV1 = FaxV1.NewDefaultApiService(c.Client)
	c.FlexV1 = FlexV1.NewDefaultApiService(c.Client)
	c.InsightsV1 = InsightsV1.NewDefaultApiService(c.Client)
	c.Ip_MessagingV1 = Ip_MessagingV1.NewDefaultApiService(c.Client)
	c.Ip_MessagingV2 = Ip_MessagingV2.NewDefaultApiService(c.Client)
	c.LookupsV1 = LookupsV1.NewDefaultApiService(c.Client)
	c.MessagingV1 = MessagingV1.NewDefaultApiService(c.Client)
	c.MonitorV1 = MonitorV1.NewDefaultApiService(c.Client)
	c.NotifyV1 = NotifyV1.NewDefaultApiService(c.Client)
	c.NumbersV2 = NumbersV2.NewDefaultApiService(c.Client)
	c.PricingV1 = PricingV1.NewDefaultApiService(c.Client)
	c.PricingV2 = PricingV2.NewDefaultApiService(c.Client)
	c.ProxyV1 = ProxyV1.NewDefaultApiService(c.Client)
	c.ServerlessV1 = ServerlessV1.NewDefaultApiService(c.Client)
	c.StudioV1 = StudioV1.NewDefaultApiService(c.Client)
	c.StudioV2 = StudioV2.NewDefaultApiService(c.Client)
	c.SupersimV1 = SupersimV1.NewDefaultApiService(c.Client)
	c.SyncV1 = SyncV1.NewDefaultApiService(c.Client)
	c.TaskrouterV1 = TaskrouterV1.NewDefaultApiService(c.Client)
	c.TrunkingV1 = TrunkingV1.NewDefaultApiService(c.Client)
	c.TrusthubV1 = TrusthubV1.NewDefaultApiService(c.Client)
	c.VerifyV2 = VerifyV2.NewDefaultApiService(c.Client)
	c.VideoV1 = VideoV1.NewDefaultApiService(c.Client)
	c.VoiceV1 = VoiceV1.NewDefaultApiService(c.Client)
	c.WirelessV1 = WirelessV1.NewDefaultApiService(c.Client)

	return c
}

// NewRestClient provides an initialized Twilio RestClient.
func NewRestClient(username string, password string) *RestClient {
	return NewRestClientWithParams(username, password, RestClientParams{
		AccountSid: username,
	})
}

// SetTimeout sets the Timeout for Twilio HTTP requests.
func (c *RestClient) SetTimeout(timeout time.Duration) {
	c.Client.SetTimeout(timeout)
}

// SetEdge sets the Edge for the Twilio request.
func (c *RestClient) SetEdge(edge string) {
	c.Client.Edge = edge
}

// SetRegion sets the Region for the Twilio request. Defaults to "us1" if an edge is provided.
func (c *RestClient) SetRegion(region string) {
	c.Client.Region = region
}
