/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// AutopilotV1AssistantWebhook struct for AutopilotV1AssistantWebhook
type AutopilotV1AssistantWebhook struct {
	AccountSid string `json:"account_sid,omitempty"`
	AssistantSid string `json:"assistant_sid,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	Events string `json:"events,omitempty"`
	Sid string `json:"sid,omitempty"`
	UniqueName string `json:"unique_name,omitempty"`
	Url string `json:"url,omitempty"`
	WebhookMethod string `json:"webhook_method,omitempty"`
	WebhookUrl string `json:"webhook_url,omitempty"`
}
