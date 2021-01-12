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
// InlineObject5 struct for InlineObject5
type InlineObject5 struct {
	// Provides a human readable descriptive text for this Deployment, up to 256 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// Provides the unique string identifier of the Twilio Sync service instance that will be linked to and accessible by this Deployment.
	SyncServiceSid string `json:"SyncServiceSid,omitempty"`
}
