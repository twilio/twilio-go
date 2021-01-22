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
// IpMessagingV1ServiceChannelInviteReadResponse struct for IpMessagingV1ServiceChannelInviteReadResponse
type IpMessagingV1ServiceChannelInviteReadResponse struct {
	Invites []IpMessagingV1ServiceChannelInvite `json:"Invites,omitempty"`
	Meta IpMessagingV1CredentialReadResponseMeta `json:"Meta,omitempty"`
}
