/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TrunkingV1TrunkCredentialList struct for TrunkingV1TrunkCredentialList
type TrunkingV1TrunkCredentialList struct {
	AccountSid   *string    `json:"account_sid,omitempty"`
	DateCreated  *time.Time `json:"date_created,omitempty"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          *string    `json:"sid,omitempty"`
	TrunkSid     *string    `json:"trunk_sid,omitempty"`
	Url          *string    `json:"url,omitempty"`
}
