/*
    * Twilio - Studio
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.10.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi
// StudioV1FlowEngagementEngagementContext struct for StudioV1FlowEngagementEngagementContext
type StudioV1FlowEngagementEngagementContext struct {
	// Account SID
	AccountSid *string `json:"AccountSid,omitempty"`
	// Flow state
	Context *map[string]interface{} `json:"Context,omitempty"`
	// Engagement SID
	EngagementSid *string `json:"EngagementSid,omitempty"`
	// Flow SID
	FlowSid *string `json:"FlowSid,omitempty"`
	// The URL of the resource
	Url *string `json:"Url,omitempty"`
}
