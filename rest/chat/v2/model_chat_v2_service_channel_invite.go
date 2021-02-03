/*
 * Twilio - Chat
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
// ChatV2ServiceChannelInvite struct for ChatV2ServiceChannelInvite
type ChatV2ServiceChannelInvite struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ChannelSid string `json:"ChannelSid,omitempty"`
	CreatedBy string `json:"CreatedBy,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Identity string `json:"Identity,omitempty"`
	RoleSid string `json:"RoleSid,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
}
