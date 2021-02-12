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
// UpdateDeploymentRequest struct for UpdateDeploymentRequest
type UpdateDeploymentRequest struct {
	// Provides a human readable descriptive text for this Deployment, up to 64 characters long
	FriendlyName string `json:"FriendlyName,omitempty"`
	// Provides the unique string identifier of the Twilio Sync service instance that will be linked to and accessible by this Deployment.
	SyncServiceSid string `json:"SyncServiceSid,omitempty"`
}
