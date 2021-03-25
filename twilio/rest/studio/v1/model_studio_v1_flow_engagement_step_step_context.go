/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// StudioV1FlowEngagementStepStepContext struct for StudioV1FlowEngagementStepStepContext
type StudioV1FlowEngagementStepStepContext struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The current state of the flow
	Context *map[string]interface{} `json:"Context,omitempty"`
	// The SID of the Engagement
	EngagementSid *string `json:"EngagementSid,omitempty"`
	// The SID of the Flow
	FlowSid *string `json:"FlowSid,omitempty"`
	// Step SID
	StepSid *string `json:"StepSid,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"Url,omitempty"`
}
