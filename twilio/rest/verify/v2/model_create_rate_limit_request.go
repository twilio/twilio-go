/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateRateLimitRequest struct for CreateRateLimitRequest
type CreateRateLimitRequest struct {
	// Description of this Rate Limit
	Description string `json:"Description,omitempty"`
	// Provides a unique and addressable name to be assigned to this Rate Limit, assigned by the developer, to be optionally used in addition to SID. **This value should not contain PII.**
	UniqueName string `json:"UniqueName"`
}
