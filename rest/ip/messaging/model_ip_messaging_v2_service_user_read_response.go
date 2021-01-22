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
// IpMessagingV2ServiceUserReadResponse struct for IpMessagingV2ServiceUserReadResponse
type IpMessagingV2ServiceUserReadResponse struct {
	Meta IpMessagingV2CredentialReadResponseMeta `json:"Meta,omitempty"`
	Users []IpMessagingV2ServiceUser `json:"Users,omitempty"`
}
