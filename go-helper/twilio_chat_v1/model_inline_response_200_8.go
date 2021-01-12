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
// InlineResponse2008 struct for InlineResponse2008
type InlineResponse2008 struct {
	Channels []ChatV1ServiceUserUserChannel `json:"channels,omitempty"`
	Meta InlineResponse200Meta `json:"meta,omitempty"`
}
