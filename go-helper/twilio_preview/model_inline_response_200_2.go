/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse2002 struct for InlineResponse2002
type InlineResponse2002 struct {
	Fleets []PreviewDeployedDevicesFleet `json:"fleets,omitempty"`
	Meta InlineResponse200Meta `json:"meta,omitempty"`
}
