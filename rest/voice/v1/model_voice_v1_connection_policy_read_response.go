/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// VoiceV1ConnectionPolicyReadResponse struct for VoiceV1ConnectionPolicyReadResponse
type VoiceV1ConnectionPolicyReadResponse struct {
	ConnectionPolicies []VoiceV1ConnectionPolicy `json:"ConnectionPolicies,omitempty"`
	Meta VoiceV1ByocTrunkReadResponseMeta `json:"Meta,omitempty"`
}
