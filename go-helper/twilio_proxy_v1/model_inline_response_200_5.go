/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse2005 struct for InlineResponse2005
type InlineResponse2005 struct {
	Interactions []ProxyV1ServiceSessionParticipantMessageInteraction `json:"interactions,omitempty"`
	Meta InlineResponse200Meta `json:"meta,omitempty"`
}
