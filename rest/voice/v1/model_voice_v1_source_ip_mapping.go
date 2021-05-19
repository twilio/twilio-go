/*
 * Twilio - Voice
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

// VoiceV1SourceIpMapping struct for VoiceV1SourceIpMapping
type VoiceV1SourceIpMapping struct {
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The unique string that identifies an IP Record
	IpRecordSid *string `json:"ip_record_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The unique string that identifies a SIP Domain
	SipDomainSid *string `json:"sip_domain_sid,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
