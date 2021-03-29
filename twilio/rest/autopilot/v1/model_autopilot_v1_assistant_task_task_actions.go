/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// AutopilotV1AssistantTaskTaskActions struct for AutopilotV1AssistantTaskTaskActions
type AutopilotV1AssistantTaskTaskActions struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Assistant that is the parent of the Task associated with the resource
	AssistantSid *string `json:"assistant_sid,omitempty"`
	// The JSON string that specifies the actions that instruct the Assistant on how to perform the task
	Data *map[string]interface{} `json:"data,omitempty"`
	// The SID of the Task associated with the resource
	TaskSid *string `json:"task_sid,omitempty"`
	// The absolute URL of the TaskActions resource
	Url *string `json:"url,omitempty"`
}
