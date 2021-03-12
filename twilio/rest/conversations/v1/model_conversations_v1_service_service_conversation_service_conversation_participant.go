/*
    * Twilio - Conversations
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.10.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi
import (
	"time"
)
// ConversationsV1ServiceServiceConversationServiceConversationParticipant struct for ConversationsV1ServiceServiceConversationServiceConversationParticipant
type ConversationsV1ServiceServiceConversationServiceConversationParticipant struct {
	// The unique ID of the Account responsible for this participant.
	AccountSid *string `json:"AccountSid,omitempty"`
	// An optional string metadata field you can use to store any data you wish.
	Attributes *string `json:"Attributes,omitempty"`
	// The SID of the Conversation Service that the resource is associated with.
	ChatServiceSid *string `json:"ChatServiceSid,omitempty"`
	// The unique ID of the Conversation for this participant.
	ConversationSid *string `json:"ConversationSid,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// A unique string identifier for the conversation participant as Conversation User.
	Identity *string `json:"Identity,omitempty"`
	// Index of last “read” message in the Conversation for the Participant.
	LastReadMessageIndex *int32 `json:"LastReadMessageIndex,omitempty"`
	// Timestamp of last “read” message in the Conversation for the Participant.
	LastReadTimestamp *string `json:"LastReadTimestamp,omitempty"`
	// Information about how this participant exchanges messages with the conversation.
	MessagingBinding *map[string]interface{} `json:"MessagingBinding,omitempty"`
	// The SID of a conversation-level Role to assign to the participant
	RoleSid *string `json:"RoleSid,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"Sid,omitempty"`
	// An absolute URL for this participant.
	Url *string `json:"Url,omitempty"`
}
