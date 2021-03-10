/*
 * Twilio - Trunking
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

// TrunkingV1TrunkIpAccessControlList struct for TrunkingV1TrunkIpAccessControlList
type TrunkingV1TrunkIpAccessControlList struct {
	AccountSid   *string    `json:"AccountSid,omitempty"`
	DateCreated  *time.Time `json:"DateCreated,omitempty"`
	DateUpdated  *time.Time `json:"DateUpdated,omitempty"`
	FriendlyName *string    `json:"FriendlyName,omitempty"`
	Sid          *string    `json:"Sid,omitempty"`
	TrunkSid     *string    `json:"TrunkSid,omitempty"`
	Url          *string    `json:"Url,omitempty"`
}
