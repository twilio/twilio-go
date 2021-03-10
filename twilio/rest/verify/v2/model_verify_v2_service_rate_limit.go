/*
 * Twilio - Verify
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

// VerifyV2ServiceRateLimit struct for VerifyV2ServiceRateLimit
type VerifyV2ServiceRateLimit struct {
	AccountSid  *string                 `json:"AccountSid,omitempty"`
	DateCreated *time.Time              `json:"DateCreated,omitempty"`
	DateUpdated *time.Time              `json:"DateUpdated,omitempty"`
	Description *string                 `json:"Description,omitempty"`
	Links       *map[string]interface{} `json:"Links,omitempty"`
	ServiceSid  *string                 `json:"ServiceSid,omitempty"`
	Sid         *string                 `json:"Sid,omitempty"`
	UniqueName  *string                 `json:"UniqueName,omitempty"`
	Url         *string                 `json:"Url,omitempty"`
}
