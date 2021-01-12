/*
 * Twilio - Fax
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	Faxes []FaxV1Fax `json:"faxes,omitempty"`
	Meta InlineResponse200Meta `json:"meta,omitempty"`
}
