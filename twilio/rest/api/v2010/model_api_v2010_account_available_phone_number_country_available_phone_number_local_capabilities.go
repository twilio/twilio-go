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
// ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities Whether a phone number can receive calls or messages
type ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities struct {
	Fax bool `json:"Fax,omitempty"`
	Mms bool `json:"Mms,omitempty"`
	Sms bool `json:"Sms,omitempty"`
	Voice bool `json:"Voice,omitempty"`
}
