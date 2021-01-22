/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TrunkingV1TrunkReadResponse struct for TrunkingV1TrunkReadResponse
type TrunkingV1TrunkReadResponse struct {
	Meta TrunkingV1TrunkReadResponseMeta `json:"Meta,omitempty"`
	Trunks []TrunkingV1Trunk `json:"Trunks,omitempty"`
}
