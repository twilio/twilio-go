/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListUserChannelResponse struct for ListUserChannelResponse
type ListUserChannelResponse struct {
	Channels []ChatV1ServiceUserUserChannel `json:"Channels,omitempty"`
	Meta ListCredentialResponseMeta `json:"Meta,omitempty"`
}
