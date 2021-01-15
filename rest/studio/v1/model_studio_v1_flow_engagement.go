/*
 * Twilio - Studio
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
// StudioV1FlowEngagement struct for StudioV1FlowEngagement
type StudioV1FlowEngagement struct {
	AccountSid string `json:"account_sid,omitempty"`
	ContactChannelAddress string `json:"contact_channel_address,omitempty"`
	ContactSid string `json:"contact_sid,omitempty"`
	Context map[string]interface{} `json:"context,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	FlowSid string `json:"flow_sid,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Sid string `json:"sid,omitempty"`
	Status string `json:"status,omitempty"`
	Url string `json:"url,omitempty"`
}
