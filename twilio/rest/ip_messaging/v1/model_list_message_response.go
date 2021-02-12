/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListMessageResponse struct for ListMessageResponse
type ListMessageResponse struct {
	Messages []IpMessagingV1ServiceChannelMessage `json:"Messages,omitempty"`
	Meta ListCredentialResponseMeta `json:"Meta,omitempty"`
}
