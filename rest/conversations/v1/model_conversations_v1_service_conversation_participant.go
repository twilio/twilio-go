/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.24.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ConversationsV1ServiceConversationParticipant struct for ConversationsV1ServiceConversationParticipant
type ConversationsV1ServiceConversationParticipant struct {
	// The unique ID of the Account responsible for this participant.
	AccountSid *string `json:"account_sid,omitempty"`
	// An optional string metadata field you can use to store any data you wish.
	Attributes *string `json:"attributes,omitempty"`
	// The SID of the Conversation Service that the resource is associated with.
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// The unique ID of the Conversation for this participant.
	ConversationSid *string `json:"conversation_sid,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// A unique string identifier for the conversation participant as Conversation User.
	Identity *string `json:"identity,omitempty"`
	// Index of last “read” message in the Conversation for the Participant.
	LastReadMessageIndex *int `json:"last_read_message_index,omitempty"`
	// Timestamp of last “read” message in the Conversation for the Participant.
	LastReadTimestamp *string `json:"last_read_timestamp,omitempty"`
	// Information about how this participant exchanges messages with the conversation.
	MessagingBinding *map[string]interface{} `json:"messaging_binding,omitempty"`
	// The SID of a conversation-level Role to assign to the participant
	RoleSid *string `json:"role_sid,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// An absolute URL for this participant.
	Url *string `json:"url,omitempty"`
}
