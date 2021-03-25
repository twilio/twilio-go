/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// IpMessagingV2ServiceChannelInvite struct for IpMessagingV2ServiceChannelInvite
type IpMessagingV2ServiceChannelInvite struct {
	AccountSid  *string    `json:"AccountSid,omitempty"`
	ChannelSid  *string    `json:"ChannelSid,omitempty"`
	CreatedBy   *string    `json:"CreatedBy,omitempty"`
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	Identity    *string    `json:"Identity,omitempty"`
	RoleSid     *string    `json:"RoleSid,omitempty"`
	ServiceSid  *string    `json:"ServiceSid,omitempty"`
	Sid         *string    `json:"Sid,omitempty"`
	Url         *string    `json:"Url,omitempty"`
}
