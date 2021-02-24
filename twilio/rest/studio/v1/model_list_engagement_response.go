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

// ListEngagementResponse struct for ListEngagementResponse
type ListEngagementResponse struct {
	Engagements []StudioV1FlowEngagement `json:"Engagements,omitempty"`
	Meta        ListFlowResponseMeta     `json:"Meta,omitempty"`
}
