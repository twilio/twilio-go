/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateNetworkAccessProfileRequest struct for CreateNetworkAccessProfileRequest
type CreateNetworkAccessProfileRequest struct {
	// List of Network SIDs that this Network Access Profile will allow connections to.
	Networks []string `json:"Networks,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
	UniqueName string `json:"UniqueName,omitempty"`
}
