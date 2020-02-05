package twilio

import (
	"encoding/json"
	"fmt"
	"time"

	twilio "github.com/twilio/twilio-go/internal"
)

// Media describes the properties of media that the service supports.
type Media struct {
	SizeLimitMB          *int    `json:"size_limit_mb"`
	CompatibilityMessage *string `json:"compatibility_message"`
}

// NewMessage sets push notification settings for when a new Message is posted to the Channel.
type NewMessage struct {
	Enabled           *bool   `json:"enabled" form:",omitempty"`
	Template          *string `json:"template" form:",omitempty"`
	Sound             *string `json:"sound" form:",omitempty"`
	BadgeCountEnabled *bool   `json:"badge_count_enabled" form:",omitempty"`
}

// BaseNotification sets push notification settings for the following types of notifications:
// Added to Channel, Invited to Channel, Removed from Channel.
type BaseNotification struct {
	Enabled  *bool   `json:"enabled" form:",omitempty"`
	Template *string `json:"template" form:",omitempty"`
	Sound    *string `json:"sound" form:",omitempty"`
}

// Notifications describes the enabled notification state of the Chat Service.
type Notifications struct {
	RemovedFromChannel *BaseNotification `json:"removed_from_channel" form:",omitempty"`
	LogEnabled         *bool             `json:"log_enabled" form:",omitempty"`
	AddedToChannel     *BaseNotification `json:"added_to_channel" form:",omitempty"`
	NewMessage         *NewMessage       `json:"new_message" form:",omitempty"`
	InvitedToChannel   *BaseNotification `json:"invited_to_channel" form:",omitempty"`
}

// ChatService is the top-level scope of all other resources in the Programmable Chat REST API.
// All other Programmable Chat resources belong to a specific Service.
// See: https://www.twilio.com/docs/chat/rest/service-resource
type ChatService struct {
	Sid                          *string            `json:"sid"`
	AccountSid                   *string            `json:"account_sid"`
	FriendlyName                 *string            `json:"friendly_name"`
	DateCreated                  *time.Time         `json:"date_created"`
	DateUpdated                  *time.Time         `json:"date_updated"`
	DefaultServiceRoleSid        *string            `json:"default_service_role_sid"`
	DefaultChannelRoleSid        *string            `json:"default_channel_role_sid"`
	DefaultChannelCreatorRoleSid *string            `json:"default_channel_creator_role_sid"`
	ReadStatusEnabled            *bool              `json:"read_status_enabled"`
	ReachabilityEnabled          *bool              `json:"reachability_enabled"`
	TypingIndicatorTimeout       *int               `json:"typing_indicator_timeout"`
	ConsumptionReportInterval    *int               `json:"consumption_report_interval"`
	Limits                       map[string]*int    `json:"limits"`
	PreWebhookURL                *string            `json:"pre_webhook_url"`
	PostWebhookURL               *string            `json:"post_webhook_url"`
	WebhookMethod                *string            `json:"webhook_method"`
	WebhookFilters               []*string          `json:"webhook_filters"`
	PreWebhookRetryCount         *int               `json:"pre_webhook_retry_count"`
	PostWebhookRetryCount        *int               `json:"post_webhook_retry_count"`
	Notifications                *Notifications     `json:"notifications,omitempty"`
	Media                        *Media             `json:"media"`
	URL                          *string            `json:"url"`
	Links                        map[string]*string `json:"links"`
}

// ChatServiceParams is the set of parameters that can be used when creating or updating a service.
type ChatServiceParams struct {
	FriendlyName                 *string         `form:",omitempty"`
	DefaultServiceRoleSID        *string         `form:"DefaultServiceRoleSid,omitempty"`
	DefaultChannelRoleSID        *string         `form:"DefaultChannelRoleSid,omitempty"`
	DefaultChannelCreatorRoleSID *string         `form:"DefaultChannelCreatorRoleSid,omitempty"`
	ReadStatusEnabled            *bool           `form:",omitempty"`
	ReachabilityEnabled          *bool           `form:",omitempty"`
	TypingIndicatorTimeout       *int            `form:",omitempty"`
	ConsumptionReportInterval    *int            `form:",omitempty"`
	Notifications                *Notifications  `form:",omitempty"`
	PreWebhookURL                *string         `form:"PreWebhookUrl,omitempty"`
	PostWebhookURL               *string         `form:"PostWebhookUrl,omitempty"`
	WebhookMethod                *string         `form:",omitempty"`
	WebhookFilters               []*string       `form:",omitempty"`
	PreWebhookRetryCount         *int            `form:",omitempty"`
	PostWebhookRetryCount        *int            `form:",omitempty"`
	Limits                       map[string]*int `form:",omitempty"`
}

// ChatServiceClient is the entrypoint for the Programmable Chat API.
type ChatServiceClient struct {
	serviceURL string
	client     *twilio.Client
}

// NewChatServiceClient constructs a new Chat Service client.
func NewChatServiceClient(request *twilio.Client) *ChatServiceClient {
	c := new(ChatServiceClient)
	c.client = request
	c.serviceURL = fmt.Sprintf("https://chat.%s/v2/Services", c.client.BaseURL)

	return c
}

// Create creates a new Chat Service.
func (c *ChatServiceClient) Create(params *ChatServiceParams) (*ChatService, error) {
	resp, err := c.client.Post(c.serviceURL, params)
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
	resp, err := c.client.Get(fmt.Sprintf("%s/%s", c.serviceURL, sid), nil)
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
	resp, err := c.client.Post(fmt.Sprintf("%s/%s", c.serviceURL, sid), params)
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
	resp, err := c.client.Delete(fmt.Sprintf("%s/%s", c.serviceURL, sid))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
