/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// IpMessagingV2ServiceChannelInvite struct for IpMessagingV2ServiceChannelInvite
type IpMessagingV2ServiceChannelInvite struct {
	AccountSid  *string    `json:"account_sid,omitempty"`
	ChannelSid  *string    `json:"channel_sid,omitempty"`
	CreatedBy   *string    `json:"created_by,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Identity    *string    `json:"identity,omitempty"`
	RoleSid     *string    `json:"role_sid,omitempty"`
	ServiceSid  *string    `json:"service_sid,omitempty"`
	Sid         *string    `json:"sid,omitempty"`
	Url         *string    `json:"url,omitempty"`
}
