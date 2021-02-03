/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// AutopilotV1AssistantTaskTaskStatistics struct for AutopilotV1AssistantTaskTaskStatistics
type AutopilotV1AssistantTaskTaskStatistics struct {
	AccountSid string `json:"AccountSid,omitempty"`
	AssistantSid string `json:"AssistantSid,omitempty"`
	FieldsCount int32 `json:"FieldsCount,omitempty"`
	SamplesCount int32 `json:"SamplesCount,omitempty"`
	TaskSid string `json:"TaskSid,omitempty"`
	Url string `json:"Url,omitempty"`
}
