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
// ProxyV1ServiceSessionInteractionReadResponse struct for ProxyV1ServiceSessionInteractionReadResponse
type ProxyV1ServiceSessionInteractionReadResponse struct {
	Interactions []ProxyV1ServiceSessionInteraction `json:"interactions,omitempty"`
	Meta ProxyV1ServiceReadResponseMeta `json:"meta,omitempty"`
}
