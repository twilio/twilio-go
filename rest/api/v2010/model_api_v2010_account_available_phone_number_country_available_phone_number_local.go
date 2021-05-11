/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal struct for ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal
type ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocal struct {
	AddressRequirements *string                  `json:"address_requirements,omitempty"`
	Beta                *bool                    `json:"beta,omitempty"`
	Capabilities        *PhoneNumberCapabilities `json:"capabilities,omitempty"`
	FriendlyName        *string                  `json:"friendly_name,omitempty"`
	IsoCountry          *string                  `json:"iso_country,omitempty"`
	Lata                *string                  `json:"lata,omitempty"`
	Latitude            *float32                 `json:"latitude,omitempty"`
	Locality            *string                  `json:"locality,omitempty"`
	Longitude           *float32                 `json:"longitude,omitempty"`
	PhoneNumber         *string                  `json:"phone_number,omitempty"`
	PostalCode          *string                  `json:"postal_code,omitempty"`
	RateCenter          *string                  `json:"rate_center,omitempty"`
	Region              *string                  `json:"region,omitempty"`
}
