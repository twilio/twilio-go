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

// VerifyV2ServiceMessagingConfiguration struct for VerifyV2ServiceMessagingConfiguration
type VerifyV2ServiceMessagingConfiguration struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The ISO-3166-1 country code of the country or `all`.
	Country *string `json:"Country,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The SID of the Messaging Service used for this configuration.
	MessagingServiceSid *string `json:"MessagingServiceSid,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"ServiceSid,omitempty"`
	// The URL of this resource.
	Url *string `json:"Url,omitempty"`
}
