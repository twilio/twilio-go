package chat

import (
	"encoding/json"
	"fmt"
	"time"

	"../../../base/backend"
)

// The Chat REST API allows you to control your Chat applications from the server in much the same way you can from the client, except from a service perspective rather than a 1st person one. You can create channels, send messages, and query the state of your messaging applications using the resources in this REST API.
// See: https://www.twilio.com/docs/chat/rest
type Chat struct {
	Request backend.Request
}

// A Service is the top-level scope of all other resources in the Programmable Chat REST API. All other Programmable Chat resources, such as Channels, Users, Messages, and Credentials belong to a specific Service.
// See: https://www.twilio.com/docs/chat/rest/service-resource
type Service struct {
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
	TypingIndicatorTimeout       string            `json:"typing_indicator_timeout,omitempty"`
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
type ServiceParams struct {
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

// Create creates a new Service.
func (chat *Chat) Create(params *ServiceParams) (*Service, error) {
	resp, err := chat.Request.Post("/Services", params)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	cs := &Service{}
	json.NewDecoder(resp.Body).Decode(cs)
	return cs, err
}

// Read returns the details of a Service.
func (chat *Chat) Read(sid string, params *ServiceParams) (*Service, error) {
	resp, err := chat.Request.Get(fmt.Sprintf("/Services/%s", sid))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	cs := &Service{}
	json.NewDecoder(resp.Body).Decode(cs)
	return cs, err
}

// Update updates a Service.
func (chat *Chat) Update(sid string, params *ServiceParams) (*Service, error) {
	resp, err := chat.Request.Post(fmt.Sprintf("/Services/%s", sid), params)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	cs := &Service{}
	json.NewDecoder(resp.Body).Decode(cs)
	return cs, err
}

// Delete deletes a Service.
func (chat *Chat) Delete(sid string, params *ServiceParams) (*Service, error) {
	resp, err := chat.Request.Delete(fmt.Sprintf("/Services/%s", sid))
	if err != nil {
		return nil, err
	}

	cs := &Service{}
	json.NewDecoder(resp.Body).Decode(cs)
	return cs, err
}
