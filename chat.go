package twilio

import (
	"time"
)

// A Service is the top-level scope of all other resources in the Programmable Chat REST API.
// All other Programmable Chat resources belong to a specific Service.
// See: https://www.twilio.com/docs/chat/rest/service-resource
type ChatService struct {
	Sid                          string            `json:"sid,omitempty"`
	AccountSid                   string            `json:"account_sid,omitempty"`
	FriendlyName                 string            `json:"friendly_name,omitempty"`
	DateCreated                  time.Time         `json:"date_created,omitempty"`
	DateUpdated                  time.Time         `json:"date_updated,omitempty"`
	DefaultServiceRoleSid        string            `json:"default_service_role_sid,omitempty"`
	DefaultChannelRoleSid        string            `json:"default_channel_role_sid,omitempty"`
	DefaultChannelCreatorRoleSid string            `json:"default_channel_creator_role_sid,omitempty"`
	ReadStatusEnabled            string            `json:"read_status_enabled,omitempty"`
	ReachabilityEnabled          string            `json:"reachability_enabled,omitempty"`
	TypingIndicatorTimeout       int               `json:"typing_indicator_timeout,omitempty"`
	ConsumptionReportInterval    string            `json:"consumption_report_interval,omitempty"`
	Limits                       map[string]int    `json:"limits,omitempty"`
	PreWebhookURL                string            `json:"pre_webhook_url,omitempty"`
	PostWebhookURL               string            `json:"post_webhook_url,omitempty"`
	WebhookMethod                string            `json:"webhook_method,omitempty"`
	WebhookFilters               string            `json:"webhook_filters,omitempty"`
	PreWebhookRetryCount         int               `json:"pre_webhook_retry_count,omitempty"`
	PostWebhookRetryCount        int               `json:"post_webhook_retry_count,omitempty"`
	Notifications                map[string]string `json:"notifications,omitempty"`
	Media                        map[string]string `json:"media,omitempty"`
	URL                          string            `json:"url,omitempty"`
	Links                        map[string]string `json:"links,omitempty"`
}

// ServiceParams is the set of parameters that can be used when creating or updating a service.
type ChatServiceParams struct {
	FriendlyName                 string            `url:"FriendlyName,omitempty"`
	DefaultServiceRoleSid        string            `url:"DefaultServiceRoleSid,omitempty"`
	DefaultChannelRoleSid        string            `url:"DefaultChannelRoleSid,omitempty"`
	DefaultChannelCreatorRoleSid string            `url:"DefaultChannelCreatorRoleSid,omitempty"`
	ReadStatusEnabled            string            `url:"ReadStatusEnabled,omitempty"`
	ReachabilityEnabled          string            `url:"ReachabilityEnabled,omitempty"`
	TypingIndicatorTimeout       string            `url:"TypingIndicatorTimeout,omitempty"`
	ConsumptionReportInterval    string            `url:"ConsumptionReportInterval,omitempty"`
	Notifications                map[string]string `url:"Notifications,omitempty"`
	PreWebhookURL                string            `url:"PreWebhookUrl,omitempty"`
	PostWebhookURL               string            `url:"PostWebhookUrl,omitempty"`
	WebhookMethod                string            `url:"WebhookMethod,omitempty"`
	WebhookFilters               string            `url:"WebhookFilters,omitempty"`
	PreWebhookRetryCount         int               `url:"PreWebhookRetryCount,omitempty"`
	PostWebhookRetryCount        int               `url:"PostWebhookRetryCount,omitempty"`
	Limits                       map[string]string `url:"Limits,omitempty"`
}
