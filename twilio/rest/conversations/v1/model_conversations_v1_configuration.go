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
// ConversationsV1Configuration struct for ConversationsV1Configuration
type ConversationsV1Configuration struct {
	// The SID of the Account responsible for this configuration.
	AccountSid *string `json:"AccountSid,omitempty"`
	// The SID of the default Conversation Service that every new conversation is associated with.
	DefaultChatServiceSid *string `json:"DefaultChatServiceSid,omitempty"`
	// Default ISO8601 duration when conversation will be switched to `closed` state.
	DefaultClosedTimer *string `json:"DefaultClosedTimer,omitempty"`
	// Default ISO8601 duration when conversation will be switched to `inactive` state.
	DefaultInactiveTimer *string `json:"DefaultInactiveTimer,omitempty"`
	// The SID of the default Messaging Service that every new conversation is associated with.
	DefaultMessagingServiceSid *string `json:"DefaultMessagingServiceSid,omitempty"`
	// Absolute URLs to access the webhook and default service configurations.
	Links *map[string]interface{} `json:"Links,omitempty"`
	// An absolute URL for this global configuration.
	Url *string `json:"Url,omitempty"`
}
