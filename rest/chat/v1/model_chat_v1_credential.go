/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ChatV1Credential struct for ChatV1Credential
type ChatV1Credential struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// [APN only] Whether to send the credential to sandbox APNs
	Sandbox *string `json:"sandbox,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The type of push-notification service the credential is for
	Type *string `json:"type,omitempty"`
	// The absolute URL of the Credential resource
	Url *string `json:"url,omitempty"`
}
