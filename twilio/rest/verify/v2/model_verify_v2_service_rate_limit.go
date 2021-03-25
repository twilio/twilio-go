/*
 * Twilio - Verify
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

// VerifyV2ServiceRateLimit struct for VerifyV2ServiceRateLimit
type VerifyV2ServiceRateLimit struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// Description of this Rate Limit
	Description *string `json:"Description,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"Links,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"ServiceSid,omitempty"`
	// A string that uniquely identifies this Rate Limit.
	Sid *string `json:"Sid,omitempty"`
	// A unique, developer assigned name of this Rate Limit.
	UniqueName *string `json:"UniqueName,omitempty"`
	// The URL of this resource.
	Url *string `json:"Url,omitempty"`
}
