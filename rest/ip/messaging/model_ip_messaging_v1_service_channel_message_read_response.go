/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// IpMessagingV1ServiceChannelMessageReadResponse struct for IpMessagingV1ServiceChannelMessageReadResponse
type IpMessagingV1ServiceChannelMessageReadResponse struct {
	Messages []IpMessagingV1ServiceChannelMessage `json:"Messages,omitempty"`
	Meta IpMessagingV1CredentialReadResponseMeta `json:"Meta,omitempty"`
}
