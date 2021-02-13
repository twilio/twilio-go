/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateAddressRequest struct for CreateAddressRequest
type CreateAddressRequest struct {
	// Whether we should automatically correct the address. Can be: `true` or `false` and the default is `true`. If empty or `true`, we will correct the address you provide if necessary. If `false`, we won't alter the address you provide.
	AutoCorrectAddress bool `json:"AutoCorrectAddress,omitempty"`
	// The city of the new address.
	City string `json:"City"`
	// The name to associate with the new address.
	CustomerName string `json:"CustomerName"`
	// Whether to enable emergency calling on the new address. Can be: `true` or `false`.
	EmergencyEnabled bool `json:"EmergencyEnabled,omitempty"`
	// A descriptive string that you create to describe the new address. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The ISO country code of the new address.
	IsoCountry string `json:"IsoCountry"`
	// The postal code of the new address.
	PostalCode string `json:"PostalCode"`
	// The state or region of the new address.
	Region string `json:"Region"`
	// The number and street address of the new address.
	Street string `json:"Street"`
}
