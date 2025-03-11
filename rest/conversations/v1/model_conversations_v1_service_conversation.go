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
	"time"
)

// ConversationsV1ServiceConversation struct for ConversationsV1ServiceConversation
type ConversationsV1ServiceConversation struct {
	// The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this conversation.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique ID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) this conversation belongs to.
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// The unique ID of the [Messaging Service](https://www.twilio.com/docs/messaging/api/service-resource) this conversation belongs to.
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// The human-readable name of this conversation, limited to 256 characters. Optional.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL.
	UniqueName *string `json:"unique_name,omitempty"`
	// An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \"{}\" will be returned.
	Attributes *string `json:"attributes,omitempty"`
	State      *string `json:"state,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Timer date values representing state update for this conversation.
	Timers *map[string]interface{} `json:"timers,omitempty"`
	// An absolute API resource URL for this conversation.
	Url *string `json:"url,omitempty"`
	// Contains absolute URLs to access the [participants](https://www.twilio.com/docs/conversations/api/conversation-participant-resource), [messages](https://www.twilio.com/docs/conversations/api/conversation-message-resource) and [webhooks](https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource) of this conversation.
	Links    *map[string]interface{} `json:"links,omitempty"`
	Bindings *map[string]interface{} `json:"bindings,omitempty"`
}
