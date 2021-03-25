/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListConnectionPolicyTargetResponse struct for ListConnectionPolicyTargetResponse
type ListConnectionPolicyTargetResponse struct {
	Meta    ListByocTrunkResponseMeta                       `json:"Meta,omitempty"`
	Targets []VoiceV1ConnectionPolicyConnectionPolicyTarget `json:"Targets,omitempty"`
}
