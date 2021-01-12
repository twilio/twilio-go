/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// StudioV1FlowEngagementStepStepContext struct for StudioV1FlowEngagementStepStepContext
type StudioV1FlowEngagementStepStepContext struct {
	AccountSid string `json:"account_sid,omitempty"`
	Context map[string]interface{} `json:"context,omitempty"`
	EngagementSid string `json:"engagement_sid,omitempty"`
	FlowSid string `json:"flow_sid,omitempty"`
	StepSid string `json:"step_sid,omitempty"`
	Url string `json:"url,omitempty"`
}
