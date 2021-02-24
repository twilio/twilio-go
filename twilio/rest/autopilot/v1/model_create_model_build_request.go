/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateModelBuildRequest struct for CreateModelBuildRequest
type CreateModelBuildRequest struct {
	// The URL we should call using a POST method to send status information to your application.
	StatusCallback string `json:"StatusCallback,omitempty"`
	// An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName string `json:"UniqueName,omitempty"`
}
