/*
    * Twilio - Chat
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.10.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi
import (
	"time"
)
// ChatV2ServiceChannelChannelWebhook struct for ChatV2ServiceChannelChannelWebhook
type ChatV2ServiceChannelChannelWebhook struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The SID of the Channel the Channel Webhook resource belongs to
	ChannelSid *string `json:"ChannelSid,omitempty"`
	// The JSON string that describes the configuration object for the channel webhook
	Configuration *map[string]interface{} `json:"Configuration,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The SID of the Service that the Channel Webhook resource is associated with
	ServiceSid *string `json:"ServiceSid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The type of webhook
	Type *string `json:"Type,omitempty"`
	// The absolute URL of the Channel Webhook resource
	Url *string `json:"Url,omitempty"`
}
