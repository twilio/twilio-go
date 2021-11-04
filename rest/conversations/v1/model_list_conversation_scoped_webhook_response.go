/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListConversationScopedWebhookResponse struct for ListConversationScopedWebhookResponse
type ListConversationScopedWebhookResponse struct {
	Meta     ListConversationResponseMeta               `json:"meta,omitempty"`
	Webhooks []ConversationsV1ConversationScopedWebhook `json:"webhooks,omitempty"`
}
