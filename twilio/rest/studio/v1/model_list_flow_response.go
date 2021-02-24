/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListFlowResponse struct for ListFlowResponse
type ListFlowResponse struct {
	Flows []StudioV1Flow       `json:"Flows,omitempty"`
	Meta  ListFlowResponseMeta `json:"Meta,omitempty"`
}
