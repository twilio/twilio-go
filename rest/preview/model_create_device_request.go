/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateDeviceRequest struct for CreateDeviceRequest
type CreateDeviceRequest struct {
	// Specifies the unique string identifier of the Deployment group that this Device is going to be associated with.
	DeploymentSid string `json:"DeploymentSid,omitempty"`
	Enabled bool `json:"Enabled,omitempty"`
	// Provides a human readable descriptive text to be assigned to this Device, up to 256 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// Provides an arbitrary string identifier representing a human user to be associated with this Device, up to 256 characters long.
	Identity string `json:"Identity,omitempty"`
	// Provides a unique and addressable name to be assigned to this Device, to be used in addition to SID, up to 128 characters long.
	UniqueName string `json:"UniqueName,omitempty"`
}
