/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListUserResponse struct for ListUserResponse
type ListUserResponse struct {
	Meta  ListConversationResponseMeta `json:"meta,omitempty"`
	Users []ConversationsV1User        `json:"users,omitempty"`
}
