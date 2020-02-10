package twilio

import (
	"encoding/json"
	"time"
)

// Media describes the properties of media that the service supports.
type Media struct {
	SizeLimitMB          *int    `json:"size_limit_mb"`
	CompatibilityMessage *string `json:"compatibility_message"`
}

// NewMessage sets push notification settings for when a new Message is posted to the Channel.
type NewMessage struct {
	Enabled           bool    `json:"enabled"`
	Template          *string `json:"template"`
	Sound             *string `json:"sound"`
	BadgeCountEnabled *bool   `json:"badge_count_enabled"`
}

// BaseNotification sets push notification settings for the following types of notifications:
// Added to Channel, Invited to Channel, Removed from Channel.
type BaseNotification struct {
	Enabled  bool    `json:"enabled"`
	Template *string `json:"template"`
	Sound    *string `json:"sound"`
}

// Notifications describes the enabled notification state of the Chat Service.
type Notifications struct {
	RemovedFromChannel *BaseNotification `json:"removed_from_channel"`
	LogEnabled         bool              `json:"log_enabled"`
	AddedToChannel     *BaseNotification `json:"added_to_channel"`
	NewMessage         *NewMessage       `json:"new_message"`
	InvitedToChannel   *BaseNotification `json:"invited_to_channel"`
}

// ChatService is the top-level scope of all other resources in the Programmable Chat REST API.
// All other Programmable Chat resources belong to a specific Service.
// See: https://www.twilio.com/docs/chat/rest/service-resource
type ChatService struct {
	Sid                          string            `json:"sid"`
	AccountSid                   string            `json:"account_sid"`
	FriendlyName                 string            `json:"friendly_name"`
	DateCreated                  time.Time         `json:"date_created"`
	DateUpdated                  time.Time         `json:"date_updated"`
	DefaultServiceRoleSid        string            `json:"default_service_role_sid"`
	DefaultChannelRoleSid        string            `json:"default_channel_role_sid"`
	DefaultChannelCreatorRoleSid string            `json:"default_channel_creator_role_sid"`
	ReadStatusEnabled            bool              `json:"read_status_enabled"`
	ReachabilityEnabled          bool              `json:"reachability_enabled"`
	TypingIndicatorTimeout       int               `json:"typing_indicator_timeout"`
	ConsumptionReportInterval    int               `json:"consumption_report_interval"`
	Limits                       map[string]int    `json:"limits"`
	PreWebhookURL                string            `json:"pre_webhook_url"`
	PostWebhookURL               string            `json:"post_webhook_url"`
	WebhookMethod                string            `json:"webhook_method"`
	WebhookFilters               []string          `json:"webhook_filters"`
	PreWebhookRetryCount         int               `json:"pre_webhook_retry_count"`
	PostWebhookRetryCount        int               `json:"post_webhook_retry_count"`
	Notifications                *Notifications    `json:"notifications,omitempty"`
	Media                        *Media            `json:"media"`
	URL                          string            `json:"url"`
	Links                        map[string]string `json:"links"`
}

// ChatServiceParams is the set of parameters that can be used when creating or updating a service.
type ChatServiceParams struct {
	FriendlyName                 string         `url:"FriendlyName,omitempty"`
	DefaultServiceRoleSid        string         `url:"DefaultServiceRoleSid,omitempty"`
	DefaultChannelRoleSid        string         `url:"DefaultChannelRoleSid,omitempty"`
	DefaultChannelCreatorRoleSid string         `url:"DefaultChannelCreatorRoleSid,omitempty"`
	ReadStatusEnabled            bool           `url:"ReadStatusEnabled,omitempty"`
	ReachabilityEnabled          bool           `url:"ReachabilityEnabled,omitempty"`
	TypingIndicatorTimeout       int            `url:"TypingIndicatorTimeout,omitempty"`
	ConsumptionReportInterval    int            `url:"ConsumptionReportInterval,omitempty"`
	Notifications                *Notifications `url:"Notifications,omitempty"`
	PreWebhookURL                string         `url:"PreWebhookUrl,omitempty"`
	PostWebhookURL               string         `url:"PostWebhookUrl,omitempty"`
	WebhookMethod                string         `url:"WebhookMethod,omitempty"`
	WebhookFilters               []string       `url:"WebhookFilters,omitempty"`
	PreWebhookRetryCount         int            `url:"PreWebhookRetryCount,omitempty"`
	PostWebhookRetryCount        int            `url:"PostWebhookRetryCount,omitempty"`
	Limits                       map[string]int `url:"Limits,omitempty"`
}

// ChatServiceClient is the entrypoint for the Programmable Chat API.
type ChatServiceClient struct {
	baseURL string
	client  *Twilio
}

// NewChatServiceClient constructs a new Chat Service client.
func NewChatServiceClient(client *Twilio) *ChatServiceClient {
	c := new(ChatServiceClient)
	c.client = client
	c.baseURL = "https://chat.twilio.com/v2"
	return c
}

// Create creates a new Chat Service.
func (c *ChatServiceClient) Create(params *ChatServiceParams) (*ChatService, error) {
	resp, err := c.client.Post(c.url("/Services"), params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cs := &ChatService{}
	if err := json.NewDecoder(resp.Body).Decode(cs); err != nil {
		return nil, err
	}

	return cs, err
}

// Read returns the details of a Chat Service.
func (c *ChatServiceClient) Read(sid string) (*ChatService, error) {
	resp, err := c.client.Get(c.url("/Services/"+sid), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cs := &ChatService{}
	if err := json.NewDecoder(resp.Body).Decode(cs); err != nil {
		return nil, err
	}

	return cs, err
}

// Update updates a Service.
func (c *ChatServiceClient) Update(sid string, params *ChatServiceParams) (*ChatService, error) {
	resp, err := c.client.Post(c.url("/Services/"+sid), params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cs := &ChatService{}
	if err := json.NewDecoder(resp.Body).Decode(cs); err != nil {
		return nil, err
	}

	return cs, err
}

// Delete deletes a Chat Service.
func (c *ChatServiceClient) Delete(sid string) error {
	resp, err := c.client.Delete(c.url("/Services/" + sid))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ChatServiceClient) url(path string) string {
	if c.client.defaultBaseURL != nil {
		return *c.client.defaultBaseURL + path
	}

	return c.baseURL + path
}
