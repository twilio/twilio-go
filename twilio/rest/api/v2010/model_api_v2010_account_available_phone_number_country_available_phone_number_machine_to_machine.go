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

// ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberMachineToMachine struct for ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberMachineToMachine
type ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberMachineToMachine struct {
	// The type of Address resource the phone number requires
	AddressRequirements *string `json:"AddressRequirements,omitempty"`
	// Whether the phone number is new to the Twilio platform
	Beta         *bool                                                                            `json:"Beta,omitempty"`
	Capabilities *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities `json:"Capabilities,omitempty"`
	// A formatted version of the phone number
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The ISO country code of this phone number
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// The LATA of this phone number
	Lata *string `json:"Lata,omitempty"`
	// The latitude of this phone number's location
	Latitude *float32 `json:"Latitude,omitempty"`
	// The locality or city of this phone number's location
	Locality *string `json:"Locality,omitempty"`
	// The longitude of this phone number's location
	Longitude *float32 `json:"Longitude,omitempty"`
	// The phone number in E.164 format
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// The postal or ZIP code of this phone number's location
	PostalCode *string `json:"PostalCode,omitempty"`
	// The rate center of this phone number
	RateCenter *string `json:"RateCenter,omitempty"`
	// The two-letter state or province abbreviation of this phone number's location
	Region *string `json:"Region,omitempty"`
}
