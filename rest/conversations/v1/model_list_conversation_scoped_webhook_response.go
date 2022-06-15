/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListConversationScopedWebhookResponse struct for ListConversationScopedWebhookResponse
type ListConversationScopedWebhookResponse struct {
	Meta     ListConfigurationAddressResponseMeta       `json:"meta,omitempty"`
	Webhooks []ConversationsV1ConversationScopedWebhook `json:"webhooks,omitempty"`
}
