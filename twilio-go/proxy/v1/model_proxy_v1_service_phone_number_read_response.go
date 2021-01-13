/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ProxyV1ServicePhoneNumberReadResponse struct for ProxyV1ServicePhoneNumberReadResponse
type ProxyV1ServicePhoneNumberReadResponse struct {
	Meta ProxyV1ServiceReadResponseMeta `json:"meta,omitempty"`
	PhoneNumbers []ProxyV1ServicePhoneNumber `json:"phone_numbers,omitempty"`
}
