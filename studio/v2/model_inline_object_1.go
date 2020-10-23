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
// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	// Description on change made in the revision.
	CommitMessage string `json:"CommitMessage,omitempty"`
	// JSON representation of flow definition.
	Definition string `json:"Definition,omitempty"`
	// The string that you assigned to describe the Flow.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The status of the Flow. Can be: `draft` or `published`.
	Status string `json:"Status"`
}
