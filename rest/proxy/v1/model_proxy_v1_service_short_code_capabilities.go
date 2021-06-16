/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ProxyV1ServiceShortCodeCapabilities The capabilities of the short code
type ProxyV1ServiceShortCodeCapabilities struct {
	Fax   bool `json:"fax,omitempty"`
	Mms   bool `json:"mms,omitempty"`
	Sms   bool `json:"sms,omitempty"`
	Voice bool `json:"voice,omitempty"`
}
