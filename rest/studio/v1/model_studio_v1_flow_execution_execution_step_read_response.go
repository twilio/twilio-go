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
// StudioV1FlowExecutionExecutionStepReadResponse struct for StudioV1FlowExecutionExecutionStepReadResponse
type StudioV1FlowExecutionExecutionStepReadResponse struct {
	Meta StudioV1FlowReadResponseMeta `json:"meta,omitempty"`
	Steps []StudioV1FlowExecutionExecutionStep `json:"steps,omitempty"`
}
