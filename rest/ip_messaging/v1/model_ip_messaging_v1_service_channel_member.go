/*
 * Twilio - Ip_messaging
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
// IpMessagingV1ServiceChannelMember struct for IpMessagingV1ServiceChannelMember
type IpMessagingV1ServiceChannelMember struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ChannelSid string `json:"ChannelSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Identity string `json:"Identity,omitempty"`
	LastConsumedMessageIndex *int32 `json:"LastConsumedMessageIndex,omitempty"`
	LastConsumptionTimestamp time.Time `json:"LastConsumptionTimestamp,omitempty"`
	RoleSid string `json:"RoleSid,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
}
