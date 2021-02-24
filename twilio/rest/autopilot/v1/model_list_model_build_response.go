/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListModelBuildResponse struct for ListModelBuildResponse
type ListModelBuildResponse struct {
	Meta        ListAssistantResponseMeta        `json:"Meta,omitempty"`
	ModelBuilds []AutopilotV1AssistantModelBuild `json:"ModelBuilds,omitempty"`
}
