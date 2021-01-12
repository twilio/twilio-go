/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject4 struct for InlineObject4
type InlineObject4 struct {
	// The optional payload needed to verify the Challenge. E.g., a TOTP would use the numeric code.
	AuthPayload string `json:"AuthPayload,omitempty"`
}
