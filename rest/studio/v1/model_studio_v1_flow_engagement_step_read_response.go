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
// StudioV1FlowEngagementStepReadResponse struct for StudioV1FlowEngagementStepReadResponse
type StudioV1FlowEngagementStepReadResponse struct {
	Meta StudioV1FlowReadResponseMeta `json:"Meta,omitempty"`
	Steps []StudioV1FlowEngagementStep `json:"Steps,omitempty"`
}
