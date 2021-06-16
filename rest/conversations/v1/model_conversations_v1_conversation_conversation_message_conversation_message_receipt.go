/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ConversationsV1ConversationConversationMessageConversationMessageReceipt struct for ConversationsV1ConversationConversationMessageConversationMessageReceipt
type ConversationsV1ConversationConversationMessageConversationMessageReceipt struct {
	// The unique ID of the Account responsible for this participant.
	AccountSid *string `json:"account_sid,omitempty"`
	// A messaging channel-specific identifier for the message delivered to participant
	ChannelMessageSid *string `json:"channel_message_sid,omitempty"`
	// The unique ID of the Conversation for this message.
	ConversationSid *string `json:"conversation_sid,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The message [delivery error code](https://www.twilio.com/docs/sms/api/message-resource#delivery-related-errors) for a `failed` status
	ErrorCode *int32 `json:"error_code,omitempty"`
	// The SID of the message the delivery receipt belongs to
	MessageSid *string `json:"message_sid,omitempty"`
	// The unique ID of the participant the delivery receipt belongs to.
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// The message delivery status
	Status *string `json:"status,omitempty"`
	// An absolute URL for this delivery receipt.
	Url *string `json:"url,omitempty"`
}
