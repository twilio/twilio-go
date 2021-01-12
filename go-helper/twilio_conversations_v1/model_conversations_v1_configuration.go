/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ConversationsV1Configuration struct for ConversationsV1Configuration
type ConversationsV1Configuration struct {
	AccountSid string `json:"account_sid,omitempty"`
	DefaultChatServiceSid string `json:"default_chat_service_sid,omitempty"`
	DefaultClosedTimer string `json:"default_closed_timer,omitempty"`
	DefaultInactiveTimer string `json:"default_inactive_timer,omitempty"`
	DefaultMessagingServiceSid string `json:"default_messaging_service_sid,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Url string `json:"url,omitempty"`
}
