/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject struct for InlineObject
type InlineObject struct {
	// Description on change made in the revision.
	CommitMessage string `json:"CommitMessage,omitempty"`
	// JSON representation of flow definition.
	Definition map[string]interface{} `json:"Definition"`
	// The string that you assigned to describe the Flow.
	FriendlyName string `json:"FriendlyName"`
	// The status of the Flow. Can be: `draft` or `published`.
	Status string `json:"Status"`
}