/*
 * Twilio - Chat
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
	Messages []ChatV1ServiceChannelMessage `json:"messages,omitempty"`
	Meta InlineResponse200Meta `json:"meta,omitempty"`
}
