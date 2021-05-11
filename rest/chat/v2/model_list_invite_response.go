/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.14.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListInviteResponse struct for ListInviteResponse
type ListInviteResponse struct {
	Invites []ChatV2ServiceChannelInvite `json:"invites,omitempty"`
	Meta    ListCredentialResponseMeta   `json:"meta,omitempty"`
}
