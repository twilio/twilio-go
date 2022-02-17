/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// StudioV2ExecutionContext struct for StudioV2ExecutionContext
type StudioV2ExecutionContext struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The current state of the flow
	Context *map[string]interface{} `json:"context,omitempty"`
	// The SID of the Execution
	ExecutionSid *string `json:"execution_sid,omitempty"`
	// The SID of the Flow
	FlowSid *string `json:"flow_sid,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
