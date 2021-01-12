/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	Assistants []AutopilotV1Assistant `json:"assistants,omitempty"`
	Meta InlineResponse200Meta `json:"meta,omitempty"`
}
