/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListExecutionResponse struct for ListExecutionResponse
type ListExecutionResponse struct {
	Executions []StudioV2Execution  `json:"executions,omitempty"`
	Meta       ListFlowResponseMeta `json:"meta,omitempty"`
}
