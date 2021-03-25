/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountOutgoingCallerId struct for ApiV2010AccountOutgoingCallerId
type ApiV2010AccountOutgoingCallerId struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"DateUpdated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The phone number in E.164 format
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"Uri,omitempty"`
}
