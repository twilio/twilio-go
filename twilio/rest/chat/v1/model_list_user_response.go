/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListUserResponse struct for ListUserResponse
type ListUserResponse struct {
	Meta  ListCredentialResponseMeta `json:"Meta,omitempty"`
	Users []ChatV1ServiceUser        `json:"Users,omitempty"`
}
