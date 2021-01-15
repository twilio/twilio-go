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
// ProxyV1ServiceSessionParticipantMessageInteractionReadResponse struct for ProxyV1ServiceSessionParticipantMessageInteractionReadResponse
type ProxyV1ServiceSessionParticipantMessageInteractionReadResponse struct {
	Interactions []ProxyV1ServiceSessionParticipantMessageInteraction `json:"interactions,omitempty"`
	Meta ProxyV1ServiceReadResponseMeta `json:"meta,omitempty"`
}
