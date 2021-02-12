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
// ListRoleResponse struct for ListRoleResponse
type ListRoleResponse struct {
	Meta ListCredentialResponseMeta `json:"Meta,omitempty"`
	Roles []ChatV1ServiceRole `json:"Roles,omitempty"`
}
