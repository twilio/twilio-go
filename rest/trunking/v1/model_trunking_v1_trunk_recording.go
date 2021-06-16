/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TrunkingV1TrunkRecording struct for TrunkingV1TrunkRecording
type TrunkingV1TrunkRecording struct {
	// The recording mode for the trunk.
	Mode *string `json:"mode,omitempty"`
	// The recording trim setting for the trunk.
	Trim *string `json:"trim,omitempty"`
}
