/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountAddress struct for ApiV2010AccountAddress
type ApiV2010AccountAddress struct {
	// The SID of the Account that is responsible for the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The city in which the address is located
	City *string `json:"City,omitempty"`
	// The name associated with the address
	CustomerName *string `json:"CustomerName,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"DateUpdated,omitempty"`
	// Whether emergency calling has been enabled on this number
	EmergencyEnabled *bool `json:"EmergencyEnabled,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The ISO country code of the address
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// The postal code of the address
	PostalCode *string `json:"PostalCode,omitempty"`
	// The state or region of the address
	Region *string `json:"Region,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The number and street address of the address
	Street *string `json:"Street,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"Uri,omitempty"`
	// Whether the address has been validated to comply with local regulation
	Validated *bool `json:"Validated,omitempty"`
	// Whether the address has been verified to comply with regulation
	Verified *bool `json:"Verified,omitempty"`
}
