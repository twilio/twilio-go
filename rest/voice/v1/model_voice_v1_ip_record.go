/*
 * Twilio - Voice
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

// VoiceV1IpRecord struct for VoiceV1IpRecord
type VoiceV1IpRecord struct {
	AccountSid       *string    `json:"account_sid,omitempty"`
	CidrPrefixLength *int32     `json:"cidr_prefix_length,omitempty"`
	DateCreated      *time.Time `json:"date_created,omitempty"`
	DateUpdated      *time.Time `json:"date_updated,omitempty"`
	FriendlyName     *string    `json:"friendly_name,omitempty"`
	IpAddress        *string    `json:"ip_address,omitempty"`
	Sid              *string    `json:"sid,omitempty"`
	Url              *string    `json:"url,omitempty"`
}
