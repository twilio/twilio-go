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

// CreateAccessTokenRequest struct for CreateAccessTokenRequest
type CreateAccessTokenRequest struct {
	// The Type of this Factor. Eg. `push`
	FactorType string `json:"FactorType"`
	// The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, and generated by your external system, such as your user's UUID, GUID, or SID.
	Identity string `json:"Identity"`
}
