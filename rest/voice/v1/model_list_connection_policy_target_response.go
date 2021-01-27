/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListConnectionPolicyTargetResponse struct for ListConnectionPolicyTargetResponse
type ListConnectionPolicyTargetResponse struct {
	Meta ListByocTrunkResponseMeta `json:"Meta,omitempty"`
	Targets []VoiceV1ConnectionPolicyConnectionPolicyTarget `json:"Targets,omitempty"`
}
