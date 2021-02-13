/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateAssistantRequest struct for CreateAssistantRequest
type CreateAssistantRequest struct {
	// Space-separated list of callback events that will trigger callbacks.
	CallbackEvents string `json:"CallbackEvents,omitempty"`
	// A user-provided URL to send event callbacks to.
	CallbackUrl string `json:"CallbackUrl,omitempty"`
	// The JSON actions to be executed when the user's input is not recognized as matching any Task.
	FallbackActions map[string]interface{} `json:"FallbackActions,omitempty"`
	// A text description for the Assistant. It is non-unique and can up to 255 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The JSON actions to be executed on inbound phone calls when the Assistant has to say something first.
	InitiationActions map[string]interface{} `json:"InitiationActions,omitempty"`
	// A boolean that specifies whether queries should be logged for 30 days further training. If false, no queries will be stored, if true, queries will be stored for 30 days and deleted thereafter. Defaults to true if no value is provided.
	LogQueries bool `json:"LogQueries,omitempty"`
	// The JSON object that holds the style sheet for the assistant
	StyleSheet map[string]interface{} `json:"StyleSheet,omitempty"`
	// A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long.
	UniqueName string `json:"UniqueName,omitempty"`
}
