/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListConnectionPolicyTargetResponse struct for ListConnectionPolicyTargetResponse
type ListConnectionPolicyTargetResponse struct {
	Meta    ListByocTrunkResponseMeta       `json:"meta,omitempty"`
	Targets []VoiceV1ConnectionPolicyTarget `json:"targets,omitempty"`
}
