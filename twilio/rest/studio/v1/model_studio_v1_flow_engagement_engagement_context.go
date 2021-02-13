/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// StudioV1FlowEngagementEngagementContext struct for StudioV1FlowEngagementEngagementContext
type StudioV1FlowEngagementEngagementContext struct {
	AccountSid    string                 `json:"AccountSid,omitempty"`
	Context       map[string]interface{} `json:"Context,omitempty"`
	EngagementSid string                 `json:"EngagementSid,omitempty"`
	FlowSid       string                 `json:"FlowSid,omitempty"`
	Url           string                 `json:"Url,omitempty"`
}
