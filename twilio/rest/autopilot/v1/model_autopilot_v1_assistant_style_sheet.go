/*
    * Twilio - Autopilot
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.10.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi
// AutopilotV1AssistantStyleSheet struct for AutopilotV1AssistantStyleSheet
type AutopilotV1AssistantStyleSheet struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The SID of the Assistant that is the parent of the resource
	AssistantSid *string `json:"AssistantSid,omitempty"`
	// The JSON string that describes the style sheet object
	Data *map[string]interface{} `json:"Data,omitempty"`
	// The absolute URL of the StyleSheet resource
	Url *string `json:"Url,omitempty"`
}
