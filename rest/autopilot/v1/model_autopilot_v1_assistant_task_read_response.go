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
// AutopilotV1AssistantTaskReadResponse struct for AutopilotV1AssistantTaskReadResponse
type AutopilotV1AssistantTaskReadResponse struct {
	Meta AutopilotV1AssistantReadResponseMeta `json:"meta,omitempty"`
	Tasks []AutopilotV1AssistantTask `json:"tasks,omitempty"`
}
