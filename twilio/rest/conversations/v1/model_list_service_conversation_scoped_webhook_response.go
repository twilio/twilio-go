/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListServiceConversationScopedWebhookResponse struct for ListServiceConversationScopedWebhookResponse
type ListServiceConversationScopedWebhookResponse struct {
	Meta     ListConversationResponseMeta                                                `json:"Meta,omitempty"`
	Webhooks []ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook `json:"Webhooks,omitempty"`
}
