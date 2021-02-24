/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// ChatV2ServiceChannelMessage struct for ChatV2ServiceChannelMessage
type ChatV2ServiceChannelMessage struct {
	AccountSid    *string                 `json:"AccountSid,omitempty"`
	Attributes    *string                 `json:"Attributes,omitempty"`
	Body          *string                 `json:"Body,omitempty"`
	ChannelSid    *string                 `json:"ChannelSid,omitempty"`
	DateCreated   *time.Time              `json:"DateCreated,omitempty"`
	DateUpdated   *time.Time              `json:"DateUpdated,omitempty"`
	From          *string                 `json:"From,omitempty"`
	Index         *int32                  `json:"Index,omitempty"`
	LastUpdatedBy *string                 `json:"LastUpdatedBy,omitempty"`
	Media         *map[string]interface{} `json:"Media,omitempty"`
	ServiceSid    *string                 `json:"ServiceSid,omitempty"`
	Sid           *string                 `json:"Sid,omitempty"`
	To            *string                 `json:"To,omitempty"`
	Type          *string                 `json:"Type,omitempty"`
	Url           *string                 `json:"Url,omitempty"`
	WasEdited     *bool                   `json:"WasEdited,omitempty"`
}
