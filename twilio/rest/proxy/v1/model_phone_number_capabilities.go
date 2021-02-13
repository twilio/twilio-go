/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PhoneNumberCapabilities struct for PhoneNumberCapabilities
type PhoneNumberCapabilities struct {
	Fax   bool `json:"Fax,omitempty"`
	Mms   bool `json:"Mms,omitempty"`
	Sms   bool `json:"Sms,omitempty"`
	Voice bool `json:"Voice,omitempty"`
}
