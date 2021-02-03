/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListExecutionResponse struct for ListExecutionResponse
type ListExecutionResponse struct {
	Executions []StudioV1FlowExecution `json:"Executions,omitempty"`
	Meta ListFlowResponseMeta `json:"Meta,omitempty"`
}
