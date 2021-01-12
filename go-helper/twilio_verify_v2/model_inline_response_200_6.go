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
// InlineResponse2006 struct for InlineResponse2006
type InlineResponse2006 struct {
	Buckets []VerifyV2ServiceRateLimitBucket `json:"buckets,omitempty"`
	Meta InlineResponse200Meta `json:"meta,omitempty"`
}
