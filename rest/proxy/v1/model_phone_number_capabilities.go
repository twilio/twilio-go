/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// PhoneNumberCapabilities struct for PhoneNumberCapabilities
type PhoneNumberCapabilities struct {
	Fax   bool `json:"fax,omitempty"`
	Mms   bool `json:"mms,omitempty"`
	Sms   bool `json:"sms,omitempty"`
	Voice bool `json:"voice,omitempty"`
}
