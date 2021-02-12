/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// ChatV1Service struct for ChatV1Service
type ChatV1Service struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ConsumptionReportInterval *int32 `json:"ConsumptionReportInterval,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	DefaultChannelCreatorRoleSid string `json:"DefaultChannelCreatorRoleSid,omitempty"`
	DefaultChannelRoleSid string `json:"DefaultChannelRoleSid,omitempty"`
	DefaultServiceRoleSid string `json:"DefaultServiceRoleSid,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Limits map[string]interface{} `json:"Limits,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	Notifications map[string]interface{} `json:"Notifications,omitempty"`
	PostWebhookUrl string `json:"PostWebhookUrl,omitempty"`
	PreWebhookUrl string `json:"PreWebhookUrl,omitempty"`
	ReachabilityEnabled bool `json:"ReachabilityEnabled,omitempty"`
	ReadStatusEnabled bool `json:"ReadStatusEnabled,omitempty"`
	Sid string `json:"Sid,omitempty"`
	TypingIndicatorTimeout *int32 `json:"TypingIndicatorTimeout,omitempty"`
	Url string `json:"Url,omitempty"`
	WebhookFilters []string `json:"WebhookFilters,omitempty"`
	WebhookMethod string `json:"WebhookMethod,omitempty"`
	Webhooks map[string]interface{} `json:"Webhooks,omitempty"`
}
