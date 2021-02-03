/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// PreviewSyncService struct for PreviewSyncService
type PreviewSyncService struct {
	AccountSid string `json:"AccountSid,omitempty"`
	AclEnabled bool `json:"AclEnabled,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	ReachabilityWebhooksEnabled bool `json:"ReachabilityWebhooksEnabled,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
	WebhookUrl string `json:"WebhookUrl,omitempty"`
}
