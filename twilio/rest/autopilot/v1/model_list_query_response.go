/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListQueryResponse struct for ListQueryResponse
type ListQueryResponse struct {
	Meta    ListAssistantResponseMeta   `json:"Meta,omitempty"`
	Queries []AutopilotV1AssistantQuery `json:"Queries,omitempty"`
}
