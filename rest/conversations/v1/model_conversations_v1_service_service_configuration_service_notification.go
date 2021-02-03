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
// ConversationsV1ServiceServiceConfigurationServiceNotification struct for ConversationsV1ServiceServiceConfigurationServiceNotification
type ConversationsV1ServiceServiceConfigurationServiceNotification struct {
	AccountSid string `json:"AccountSid,omitempty"`
	AddedToConversation map[string]interface{} `json:"AddedToConversation,omitempty"`
	ChatServiceSid string `json:"ChatServiceSid,omitempty"`
	LogEnabled bool `json:"LogEnabled,omitempty"`
	NewMessage map[string]interface{} `json:"NewMessage,omitempty"`
	RemovedFromConversation map[string]interface{} `json:"RemovedFromConversation,omitempty"`
	Url string `json:"Url,omitempty"`
}
