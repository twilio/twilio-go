/*
 * Twilio - Proxy
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

// ProxyV1ServicePhoneNumber struct for ProxyV1ServicePhoneNumber
type ProxyV1ServicePhoneNumber struct {
	// The SID of the Account that created the resource
	AccountSid   *string                                `json:"account_sid,omitempty"`
	Capabilities *ProxyV1ServicePhoneNumberCapabilities `json:"capabilities,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The number of open session assigned to the number.
	InUse *int32 `json:"in_use,omitempty"`
	// Reserve the phone number for manual assignment to participants only
	IsReserved *bool `json:"is_reserved,omitempty"`
	// The ISO Country Code
	IsoCountry *string `json:"iso_country,omitempty"`
	// The phone number in E.164 format
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The SID of the PhoneNumber resource's parent Service resource
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the PhoneNumber resource
	Url *string `json:"url,omitempty"`
}
