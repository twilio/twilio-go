/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// AutopilotV1AssistantDialogue struct for AutopilotV1AssistantDialogue
type AutopilotV1AssistantDialogue struct {
	AccountSid   *string                 `json:"account_sid,omitempty"`
	AssistantSid *string                 `json:"assistant_sid,omitempty"`
	Data         *map[string]interface{} `json:"data,omitempty"`
	Sid          *string                 `json:"sid,omitempty"`
	Url          *string                 `json:"url,omitempty"`
}
