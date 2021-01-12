/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject45 struct for InlineObject45
type InlineObject45 struct {
	// A user-provided JSON object encoded as a string to specify the actions for this task. It is optional and non-unique.
	Actions map[string]interface{} `json:"Actions,omitempty"`
	// User-provided HTTP endpoint where from the assistant fetches actions
	ActionsUrl string `json:"ActionsUrl,omitempty"`
	// A user-provided string that identifies this resource. It is non-unique and can up to 255 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long.
	UniqueName string `json:"UniqueName,omitempty"`
}
