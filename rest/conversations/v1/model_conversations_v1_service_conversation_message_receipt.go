/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"github.com/twilio/twilio-go/client"
	"time"
)

// ConversationsV1ServiceConversationMessageReceipt struct for ConversationsV1ServiceConversationMessageReceipt
type ConversationsV1ServiceConversationMessageReceipt struct {
	// The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this participant.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Message resource is associated with.
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
	ConversationSid *string `json:"conversation_sid,omitempty"`
	// The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to
	MessageSid *string `json:"message_sid,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// A messaging channel-specific identifier for the message delivered to participant e.g. `SMxx` for SMS, `WAxx` for Whatsapp etc.
	ChannelMessageSid *string `json:"channel_message_sid,omitempty"`
	// The unique ID of the participant the delivery receipt belongs to.
	ParticipantSid *string `json:"participant_sid,omitempty"`
	Status         *string `json:"status,omitempty"`
	// The message [delivery error code](https://www.twilio.com/docs/sms/api/message-resource#delivery-related-errors) for a `failed` status,
	ErrorCode int `json:"error_code,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this resource was last updated. `null` if the delivery receipt has not been updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// An absolute API resource URL for this delivery receipt.
	Url *string `json:"url,omitempty"`
}
