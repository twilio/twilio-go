package twilio

import (
	"time"
)

// Media describes the properties of media that the service supports.
type Media struct {
	SizeLimitMB          int    `json:"size_limit_mb"`
	CompatibilityMessage string `json:"compatibility_message"`
}

// NewMessage describes notification setting state
type NewMessage struct {
	Enabled           bool   `json:"enabled"`
	Template          string `json:"template,omitempty"`
	Sound             string `json:"sound,omitempty"`
	BadgeCountEnabled bool   `json:"badge_count_enabled,omitempty"`
}

// BaseNotification describes notification setting state
type BaseNotification struct {
	Enabled  bool    `json:"enabled"`
	Template *string `json:"template,omitempty"`
	Sound    string  `json:"sound,omitempty"`
}

// Notifications describes the enabled notification state of the Chat Service
type Notifications struct {
	RemovedFromChannel *BaseNotification `json:"removed_from_channel,omitempty"`
	LogEnabled         bool              `json:"log_enabled,omitempty"`
	AddedToChannel     *BaseNotification `json:"added_to_channel,omitempty"`
	NewMessage         *NewMessage       `json:"new_message,omitempty"`
	InvitedToChannel   *BaseNotification `json:"invited_to_channel,omitempty"`
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
	ReadStatusEnabled            bool           `url:"ReadStatusEnabled"`
	ReachabilityEnabled          bool           `url:"ReachabilityEnabled"`
	TypingIndicatorTimeout       int            `url:"TypingIndicatorTimeout,omitempty"`
	ConsumptionReportInterval    int            `url:"ConsumptionReportInterval,omitempty"`
	Notifications                *Notifications `url:"Notifications,omitempty"`
	PreWebhookURL                string         `url:"PreWebhookUrl,omitempty"`
	PostWebhookURL               string         `url:"PostWebhookUrl,omitempty"`
	WebhookMethod                string         `url:"WebhookMethod,omitempty"`
	WebhookFilters               []interface{}  `url:"WebhookFilters,omitempty"`
	PreWebhookRetryCount         int            `url:"PreWebhookRetryCount,omitempty"`
	PostWebhookRetryCount        int            `url:"PostWebhookRetryCount,omitempty"`
	Limits                       map[string]int `url:"Limits,omitempty"`
}
