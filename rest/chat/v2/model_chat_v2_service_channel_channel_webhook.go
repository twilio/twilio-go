/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ChatV2ServiceChannelChannelWebhook struct for ChatV2ServiceChannelChannelWebhook
type ChatV2ServiceChannelChannelWebhook struct {
	AccountSid    *string                 `json:"account_sid,omitempty"`
	ChannelSid    *string                 `json:"channel_sid,omitempty"`
	Configuration *map[string]interface{} `json:"configuration,omitempty"`
	DateCreated   *time.Time              `json:"date_created,omitempty"`
	DateUpdated   *time.Time              `json:"date_updated,omitempty"`
	ServiceSid    *string                 `json:"service_sid,omitempty"`
	Sid           *string                 `json:"sid,omitempty"`
	Type          *string                 `json:"type,omitempty"`
	Url           *string                 `json:"url,omitempty"`
}
