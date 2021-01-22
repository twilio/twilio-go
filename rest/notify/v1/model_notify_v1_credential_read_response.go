/*
 * Twilio - Notify
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// NotifyV1CredentialReadResponse struct for NotifyV1CredentialReadResponse
type NotifyV1CredentialReadResponse struct {
	Credentials []NotifyV1Credential `json:"Credentials,omitempty"`
	Meta NotifyV1CredentialReadResponseMeta `json:"Meta,omitempty"`
}
