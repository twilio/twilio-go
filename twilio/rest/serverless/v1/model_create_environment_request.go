/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateEnvironmentRequest struct for CreateEnvironmentRequest
type CreateEnvironmentRequest struct {
	// A URL-friendly name that represents the environment and forms part of the domain name. It can be a maximum of 16 characters.
	DomainSuffix string `json:"DomainSuffix,omitempty"`
	// A user-defined string that uniquely identifies the Environment resource. It can be a maximum of 100 characters.
	UniqueName string `json:"UniqueName"`
}
