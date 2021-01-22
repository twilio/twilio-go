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
// ConversationsV1ConversationConversationScopedWebhookReadResponse struct for ConversationsV1ConversationConversationScopedWebhookReadResponse
type ConversationsV1ConversationConversationScopedWebhookReadResponse struct {
	Meta ConversationsV1ConversationReadResponseMeta `json:"Meta,omitempty"`
	Webhooks []ConversationsV1ConversationConversationScopedWebhook `json:"Webhooks,omitempty"`
}
