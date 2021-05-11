/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountAddress struct for ApiV2010AccountAddress
type ApiV2010AccountAddress struct {
	AccountSid       *string `json:"account_sid,omitempty"`
	City             *string `json:"city,omitempty"`
	CustomerName     *string `json:"customer_name,omitempty"`
	DateCreated      *string `json:"date_created,omitempty"`
	DateUpdated      *string `json:"date_updated,omitempty"`
	EmergencyEnabled *bool   `json:"emergency_enabled,omitempty"`
	FriendlyName     *string `json:"friendly_name,omitempty"`
	IsoCountry       *string `json:"iso_country,omitempty"`
	PostalCode       *string `json:"postal_code,omitempty"`
	Region           *string `json:"region,omitempty"`
	Sid              *string `json:"sid,omitempty"`
	Street           *string `json:"street,omitempty"`
	Uri              *string `json:"uri,omitempty"`
	Validated        *bool   `json:"validated,omitempty"`
	Verified         *bool   `json:"verified,omitempty"`
}
