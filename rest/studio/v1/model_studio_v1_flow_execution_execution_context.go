/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// StudioV1FlowExecutionExecutionContext struct for StudioV1FlowExecutionExecutionContext
type StudioV1FlowExecutionExecutionContext struct {
	AccountSid   *string                 `json:"account_sid,omitempty"`
	Context      *map[string]interface{} `json:"context,omitempty"`
	ExecutionSid *string                 `json:"execution_sid,omitempty"`
	FlowSid      *string                 `json:"flow_sid,omitempty"`
	Url          *string                 `json:"url,omitempty"`
}
