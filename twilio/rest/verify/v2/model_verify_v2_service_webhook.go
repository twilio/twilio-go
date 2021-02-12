/*
 * Twilio - Verify
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
// VerifyV2ServiceWebhook struct for VerifyV2ServiceWebhook
type VerifyV2ServiceWebhook struct {
	AccountSid string `json:"AccountSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	EventTypes []string `json:"EventTypes,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Status Status `json:"Status,omitempty"`
	Url string `json:"Url,omitempty"`
	WebhookMethod Methods `json:"WebhookMethod,omitempty"`
	WebhookUrl string `json:"WebhookUrl,omitempty"`
}
