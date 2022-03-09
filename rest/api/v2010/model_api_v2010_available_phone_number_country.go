/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AvailablePhoneNumberCountry struct for ApiV2010AvailablePhoneNumberCountry
type ApiV2010AvailablePhoneNumberCountry struct {
	// Whether all phone numbers available in the country are new to the Twilio platform.
	Beta *bool `json:"beta,omitempty"`
	// The name of the country
	Country *string `json:"country,omitempty"`
	// The ISO-3166-1 country code of the country.
	CountryCode *string `json:"country_code,omitempty"`
	// A list of related resources identified by their relative URIs
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	// The URI of the Country resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}
