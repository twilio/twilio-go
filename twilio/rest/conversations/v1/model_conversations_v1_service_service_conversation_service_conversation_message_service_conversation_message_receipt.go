/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt struct for ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt
type ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ChannelMessageSid string `json:"ChannelMessageSid,omitempty"`
	ChatServiceSid string `json:"ChatServiceSid,omitempty"`
	ConversationSid string `json:"ConversationSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	ErrorCode *int32 `json:"ErrorCode,omitempty"`
	MessageSid string `json:"MessageSid,omitempty"`
	ParticipantSid string `json:"ParticipantSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Status DeliveryStatus `json:"Status,omitempty"`
	Url string `json:"Url,omitempty"`
}