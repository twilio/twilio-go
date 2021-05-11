/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListEngagementResponse struct for ListEngagementResponse
type ListEngagementResponse struct {
	Engagements []StudioV1FlowEngagement `json:"engagements,omitempty"`
	Meta        ListFlowResponseMeta     `json:"meta,omitempty"`
}
