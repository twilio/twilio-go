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
)

// ConversationsV1Configuration struct for ConversationsV1Configuration
type ConversationsV1Configuration struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this configuration.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the default [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) used when creating a conversation.
	DefaultChatServiceSid *string `json:"default_chat_service_sid,omitempty"`
	// The SID of the default [Messaging Service](https://www.twilio.com/docs/messaging/api/service-resource) used when creating a conversation.
	DefaultMessagingServiceSid *string `json:"default_messaging_service_sid,omitempty"`
	// Default ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute.
	DefaultInactiveTimer *string `json:"default_inactive_timer,omitempty"`
	// Default ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes.
	DefaultClosedTimer *string `json:"default_closed_timer,omitempty"`
	// An absolute API resource URL for this global configuration.
	Url *string `json:"url,omitempty"`
	// Contains absolute API resource URLs to access the webhook and default service configurations.
	Links *map[string]interface{} `json:"links,omitempty"`
}
