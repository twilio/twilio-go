/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListChannelResponse struct for ListChannelResponse
type ListChannelResponse struct {
	Channels []IpMessagingV1ServiceChannel `json:"Channels,omitempty"`
	Meta     ListCredentialResponseMeta    `json:"Meta,omitempty"`
}
