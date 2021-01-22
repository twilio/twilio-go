/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// AutopilotV1AssistantQueryReadResponse struct for AutopilotV1AssistantQueryReadResponse
type AutopilotV1AssistantQueryReadResponse struct {
	Meta AutopilotV1AssistantReadResponseMeta `json:"Meta,omitempty"`
	Queries []AutopilotV1AssistantQuery `json:"Queries,omitempty"`
}
