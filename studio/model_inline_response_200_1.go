/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twilio
// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	Flows []StudioV2Flow `json:"flows,omitempty"`
	Meta InlineResponse200Meta `json:"meta,omitempty"`
}
