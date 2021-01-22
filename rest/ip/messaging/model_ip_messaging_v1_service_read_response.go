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
// IpMessagingV1ServiceReadResponse struct for IpMessagingV1ServiceReadResponse
type IpMessagingV1ServiceReadResponse struct {
	Meta IpMessagingV1CredentialReadResponseMeta `json:"Meta,omitempty"`
	Services []IpMessagingV1Service `json:"Services,omitempty"`
}
