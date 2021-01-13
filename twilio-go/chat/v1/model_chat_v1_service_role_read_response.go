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
// ChatV1ServiceRoleReadResponse struct for ChatV1ServiceRoleReadResponse
type ChatV1ServiceRoleReadResponse struct {
	Meta ChatV1CredentialReadResponseMeta `json:"meta,omitempty"`
	Roles []ChatV1ServiceRole `json:"roles,omitempty"`
}
